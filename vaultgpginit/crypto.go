package vaultgpginit

import (
	"errors"
	"fmt"

	"github.com/shifty21/scone/gpgcrypto"
	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//InitResponse for storing intermediate encrypted response
var InitResponse *utils.InitResponse

//DecryptPGPInitResponse decrypts all the fields of json
func DecryptPGPInitResponse(encryptedResponse *utils.InitResponse, vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	if encryptedResponse == nil {
		return nil, errors.New("No InitResponse found in memory")
	}
	decryptedInitResponseJSON := &utils.InitResponse{
		Keys:               encryptedResponse.Keys,
		KeysBase64:         encryptedResponse.KeysBase64,
		RecoveryKeys:       encryptedResponse.RecoveryKeys,
		RecoveryKeysBase64: make([]string, len(encryptedResponse.RecoveryKeysBase64)),
		RootToken:          encryptedResponse.RootToken,
	}

	// //decrypt base64
	// for key, value := range encryptedResponse.KeysBase64 {
	// 	encryptedText, err := gpgcrypto.DecryptBytes(value, vault.GPGCrypto.PrivateKey[0])
	// 	if err != nil {
	// 		return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %w", err)
	// 	}
	// 	decryptedInitResponseJSON.KeysBase64[key] = encryptedText.String()
	// }
	//decrypt RecoveryKeybase64
	for key, value := range encryptedResponse.RecoveryKeysBase64 {
		encryptedText, err := gpgcrypto.DecryptBytes(value, vault.GPGCrypto.PrivateKey[0])
		if err != nil {
			return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %w", err)
		}
		decryptedInitResponseJSON.RecoveryKeysBase64[key] = encryptedText.String()
	}

	// // roottoken
	// encryptedText, err := gpgcrypto.DecryptBytes(encryptedResponse.RootToken, vault.GPGCrypto.PrivateKey[0])
	// if err != nil {
	// 	return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt RootToken %w ", err)
	// }
	// decryptedInitResponseJSON.RootToken = encryptedText.String()
	return decryptedInitResponseJSON, nil

}
