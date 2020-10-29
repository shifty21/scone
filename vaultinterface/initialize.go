package vaultinterface

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/gpgcrypto"
	"github.com/shifty21/scone/rsacrypto"
	"github.com/shifty21/scone/utils"
)

//Vault struct contains vault initialization realted components
type Vault struct {
	//HTTPClient for quering vault
	HTTPClient http.Client

	SignalCh      chan os.Signal
	Stop          func()
	Opt           *Options
	EncryptKeyFun EncryptKeyFun
	ProcessKeyFun ProcessKeyFun
	//InitResponse from vault
	InitRequest *utils.InitRequest
	//InitResponse from vault
	InitResponse *utils.InitResponse
	//EncryptedResponse
	EncryptedResponse *utils.InitResponse
	//DecryptedInitResponse if response from vault is encrypted
	//or we are encrypting the response this will be used
	DecryptedInitResponse *utils.InitResponse
}

//SetConfig sets config that will verify the clients CAS config
func SetConfig(config *config.Vault) Option {
	return func(o *Options) {
		o.VaultConfig = config
	}
}

//Initialize reads config from env variables or set them to default values
func NewVaultInterface() (*Vault, error) {
	log.Println("GetConfig|Starting the vault-init service...")
	v := &Vault{
		HTTPClient: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		SignalCh: make(chan os.Signal),
		Stop: func() {
			log.Println("Shutting down")
			os.Exit(0)
		},
		InitResponse:          &utils.InitResponse{},
		InitRequest:           &utils.InitRequest{},
		DecryptedInitResponse: &utils.InitResponse{},
	}
	signal.Notify(v.SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)

	return v, nil
}

//Finalize creates objects
func (v *Vault) Finalize(option ...Option) error {
	options := &Options{
		InitializationType: "vanilla",
	}
	for _, o := range option {
		o(options)
	}
	v.Opt = options
	log.Printf("Finalizing %v", v.Opt)
	if v.Opt.EnableGPGEncryption {
		log.Println("VaultInterface|EnableGPGEncryption")
		if v.Opt.GPGCryptoConfig == nil {
			return fmt.Errorf("Please add GPGCryptoConfig")
		}
		gpgcrypto, err := gpgcrypto.InitGPGCrypto(v.Opt.GPGCryptoConfig)
		if err != nil {
			return fmt.Errorf("Error while initializing GPGCrypto : %v", err)
		}
		v.Opt.GPGCrypto = gpgcrypto
	}
	if v.Opt.SconeCryptoConfig != nil {
		log.Println("VaultInterface|Enabling SconeCrypto")
		crypto, err := rsacrypto.InitCrypto(v.Opt.SconeCryptoConfig)
		if err != nil {
			return fmt.Errorf("Error while initializing crypto module: %w", err)
		}
		v.Opt.SconeCrypto = crypto
	}
	log.Printf("Finalize|v.Opt.IsAutoInitilization %v", v.Opt.IsAutoInitilization)
	if v.Opt.IsAutoInitilization {
		log.Println("Setting Recover shares")
		v.InitRequest.RecoveryShares = 1
		v.InitRequest.RecoveryThreshold = 1
		if v.Opt.EnableGPGEncryption {
			v.InitRequest.RecoveryPGPKeys = make([]string, v.InitRequest.RecoveryShares)
			v.InitRequest.RecoveryPGPKeys[0] = v.Opt.GPGCrypto.PublicKey[0]
			v.InitRequest.RootTokenPGPKey = v.Opt.GPGCrypto.PublicKey[0]
		}
	} else {
		log.Println("Setting shamir shares")
		v.InitRequest.SecretShares = 5
		v.InitRequest.SecretThreshold = 3
		if v.Opt.EnableGPGEncryption {
			v.InitRequest.PGPKeys = make([]string, v.InitRequest.RecoveryShares)
			for x := 0; x < v.InitRequest.RecoveryShares; x++ {
				v.InitRequest.PGPKeys[x] = v.Opt.GPGCrypto.PublicKey[0]
			}
			v.InitRequest.RootTokenPGPKey = v.Opt.GPGCrypto.PublicKey[0]
		}
	}
	return nil
}

//ProcessInitVault vault, encrypt and store encrypted keys
func (v *Vault) ProcessInitVault(initRequest *utils.InitRequest) error {

	initRequestData, err := json.Marshal(&initRequest)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error marshalling initRequest %w", err)
	}
	// log.Printf("Initializing vault with %v\n", string(initRequestData))
	r := bytes.NewReader(initRequestData)
	request, err := http.NewRequest("PUT", v.Opt.VaultConfig.Address()+"/v1/sys/init", r)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error while creating init put request %w", err)
	}
	response, err := v.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error while posting init request %w", err)
	}
	defer response.Body.Close()

	initResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error reading initialization response %w", err)
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("ProcessInitVault|Non 200 status code %v with body %v: %w", response.StatusCode, string(initResponseBody), err)
	}
	var initResponse *utils.InitResponse
	if err := json.Unmarshal(initResponseBody, &initResponse); err != nil {
		return fmt.Errorf("ProcessInitVault|Error while unmarshalling reponse %w", err)
	}
	//encrypt root token
	v.InitResponse = initResponse
	err = v.EncryptKeyFun(v)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error while saving InitResponse %w", err)
	}
	return nil
}

//EncryptKeyFun takes
type EncryptKeyFun func(vault *Vault) error

//ProcessKeyFun reads key as per the type of initialization specified in configuration
type ProcessKeyFun func(vault *Vault) (*utils.InitResponse, error)

//Run vault init process
func (v *Vault) Run() {
	go func() {
		<-v.SignalCh
		v.Stop()
	}()
	for {
		err := v.CheckInitStatus()
		if err != nil {
			log.Printf("Run|Error checking initStatus %v", err)
			os.Exit(1)
		}
		response, err := v.HTTPClient.Get(v.Opt.VaultConfig.Address() + "/v1/sys/health")

		if response != nil && response.Body != nil {
			response.Body.Close()
		}

		if err != nil {
			log.Panicf("Error while checking health of vault %v", err)
			os.Exit(1)
		}
		log.Printf("Run|Response of vault health %v \n", response.StatusCode)
		switch response.StatusCode {
		case 200:
			log.Println("Run|Vault is initialized and unsealed. Exiting Program")
			os.Exit(0)
		case 429:
			log.Println("Run|Vault is unsealed and in standby mode. Exiting Program")
			os.Exit(0)
		case 501:
			log.Println("Run|Vault is not initialized. Initializing and unsealing...")
			err := v.ProcessInitVault(v.InitRequest)
			if err != nil {
				log.Printf("Run|Error with initprocess: %v", err)
				os.Exit(1)
			}
			err = v.Unseal(v.ProcessKeyFun)
			if err != nil {
				log.Printf("Run|Error unsealing: %v", err)
				os.Exit(1)
			}
		case 503:
			log.Println("Run|Vault is initialized but in sealed state. Unsealing...")
			err := v.Unseal(v.ProcessKeyFun)
			if err != nil {
				log.Printf("Run|Error Unsealing: %v", err)
				os.Exit(1)
			}
		default:
			log.Panicf("Run|Vault is in an unknown state. Status: %v", response)
			os.Exit(1)
		}
		time.After(v.Opt.VaultConfig.CheckInterval())
	}
}
