package vaultautoinit

import (
	"fmt"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(vault *vaultinterface.Vault) error {
	vault.EncryptedResponse = vault.InitResponse
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	var err error
	var decryptedInitResponse *utils.InitResponse
	if vault.Opt.EnableGPGEncryption {
		decryptedInitResponse, err = DecryptPGPInitResponse(vault.EncryptedResponse, vault)
		if err != nil {
			return nil, fmt.Errorf("ProcessKeyFun|Error while decrypting initResponse: %w", err)
		}
		vault.DecryptedInitResponse = decryptedInitResponse
	} else {
		vault.DecryptedInitResponse = vault.InitResponse
	}

	return decryptedInitResponse, nil
}
