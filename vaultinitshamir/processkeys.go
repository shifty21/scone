package vaultinitshamir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ykhedar/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *vaultinterface.InitResponse) error {
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

	return nil, nil
}
