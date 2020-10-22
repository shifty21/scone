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
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/utils"
)

//Vault struct contains vault initialization realted components
type Vault struct {
	//HTTPClient for quering vault
	HTTPClient http.Client

	SignalCh chan os.Signal
	Stop     func()
	//InitResp from vault
	InitResp          *utils.InitResponse
	Config            *config.Vault
	CASConfig         *config.VaultCAS
	Crypto            *crypto.Crypto
	AutoInitilization bool
}

//Initialize reads config from env variables or set them to default values
func Initialize(config *config.Configuration, crypto *crypto.Crypto) *Vault {
	log.Println("GetConfig|Starting the vault-init service...")
	VaultInterface := &Vault{
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
		Config:    config.GetVaultConfig(),
		CASConfig: config.GetCASConfig(),
		Crypto:    crypto,
	}
	signal.Notify(VaultInterface.SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	return VaultInterface
}

//ProcessInitVault vault, encrypt and store encrypted keys
func (v *Vault) ProcessInitVault(encryptKey EncryptKeyFun) error {
	var initRequest utils.InitRequest
	if !v.AutoInitilization {
		initRequest.SecretShares = 5
		initRequest.SecretThreshold = 3

	} else {
		initRequest.RecoveryShares = 1
		initRequest.RecoveryThreshold = 1
	}
	log.Printf("Initializing vault with %v\n", initRequest)
	initRequestData, err := json.Marshal(&initRequest)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error marshalling initRequest %w", err)
	}

	r := bytes.NewReader(initRequestData)
	request, err := http.NewRequest("PUT", v.Config.Address()+"/v1/sys/init", r)
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
		return fmt.Errorf("ProcessInitVault|Found non 200 status code %v", response.StatusCode)
	}
	fmt.Printf("initResponseBody %v", string(initResponseBody))
	var initResponse *utils.InitResponse
	if err := json.Unmarshal(initResponseBody, &initResponse); err != nil {
		return fmt.Errorf("ProcessInitVault|Error while unmarshalling reponse %w", err)
	}

	log.Println("ProcessInitVault|Encrypting unseal keys and the root token...")

	//encrypt root token
	err = encryptKey(initResponse, v)
	if err != nil {
		return fmt.Errorf("ProcessInitVault|Error while saving InitResponse %w", err)
	}
	return nil
}

//EncryptKeyFun takes
type EncryptKeyFun func(initResponse *utils.InitResponse, vault *Vault) error

//ProcessKeyFun reads key as per the type of initialization specified in configuration
type ProcessKeyFun func(vault *Vault) (*utils.InitResponse, error)

//Run vault init process
func (v *Vault) Run(encryptKeyFun EncryptKeyFun, processKeyFun ProcessKeyFun) {
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
		response, err := v.HTTPClient.Get(v.Config.Address() + "/v1/sys/health")

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
			err := v.ProcessInitVault(encryptKeyFun)
			if err != nil {
				log.Printf("Run|Error with initprocess: %v", err)
				os.Exit(1)
			}
			err = v.Unseal(processKeyFun)
			if err != nil {
				log.Printf("Run|Error unsealing: %v", err)
				os.Exit(1)
			}
		case 503:
			log.Println("Run|Vault is initialized but in sealed state. Unsealing...")
			err := v.Unseal(processKeyFun)
			if err != nil {
				log.Printf("Run|Error Unsealing: %v", err)
				os.Exit(1)
			}
		default:
			log.Panicf("Run|Vault is in an unknown state. Status: %v", response)
			os.Exit(1)
		}
		time.After(v.Config.CheckInterval())
	}
}
