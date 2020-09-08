package vaultinitshamir

import (
	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
func (v *VaultInitShamir) EncryptKeyFun(initResponse *vaultinterface.InitResponse) (*vaultinterface.InitResponse, error) {
	//use a public key to encrypt the response
	encryptedInitResponse, err := v.EncryptInitResponse(initResponse)
	if err != nil {
		logger.Error.Printf("EncryptKeyFun|Error while encrypting initresponse %v", err)
		return nil, err

	}
	if err != nil {
		logger.Error.Printf("EncryptKeyFun|Error while saving file to json")
		return nil, err
	}
	return encryptedInitResponse, nil
}

//ProcessKeyFun stores keys as required by cas unseal process
func (v *VaultInitShamir) ProcessKeyFun(encryptedInitResponseJSON *vaultinterface.InitResponse) (*vaultinterface.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	decryptedInitResponse, err := v.DecryptInitResponse(encryptedInitResponseJSON)
	if err != nil {
		logger.Error.Printf("ProcessKeyFun|Error while decrypting initResponse\n")
		return nil, err
	}
	return decryptedInitResponse, nil
}
