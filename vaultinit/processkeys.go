package vaultinit

import (
	"github.com/shifty21/scone/utils"
)

//InitResponse response
var InitResponse *utils.InitResponse

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *utils.InitResponse) error {
	InitResponse = initResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func() (*utils.InitResponse, error) {
	initResponseJSON := InitResponse
	return initResponseJSON, nil
}
