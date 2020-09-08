package vaultinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/logger"
)

//CheckInitStatus of vault
func CheckInitStatus() bool {
	logger.Info.Printf("CheckInitStatus|Checking InitStatus of vault")
	response, err := httpClient.Get(vaultAddr + "/v1/sys/init")
	if err != nil {
		log.Println(err)
		return false
	}
	defer response.Body.Close()

	initStatusResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	if response.StatusCode != 200 {
		logger.Error.Printf("CheckInitStatus|Non 200 status code: %v\n", response)
		return false
	}

	var initStatus InitStatus

	if err := json.Unmarshal(initStatusResponseBody, &initStatus); err != nil {
		logger.Error.Printf("CheckInitStatus|Unable to unmarshal status %v\n", err)
		return false
	}

	return initStatus.Initialized
}

//Unseal reads encrypted keys, decrypt them and calls unseal for each of them.
func Unseal(processKeyFun ProcessKeyFun) {
	initResponse, err := processKeyFun()
	if err != nil {
		logger.Error.Printf("Unseal|unable to read keys for vault unseal %v\n", err)
		return
	}
	fmt.Println("Unseal|Starting the unsealing process with InitResponse")

	for i, key := range initResponse.KeysBase64 {
		done, err := UnsealPerKey(key)
		if done {
			return
		}

		if err != nil {
			logger.Error.Printf("Unseal|Error while unsealing %d key %v", i, err)
			return
		}
	}
}

//UnsealPerKey calls vaults unseal api with give key
func UnsealPerKey(key string) (bool, error) {
	logger.Info.Printf("UnsealPerKey|Unsealing for key %v\n", key)
	unsealRequest := UnsealRequest{
		Key: key,
	}

	unsealRequestData, err := json.Marshal(&unsealRequest)
	if err != nil {
		logger.Error.Println("UnsealPerKey|Error while marshalling unsealrequest")
		return false, err
	}

	r := bytes.NewReader(unsealRequestData)
	request, err := http.NewRequest(http.MethodPut, vaultAddr+"/v1/sys/unseal", r)
	if err != nil {
		logger.Error.Println("UnsealPerKey|Error while creating unsealrequest")
		return false, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logger.Error.Println("UnsealPerKey|Unknown status code")
		return false, fmt.Errorf("UnsealPerKey|status code: %d", response.StatusCode)
	}

	unsealRequestResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error.Println("UnsealPerKey|Error while reading response body")
		return false, err
	}

	var unsealResponse UnsealResponse
	if err := json.Unmarshal(unsealRequestResponseBody, &unsealResponse); err != nil {
		logger.Error.Println("UnsealPerKey|Error while unmarshalling unsealResponse")
		return false, err
	}

	if !unsealResponse.Sealed {
		return true, nil
	}

	return false, nil
}
