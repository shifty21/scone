package vaultcryptoinit

import (
	"fmt"
	"log"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultgpginit"
	"github.com/shifty21/scone/vaultinterface"
)

//EncryptKeyFun stores keys as required by cas unseal process
var EncryptKeyFun = func(vault *vaultinterface.Vault) error {
	//use a public key to encrypt the response
	// fmt.Printf("InitResponse %v", initResponse)
	if !vault.Opt.EnableGPGEncryption {
		encryptedInitResponse, err := EncryptInitResponse(vault.InitResponse, vault)
		if err != nil {
			return fmt.Errorf("EncryptKeyFun|Error while encrypting initresponse %v", err)
		}
		// fmt.Printf("EncryptedInitResponse %v", encryptedInitResponse)
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
	if vault.Opt.EnableGPGEncryption {
		log.Println("GPGDecryption")
		decryptedInitResponse, err = vaultgpginit.DecryptPGPInitResponse(vault.EncryptedResponse, vault)
		if err != nil {
			return nil, fmt.Errorf("ProcessKeyFun|Error while decrypting initResponse: %w", err)
		}
	} else {
		log.Println("RSADecryption")
		decryptedInitResponse, err = DecryptInitResponse(vault.EncryptedResponse, vault)
		if err != nil {
			return nil, fmt.Errorf("ProcessKeyFun|Error while decrypting initResponse: %w", err)
		}
	}
	vault.DecryptedInitResponse = decryptedInitResponse
	// fmt.Printf("decryptedInitResponse|Response %v", decryptedInitResponse)
	return decryptedInitResponse, nil
}
