package vaultgpginit

import (
	"fmt"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(initResponse *utils.InitResponse, vault *vaultinterface.Vault) error {
	InitResponse = initResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	encryptedInitResponseJSON := InitResponse
	decryptedInitResponse, err := DecryptPGPInitResponse(encryptedInitResponseJSON, vault)
	if err != nil {
		return nil, fmt.Errorf("ProcessKeyFun|Error while decrypting initResponse: %w", err)
	}
	fmt.Printf("decryptedInitResponse|Response %v", decryptedInitResponse)
	return decryptedInitResponse, nil
}
