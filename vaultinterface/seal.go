package vaultinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CheckInitStatus of vault
func CheckInitStatus() bool {
	fmt.Println("Checking InitStatus of vault...")
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
		fmt.Printf("CheckInitStatus| non 200 status code: %v\n", response)
		return false
	}

	var initStatus InitStatus

	if err := json.Unmarshal(initStatusResponseBody, &initStatus); err != nil {
		fmt.Printf("CheckInitStatus| unable to unmarshal status %v\n", err)
		return false
	}

	return initStatus.Initialized
}

//Unseal reads encrypted keys, decrypt them and calls unseal for each of them.
func Unseal(processKeyFun ProcessKeyFun) {
	initResponse, err := processKeyFun()
	if err != nil {
		fmt.Printf("Unseal|unable to read keys for vault unseal %v\n", err)
		return
	}
	fmt.Println("Unseal|Starting the unsealing process with InitResponse")

	for _, key := range initResponse.KeysBase64 {
		done, err := UnsealPerKey(key)
		if done {
			return
		}

		if err != nil {
			log.Println(err)
			return
		}
	}
}

//UnsealPerKey calls vaults unseal api with give key
func UnsealPerKey(key string) (bool, error) {
	fmt.Printf("UnsealPerKey|Unsealing for key %v\n", key)
	unsealRequest := UnsealRequest{
		Key: key,
	}

	unsealRequestData, err := json.Marshal(&unsealRequest)
	if err != nil {
		return false, err
	}

	r := bytes.NewReader(unsealRequestData)
	request, err := http.NewRequest(http.MethodPut, vaultAddr+"/v1/sys/unseal", r)
	if err != nil {
		return false, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return false, fmt.Errorf("UnsealPerKey|status code: %d", response.StatusCode)
	}

	unsealRequestResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	var unsealResponse UnsealResponse
	if err := json.Unmarshal(unsealRequestResponseBody, &unsealResponse); err != nil {
		return false, err
	}

	if !unsealResponse.Sealed {
		return true, nil
	}

	return false, nil
}
