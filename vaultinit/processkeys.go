package vaultinit

import (
	"errors"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *utils.InitResponse, vault *vaultinterface.Vault) error {
	vault.InitResp = initResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	auth := utils.AuthVaultByCAS(vault.CASConfig)
	if auth == false {
		return nil, errors.New("Error while authenticating with CAS")
	}
	initResponseJSON := vault.InitResp
	return initResponseJSON, nil
}
