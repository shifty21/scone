package vaultinitcas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *vaultinterface.InitResponse) error {
	initResponseJSON, _ := json.Marshal(initResponse)
	err := ioutil.WriteFile("initResponse.json", initResponseJSON, 0644)
	if err != nil {
		fmt.Printf("Error while saving file to json")
		return err
	}
	// fmt.Printf("%+v\n", initResponseJSON)
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func() (*vaultinterface.InitResponse, error) {
	auth := AuthCAS(&CASCONFIG)
	if auth == false {
		return nil, errors.New("Error while authenticating with CAS")
	}
	initResponseByte, err := ioutil.ReadFile("initResponse.json")
	var initResponseJSON vaultinterface.InitResponse
	err = json.Unmarshal(initResponseByte, &initResponseJSON)
	if err != nil {
		fmt.Printf("Error while saving file to json %v", err)
		return &initResponseJSON, err
	}
	// fmt.Printf("%+v\n", initResponseJSON)
	return &initResponseJSON, nil
}
