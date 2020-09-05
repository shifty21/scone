package vaultinitshamir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *vaultinterface.InitResponse) error {
	//use a public key to encrypt the response
	initResponseJSON, _ := json.Marshal(initResponse)
	err := ioutil.WriteFile("initResponse.json", initResponseJSON, 0644)
	if err != nil {
		fmt.Printf("Error while saving file to json")
		return err
	}
	fmt.Printf("%+v\n", initResponseJSON)
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func() (*vaultinterface.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	return nil, nil
}
