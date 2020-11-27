package vaultsconeinit

import (
	"fmt"
	"log"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultautoinit"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(vault *vaultinterface.Vault) error {
	//use a public key to encrypt the response
	// fmt.Printf("InitResponse %v", initResponse)
	if !vault.Opt.EnableGPGEncryption {
		log.Println("EncryptKeyFun|Encrypting unseal keys and the root token...")
		encryptedInitResponse, err := EncryptInitResponse(vault.InitResponse, vault)
		if err != nil {
			return fmt.Errorf("EncryptKeyFun|Error while encrypting initresponse %v", err)
		}
		vault.EncryptedResponse = encryptedInitResponse
	} else {
		vault.EncryptedResponse = vault.InitResponse
	}
	return nil
}

//ProcessKeyFun stores keys as required by cas unseal process
var ProcessKeyFun = func(vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	//CAS session can contain private keys, we can use public key to encrypt and upon
	//verification we can use the provided privated key by CAS to decrypt the unseal keys
	var err error
	var decryptedInitResponse *utils.InitResponse
	if !vault.Opt.EnableGPGEncryption {
		log.Println("RSADecryption")
		decryptedInitResponse, err = DecryptInitResponse(vault.EncryptedResponse, vault)
		if err != nil {
			return nil, fmt.Errorf("ProcessKeyFun|Error while decrypting initResponse: %w", err)
		}
	} else {
		log.Printf("GPGDecryption %v", vault.EncryptedResponse)
		decryptedInitResponse, err = vaultautoinit.DecryptPGPInitResponse(vault.EncryptedResponse, vault)
		if err != nil {
			return nil, fmt.Errorf("ProcessKeyFun|Error while decrypting initResponse: %w", err)
		}
	}
	vault.DecryptedInitResponse = decryptedInitResponse
	// fmt.Printf("decryptedInitResponse|Response %v", decryptedInitResponse)
	return decryptedInitResponse, nil
}
