package vaultinit

import (
	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(vault *vaultinterface.Vault) error {
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	// fmt.Printf("decryptedInitResponse|Response %v", initResponseJSON)
	return vault.InitResponse, nil
}
