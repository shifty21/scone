package vaultinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/utils"
)

//CheckInitStatus of vault
func (v *Vault) CheckInitStatus() bool {
	logger.Info.Printf("CheckInitStatus|Checking InitStatus of vault")
	response, err := v.HTTPClient.Get(v.Config.Address() + "/v1/sys/init")
	if err != nil {
		logger.Error.Printf("CheckInitStatus vault initi api response error %v", err)
		return false
	}
	defer response.Body.Close()

	initStatusResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error.Printf("CheckInitStatus vault initi api response read error %v", err)
		return false
	}

	if response.StatusCode != 200 {
		logger.Error.Printf("CheckInitStatus|Non 200 status code: %v\n", response)
		return false
	}

	var initStatus utils.InitStatus

	if err := json.Unmarshal(initStatusResponseBody, &initStatus); err != nil {
		logger.Error.Printf("CheckInitStatus|Unable to unmarshal status %v\n", err)
		return false
	}
	logger.Info.Printf("CheckInitStatus|Vault initialized %v\n", initStatus.Initialized)
	return initStatus.Initialized
}

//Unseal reads encrypted keys, decrypt them and calls unseal for each of them.
func (v *Vault) Unseal(processKeyFun ProcessKeyFun) {
	// auth := utils.AuthVaultByCAS(v.CASConfig)
	// if auth == false {
	// 	logger.Error.Println("Error while authenticating with CAS")
	// 	return
	// }
	initResponse, err := processKeyFun(v)
	if err != nil {
		logger.Error.Printf("Unseal|unable to read keys for vault unseal %v\n", err)
		return
	}
	fmt.Println("Unseal|Starting the unsealing process with InitResponse")

	for i, key := range initResponse.KeysBase64 {
		done, err := v.UnsealPerKey(key)
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
func (v *Vault) UnsealPerKey(key string) (bool, error) {
	logger.Info.Printf("UnsealPerKey|Unsealing for key %v\n", key)
	unsealRequest := utils.UnsealRequest{
		Key: key,
	}

	unsealRequestData, err := json.Marshal(&unsealRequest)
	if err != nil {
		logger.Error.Println("UnsealPerKey|Error while marshalling UnsealRequest")
		return false, err
	}

	r := bytes.NewReader(unsealRequestData)
	request, err := http.NewRequest(http.MethodPut, v.Config.Address()+"/v1/sys/unseal", r)
	if err != nil {
		logger.Error.Println("UnsealPerKey|Error while creating UnsealRequest")
		return false, err
	}

	response, err := v.HTTPClient.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logger.Error.Println("UnsealPerKey|Unknown status code")
		return false, fmt.Errorf("UnsealPerKey|Unknown status code: %d", response.StatusCode)
	}

	unsealRequestResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error.Println("UnsealPerKey|Error while reading response body")
		return false, err
	}

	var unsealResponse utils.UnsealResponse
	if err := json.Unmarshal(unsealRequestResponseBody, &unsealResponse); err != nil {
		logger.Error.Println("UnsealPerKey|Error while unmarshalling unsealResponse")
		return false, err
	}

	if !unsealResponse.Sealed {
		return true, nil
	}

	return false, nil
}
