package vaultcryptoinit

import (
	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

var InitResponse *utils.InitResponse

// EncryptInitResponse encryptes all fields one by one and sends back encrypted response
func EncryptInitResponse(initResponse *utils.InitResponse, vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	encryptedInitResponseJSON := &utils.InitResponse{
		Keys:       make([]string, len(initResponse.Keys)),
		KeysBase64: make([]string, len(initResponse.KeysBase64)),
	}
	// //Keys
	for key, plainText := range initResponse.Keys {
		encryptedText, err := vault.Crypto.EncryptText(plainText)
		if err != nil {
			logger.Error.Printf("EncryptInitResponse|Unable to encrypt one of the Keys %d\n", key)
			return nil, err
		}
		encryptedInitResponseJSON.Keys[key] = *encryptedText
	}
	//base64
	for key, plainText := range initResponse.KeysBase64 {
		encryptedText, err := vault.Crypto.EncryptText(plainText)
		if err != nil {
			logger.Error.Printf("EncryptInitResponse|Unable to encrypt one of the KeysBase64 %d\n", key)
			return nil, err
		}
		encryptedInitResponseJSON.KeysBase64[key] = *encryptedText
	}
	//roottoken
	encryptedText, err := vault.Crypto.EncryptText(initResponse.RootToken)
	if err != nil {
		logger.Error.Printf("EncryptInitResponse|Unable to encrypt one of the RootToken")
		return nil, err
	}
	encryptedInitResponseJSON.RootToken = *encryptedText
	return encryptedInitResponseJSON, nil

}

//DecryptInitResponse decrypts all the fields of json
func DecryptInitResponse(encryptedResponse *utils.InitResponse, vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	decryptedInitResponseJSON := &utils.InitResponse{
		Keys:       make([]string, len(encryptedResponse.Keys)),
		KeysBase64: make([]string, len(encryptedResponse.KeysBase64)),
	}
	//Keys
	for key, value := range encryptedResponse.Keys {
		encryptedText, err := vault.Crypto.DecryptText(value)
		if err != nil {
			logger.Error.Printf("DecryptInitResponse|Unable to decrypt one of the Key %d\n", key)
			return nil, err
		}
		decryptedInitResponseJSON.Keys[key] = *encryptedText
	}
	//base64
	for key, value := range encryptedResponse.KeysBase64 {
		encryptedText, err := vault.Crypto.DecryptText(value)
		if err != nil {
			logger.Error.Printf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %d", key)
			return nil, err
		}
		decryptedInitResponseJSON.KeysBase64[key] = *encryptedText
	}
	// roottoken
	encryptedText, err := vault.Crypto.DecryptText(encryptedResponse.RootToken)
	if err != nil {
		logger.Error.Printf("DecryptInitResponse|Unable to decrypt RootToken\n ")
		return nil, err
	}
	decryptedInitResponseJSON.RootToken = *encryptedText
	return decryptedInitResponseJSON, nil

}
