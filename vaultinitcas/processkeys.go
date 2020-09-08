package vaultinitcas

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *vaultinterface.InitResponse) error {
	initResponseJSON, _ := json.Marshal(initResponse)
	err := ioutil.WriteFile("initResponse.json", initResponseJSON, 0644)
	if err != nil {
		logger.Info.Printf("Error while saving file to json")
		return err
	}
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func() (*vaultinterface.InitResponse, error) {
	auth := AuthVaultByCAS(&VCASConfig)
	if auth == false {
		return nil, errors.New("Error while authenticating with CAS")
	}
	initResponseByte, err := ioutil.ReadFile("initResponse.json")
	var initResponseJSON vaultinterface.InitResponse
	err = json.Unmarshal(initResponseByte, &initResponseJSON)
	if err != nil {
		logger.Info.Printf("Error while saving file to json %v", err)
		return &initResponseJSON, err
	}
	// logger.Info.Printf("%+v\n", initResponseJSON)
	return &initResponseJSON, nil
}
