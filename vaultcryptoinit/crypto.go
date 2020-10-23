package vaultcryptoinit

import (
	"errors"
	"fmt"

	"github.com/shifty21/scone/utils"
	"github.com/shifty21/scone/vaultinterface"
)

//InitResponse for storing intermediate encrypted response
var InitResponse *utils.InitResponse

// EncryptInitResponse encryptes all fields one by one and sends back encrypted response
func EncryptInitResponse(initResponse *utils.InitResponse, vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	encryptedInitResponseJSON := &utils.InitResponse{
		Keys:               make([]string, len(initResponse.Keys)),
		KeysBase64:         make([]string, len(initResponse.KeysBase64)),
		RecoveryKeys:       make([]string, len(initResponse.RecoveryKeys)),
		RecoveryKeysBase64: make([]string, len(initResponse.RecoveryKeysBase64)),
	}
	//Keys
	for key, plainText := range initResponse.Keys {
		encryptedText, err := vault.Crypto.EncryptString(plainText)
		if err != nil {
			return nil, fmt.Errorf("EncryptInitResponse|Unable to encrypt one of the Keys %w", err)
		}
		encryptedInitResponseJSON.Keys[key] = *encryptedText
	}
	//base64
	for key, plainText := range initResponse.KeysBase64 {
		encryptedText, err := vault.Crypto.EncryptString(plainText)
		if err != nil {
			return nil, fmt.Errorf("EncryptInitResponse|Unable to encrypt one of the KeysBase64 %w", err)
		}
		encryptedInitResponseJSON.KeysBase64[key] = *encryptedText
	}

	//Recovery Keys
	for key, plainText := range initResponse.RecoveryKeys {
		encryptedText, err := vault.Crypto.EncryptString(plainText)
		if err != nil {
			return nil, fmt.Errorf("EncryptInitResponse|Unable to encrypt one of the Keys %w", err)
		}
		encryptedInitResponseJSON.RecoveryKeys[key] = *encryptedText
	}
	//recovery key base64
	for key, plainText := range initResponse.RecoveryKeysBase64 {
		encryptedText, err := vault.Crypto.EncryptString(plainText)
		if err != nil {
			return nil, fmt.Errorf("EncryptInitResponse|Unable to encrypt one of the KeysBase64 %w", err)
		}
		encryptedInitResponseJSON.RecoveryKeysBase64[key] = *encryptedText
	}
	//root token
	encryptedText, err := vault.Crypto.EncryptString(initResponse.RootToken)
	if err != nil {
		return nil, fmt.Errorf("EncryptInitResponse|Unable to encrypt one of the RootToken %w", err)
	}
	encryptedInitResponseJSON.RootToken = *encryptedText
	return encryptedInitResponseJSON, nil

}

//DecryptInitResponse decrypts all the fields of json
func DecryptInitResponse(encryptedResponse *utils.InitResponse, vault *vaultinterface.Vault) (*utils.InitResponse, error) {
	if encryptedResponse == nil {
		return nil, errors.New("No InitResponse found in memory")
	}
	decryptedInitResponseJSON := &utils.InitResponse{
		Keys:               make([]string, len(encryptedResponse.Keys)),
		KeysBase64:         make([]string, len(encryptedResponse.KeysBase64)),
		RecoveryKeys:       make([]string, len(encryptedResponse.RecoveryKeys)),
		RecoveryKeysBase64: make([]string, len(encryptedResponse.RecoveryKeysBase64)),
	}
	//Keys
	for key, value := range encryptedResponse.Keys {
		encryptedText, err := vault.Crypto.DecryptString(value)
		if err != nil {
			return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the Key %w", err)
		}
		decryptedInitResponseJSON.Keys[key] = *encryptedText
	}
	//base64
	for key, value := range encryptedResponse.KeysBase64 {
		encryptedText, err := vault.Crypto.DecryptString(value)
		if err != nil {
			return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %w", err)
		}
		decryptedInitResponseJSON.KeysBase64[key] = *encryptedText
	}
	//RecoveryKeys
	for key, value := range encryptedResponse.RecoveryKeys {
		encryptedText, err := vault.Crypto.DecryptString(value)
		if err != nil {
			return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the Key %w", err)
		}
		decryptedInitResponseJSON.RecoveryKeys[key] = *encryptedText
	}
	//RecoveryKeybase64
	for key, value := range encryptedResponse.RecoveryKeysBase64 {
		encryptedText, err := vault.Crypto.DecryptString(value)
		if err != nil {
			return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %w", err)
		}
		decryptedInitResponseJSON.RecoveryKeysBase64[key] = *encryptedText
	}

	// roottoken
	encryptedText, err := vault.Crypto.DecryptString(encryptedResponse.RootToken)
	if err != nil {
		return nil, fmt.Errorf("DecryptInitResponse|Unable to decrypt RootToken %w ", err)
	}
	decryptedInitResponseJSON.RootToken = *encryptedText
	return decryptedInitResponseJSON, nil

}
