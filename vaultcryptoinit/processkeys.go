package vaultcryptoinit

import (
	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *utils.InitResponse, vault *vaultinterface.Vault) error {
	//use a public key to encrypt the response
	encryptedInitResponse, err := EncryptInitResponse(initResponse, vault)
	if err != nil {
		logger.Error.Printf("EncryptKeyFun|Error while encrypting initresponse %v", err)
		return err
	}
	InitResponse = encryptedInitResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	encryptedInitResponseJSON := InitResponse
	decryptedInitResponse, err := DecryptInitResponse(encryptedInitResponseJSON, vault)
	if err != nil {
		logger.Error.Printf("ProcessKeyFun|Error while decrypting initResponse\n")
		return nil, err
	}
	return decryptedInitResponse, nil
}
