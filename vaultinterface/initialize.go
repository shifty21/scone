package vaultinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Initialize vault, encrypt and store encrypted keys
func Initialize(encryptKey EncryptKeyFun) bool {

	initRequest := InitRequest{
		SecretShares:    5,
		SecretThreshold: 3,
	}
	fmt.Printf("Initializing vault with %v\n", initRequest)
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

	log.Println("Encrypting unseal keys and the root token...")

	//encrypt root token

	// rootTokenEncryptRequest := &cloudkms.EncryptRequest{
	// 	Plaintext: base64.StdEncoding.EncodeToString([]byte(initResponse.RootToken)),
	// }

	// rootTokenEncryptResponse, err := kmsService.Projects.Locations.KeyRings.CryptoKeys.Encrypt(kmsKeyID, rootTokenEncryptRequest).Do()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	//encrypt unseal key
	// unsealKeysEncryptRequest := &cloudkms.EncryptRequest{
	// 	Plaintext: base64.StdEncoding.EncodeToString(initRequestResponseBody),
	// }

	// unsealKeysEncryptResponse, err := kmsService.Projects.Locations.KeyRings.CryptoKeys.Encrypt(kmsKeyID, unsealKeysEncryptRequest).Do()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	//write the encrypted keys to buckets
	// bucket := storageClient.Bucket(gcsBucketName)

	// // Save the encrypted unseal keys.
	// ctx := context.Background()
	// unsealKeysObject := bucket.Object("unseal-keys.json.enc").NewWriter(ctx)
	// defer unsealKeysObject.Close()

	// _, err = unsealKeysObject.Write([]byte(unsealKeysEncryptResponse.Ciphertext))
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Printf("Unseal keys written to gs://%s/%s", gcsBucketName, "unseal-keys.json.enc")

	// // Save the encrypted root token.
	// rootTokenObject := bucket.Object("root-token.enc").NewWriter(ctx)
	// defer rootTokenObject.Close()

	// _, err = rootTokenObject.Write([]byte(rootTokenEncryptResponse.Ciphertext))
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Printf("Root token written to gs://%s/%s", gcsBucketName, "root-token.enc")
	log.Println("Initialization complete")
	err = encryptKey(&initResponse)
	if err != nil {
		fmt.Printf("Initialize|Error while saving InitResponse %v", err)
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
			log.Println(err)
			time.Sleep(checkIntervalDuration)
			continue
		}

		switch response.StatusCode {
		case 200:
			log.Println("Vault is initialized and unsealed.")
		case 429:
			log.Println("Vault is unsealed and in standby mode.")
		case 501:
			log.Println("Vault is not initialized. Initializing and unsealing...")
			status := Initialize(encryptKeyFun)
			if status == false {
				log.Printf("Unable to complete initialization process of vault, continuing..")
				continue
			}
			Unseal(processKeyFun)
		case 503:
			log.Println("Vault is sealed. Unsealing...")
			Unseal(processKeyFun)
		default:
			log.Printf("Vault is in an unknown state. Status: %v", response)
		}

		log.Printf("Next check in %s", checkIntervalDuration)

		select {
		case <-signalCh:
			stop()
		case <-time.After(checkIntervalDuration):
		}
	}
}
