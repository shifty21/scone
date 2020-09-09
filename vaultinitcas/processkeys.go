package vaultinitcas

import (
	"errors"

	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *vaultinterface.InitResponse) error {
	vaultinterface.SetInitResponse(initResponse)
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func() (*vaultinterface.InitResponse, error) {
	auth := AuthVaultByCAS(&VCASConfig)
	if auth == false {
		return nil, errors.New("Error while authenticating with CAS")
	}
	initResponseJSON := vaultinterface.GetInitResponse()
	// logger.Info.Printf("%+v\n", initResponseJSON)
	return initResponseJSON, nil
}
