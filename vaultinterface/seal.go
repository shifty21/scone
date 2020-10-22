package vaultinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/utils"
)

//CheckInitStatus of vault
func (v *Vault) CheckInitStatus() error {
	log.Println("CheckInitStatus|Checking InitStatus of vault")
	response, err := v.HTTPClient.Get(v.Config.Address() + "/v1/sys/init")
	if err != nil {
		return fmt.Errorf("CheckInitStatus|vault init api response error %w", err)
	}
	defer response.Body.Close()

	initStatusResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("CheckInitStatus vault init api response read error %w", err)
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("CheckInitStatus|Non 200 status code %v with body %v: %w", response.StatusCode, string(initStatusResponseBody), err)
	}

	var initStatus utils.InitStatus

	if err := json.Unmarshal(initStatusResponseBody, &initStatus); err != nil {
		return fmt.Errorf("CheckInitStatus|Unable to unmarshal status %w", err)
	}
	log.Printf("CheckInitStatus|Vault initialized %v\n", initStatus.Initialized)
	return nil
}

//Unseal reads encrypted keys, decrypt them and calls unseal for each of them.
func (v *Vault) Unseal(processKeyFun ProcessKeyFun) error {
	initResponse, err := processKeyFun(v)
	if err != nil {
		return fmt.Errorf("Unseal|unable to read keys for vault unseal: %w", err)
	}
	log.Println("Unseal|Starting the unsealing process with InitResponse")
	var keysSet []string
	if v.AutoInitilization {
		keysSet = initResponse.RecoveryKeysBase64
	} else {
		keysSet = initResponse.Keys
	}
	for _, key := range keysSet {
		done, err := v.UnsealPerKey(key)
		if done {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Unseal|Error while unsealing key %w", err)
		}
	}
	return nil
}

//UnsealPerKey calls vaults unseal api with give key
func (v *Vault) UnsealPerKey(key string) (bool, error) {
	log.Printf("UnsealPerKey|Unsealing for key %v\n", key)
	unsealRequest := utils.UnsealRequest{
		Key: key,
	}

	unsealRequestData, err := json.Marshal(&unsealRequest)
	if err != nil {
		return false, fmt.Errorf("UnsealPerKey|Error while marshalling UnsealRequest. %w", err)
	}

	r := bytes.NewReader(unsealRequestData)
	request, err := http.NewRequest(http.MethodPut, v.Config.Address()+"/v1/sys/unseal", r)
	if err != nil {
		return false, fmt.Errorf("UnsealPerKey|Error while creating UnsealRequest. %w", err)
	}

	response, err := v.HTTPClient.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return false, fmt.Errorf("UnsealPerKey|Unknown status code: %d", response.StatusCode)
	}

	unsealRequestResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, fmt.Errorf("UnsealPerKey|Error while reading response body. %w", err)
	}
	fmt.Printf("unsealRequestResponseBody %v", string(unsealRequestResponseBody))
	var unsealResponse utils.UnsealResponse
	if err := json.Unmarshal(unsealRequestResponseBody, &unsealResponse); err != nil {
		return false, fmt.Errorf("UnsealPerKey|Error while unmarshalling unsealResponse. %w", err)
	}
	fmt.Printf("unsealResponse %v", unsealResponse)

	if !unsealResponse.Sealed {
		return true, nil
	}

	return false, nil
}
