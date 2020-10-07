package vaultinterface

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/utils"
)

//Vault struct contains vault initialization realted components
type Vault struct {
	//HTTPClient for quering vault
	HTTPClient http.Client

	SignalCh chan os.Signal
	Stop     func()
	//InitResp from vault
	InitResp  *utils.InitResponse
	Config    *config.Vault
	CASConfig *config.VaultCAS
	Crypto    *crypto.Crypto
}

//Initialize reads config from env variables or set them to default values
func Initialize(config *config.Configuration, crypto *crypto.Crypto) *Vault {
	logger.Info.Println("GetConfig|Starting the vault-init service...")
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
			logger.Error.Println("Shutting down")
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
func (v *Vault) ProcessInitVault(encryptKey EncryptKeyFun) bool {

	initRequest := utils.InitRequest{
		SecretShares:    5,
		SecretThreshold: 3,
	}
	logger.Info.Printf("Initializing vault with %v\n", initRequest)
	initRequestData, err := json.Marshal(&initRequest)
	if err != nil {
		log.Println(err)
		return false
	}

	r := bytes.NewReader(initRequestData)
	request, err := http.NewRequest("PUT", v.Config.Address()+"/v1/sys/init", r)
	if err != nil {
		log.Printf("Error while creating put request to initialize %v\n", err)
		return false
	}
	response, err := v.HTTPClient.Do(request)
	if err != nil {
		log.Printf("Error while posting request to initialize %v\n", err)
		return false
	}
	defer response.Body.Close()

	initRequestResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading initialization response %v\n", err)
		return false
	}

	if response.StatusCode != 200 {
		log.Printf("init: non 200 status code: %d", response.StatusCode)
		return false
	}

	var initResponse *utils.InitResponse
	if err := json.Unmarshal(initRequestResponseBody, &initResponse); err != nil {
		log.Printf("Error while unmarshalling reponse %v\n", err)
		return false
	}

	logger.Info.Println("Encrypting unseal keys and the root token...")

	//encrypt root token

	logger.Info.Println("Initialization complete")
	err = encryptKey(initResponse, v)
	if err != nil {
		logger.Info.Printf("Initialize|Error while saving InitResponse %v", err)
		return false
	}
	return true
}

//EncryptKeyFun takes
type EncryptKeyFun func(initResponse *utils.InitResponse, vault *Vault) error

//ProcessKeyFun reads key as per the type of initialization specified in configuration
type ProcessKeyFun func(vault *Vault) (*utils.InitResponse, error)

//Run vault init process
func (v *Vault) Run(encryptKeyFun EncryptKeyFun, processKeyFun ProcessKeyFun) {
	// auth := vaultinitcas.AuthVaultByCAS(&vaultinitcas.VCASConfig)
	// if auth == false {
	// 	logger.Error.Println("Unable to authenticate Vault, check previous logs for details.")
	// 	stop()
	// }
	for {
		select {
		case <-v.SignalCh:
			v.Stop()
		default:
			v.CheckInitStatus()
		}
		response, err := v.HTTPClient.Get(v.Config.Address() + "/v1/sys/health")

		if response != nil && response.Body != nil {
			response.Body.Close()
		}

		if err != nil {
			logger.Error.Printf("Error while checking health of vault %v", err)
			os.Exit(-1)
		}
		logger.Info.Printf("Response of vault health %v", response.StatusCode)
		switch response.StatusCode {
		case 200:
			logger.Info.Println("Run|Vault is initialized and unsealed. Exiting Program")
			return
		case 429:
			logger.Info.Println("Run|Vault is unsealed and in standby mode.")
		case 501:
			logger.Info.Println("Run|Vault is not initialized. Initializing and unsealing...")
			status := v.ProcessInitVault(encryptKeyFun)
			if status == false {
				logger.Error.Println("Run|Unable to complete initialization process of vault, continuing..")
				continue
			}
			v.Unseal(processKeyFun)
		case 503:
			logger.Info.Println("Run|Vault is sealed. Unsealing...")
			v.Unseal(processKeyFun)
		default:
			logger.Info.Printf("Run|Vault is in an unknown state. Status: %v", response)
		}

		logger.Info.Printf("Run|Next check in %s", v.Config.CheckInterval())

		select {
		case <-v.SignalCh:
			v.Stop()
		case <-time.After(v.Config.CheckInterval()):
		}
	}
}
