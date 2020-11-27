package vaultautoinit

import (
	"fmt"

	"github.com/shifty21/scone/gpgcrypto"
	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//DecryptPGPInitResponse decrypts all the fields of json
func DecryptPGPInitResponse(encryptedResponse *utils.InitResponse, vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	if encryptedResponse == nil {
		return nil, fmt.Errorf("DecryptPGPInitResponse|No InitResponse found in memory")
	}
	decryptedInitResponseJSON := &utils.InitResponse{
		Keys:               encryptedResponse.Keys,
		KeysBase64:         encryptedResponse.KeysBase64,
		RecoveryKeys:       make([]string, len(encryptedResponse.RecoveryKeys)),
		RecoveryKeysBase64: make([]string, len(encryptedResponse.RecoveryKeysBase64)),
		RootToken:          encryptedResponse.RootToken,
	}

	//Decrypt KeysBase64
	if !vault.Opt.IsAutoInitilization {

		for key, value := range encryptedResponse.KeysBase64 {
			encryptedText, err := gpgcrypto.DecryptBytes(value, vault.Opt.GPGCrypto[key].PrivateKey, vault.Opt.GPGCrypto[key].PassPhrase)
			if err != nil {
				return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %w value %v", err, value)
			}
			decryptedInitResponseJSON.KeysBase64[key] = encryptedText.String()
		}
	} else {
		//Decrypt RecoveryKeybase64
		for key, value := range encryptedResponse.RecoveryKeysBase64 {
			encryptedText, err := gpgcrypto.DecryptBytes(value, vault.Opt.GPGCrypto[key].PrivateKey, vault.Opt.GPGCrypto[key].PassPhrase)
			if err != nil {
				return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %w", err)
			}
			decryptedInitResponseJSON.RecoveryKeysBase64[key] = encryptedText.String()
		}
	}
	//Decrypt RootToken
	encryptedText, err := gpgcrypto.DecryptBytes(encryptedResponse.RootToken, vault.Opt.GPGCrypto[0].PrivateKey, vault.Opt.GPGCrypto[0].PassPhrase)
	if err != nil {
		return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt RootToken %w ", err)
	}
	decryptedInitResponseJSON.RootToken = encryptedText.String()
	return decryptedInitResponseJSON, nil

}
