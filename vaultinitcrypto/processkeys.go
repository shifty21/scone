package vaultinitcrypto

import (
	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/utils"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *utils.InitResponse) error {
	//use a public key to encrypt the response
	encryptedInitResponse, err := VaultInit.EncryptInitResponse(initResponse)
	if err != nil {
		logger.Error.Printf("EncryptKeyFun|Error while encrypting initresponse %v", err)
		return err
	}
	InitResponse = encryptedInitResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func() (*utils.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	encryptedInitResponseJSON := InitResponse
	decryptedInitResponse, err := VaultInit.DecryptInitResponse(encryptedInitResponseJSON)
	if err != nil {
		logger.Error.Printf("ProcessKeyFun|Error while decrypting initResponse\n")
		return nil, err
	}
	return decryptedInitResponse, nil
}
