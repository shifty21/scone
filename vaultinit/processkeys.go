package vaultinit

import (
	"fmt"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//InitResponse response
var InitResponse *utils.InitResponse

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *utils.InitResponse, vault *vaultinterface.Vault) error {
	InitResponse = initResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	initResponseJSON := InitResponse
	fmt.Printf("decryptedInitResponse|Response %v", initResponseJSON)
	return initResponseJSON, nil
}
