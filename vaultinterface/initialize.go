package vaultinterface

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shifty21/scone/logger"
)

//Initialize vault, encrypt and store encrypted keys
func Initialize(encryptKey EncryptKeyFun) bool {

	initRequest := InitRequest{
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
	request, err := http.NewRequest("PUT", vaultAddr+"/v1/sys/init", r)
	if err != nil {
		log.Printf("Error while creating put request to initialize %v\n", err)
		return false
	}
	response, err := httpClient.Do(request)
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

	var initResponse InitResponse
	if err := json.Unmarshal(initRequestResponseBody, &initResponse); err != nil {
		log.Printf("Error while unmarshalling reponse %v\n", err)
		return false
	}
	InitResp = initResponse

	logger.Info.Println("Encrypting unseal keys and the root token...")

	//encrypt root token

	logger.Info.Println("Initialization complete")
	err = encryptKey(&initResponse)
	if err != nil {
		logger.Info.Printf("Initialize|Error while saving InitResponse %v", err)
		return false
	}
	return true
}

//EncryptKeyFun takes
type EncryptKeyFun func(*InitResponse) error

//ProcessKeyFun reads key as per the type of initialization specified in configuration
type ProcessKeyFun func() (*InitResponse, error)

//Run vault init process
func Run(encryptKeyFun EncryptKeyFun, processKeyFun ProcessKeyFun) {

	for {
		select {
		case <-signalCh:
			stop()
		default:
			CheckInitStatus()
		}
		response, err := httpClient.Get(vaultAddr + "/v1/sys/health")

		if response != nil && response.Body != nil {
			response.Body.Close()
		}

		if err != nil {
			logger.Error.Println(err)
			time.Sleep(checkIntervalDuration)
			continue
		}

		switch response.StatusCode {
		case 200:
			logger.Info.Println("Run|Vault is initialized and unsealed.")
		case 429:
			logger.Info.Println("Run|Vault is unsealed and in standby mode.")
		case 501:
			logger.Info.Println("Run|Vault is not initialized. Initializing and unsealing...")
			status := Initialize(encryptKeyFun)
			if status == false {
				logger.Error.Println("Run|Unable to complete initialization process of vault, continuing..")
				continue
			}
			Unseal(processKeyFun)
		case 503:
			logger.Info.Println("Run|Vault is sealed. Unsealing...")
			Unseal(processKeyFun)
		default:
			logger.Info.Printf("Run|Vault is in an unknown state. Status: %v", response)
		}

		logger.Info.Printf("Run|Next check in %s", checkIntervalDuration)

		select {
		case <-signalCh:
			stop()
		case <-time.After(checkIntervalDuration):
		}
	}
}
