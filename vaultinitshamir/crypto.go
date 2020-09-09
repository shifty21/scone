package vaultinitshamir

import (
	"crypto/rsa"

	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/vaultinterface"
)

// EncryptInitResponse encryptes all fields one by one and sends back encrypted response
func (c *Crypto) EncryptInitResponse(initResponse *vaultinterface.InitResponse) (*vaultinterface.InitResponse, error) {
	encryptedInitResponseJSON := &vaultinterface.InitResponse{
		Keys:       make([]string, len(initResponse.Keys)),
		KeysBase64: make([]string, len(initResponse.KeysBase64)),
	}
	// //Keys
	for key, plainText := range initResponse.Keys {
		encryptedText, err := c.EncryptText(plainText)
		if err != nil {
			logger.Error.Printf("EncryptInitResponse|Unable to encrypt one of the Keys %d\n", key)
			return nil, err
		}
		encryptedInitResponseJSON.Keys[key] = *encryptedText
	}
	//base64
	for key, plainText := range initResponse.KeysBase64 {
		encryptedText, err := c.EncryptText(plainText)
		if err != nil {
			logger.Error.Printf("EncryptInitResponse|Unable to encrypt one of the KeysBase64 %d\n", key)
			return nil, err
		}
		encryptedInitResponseJSON.KeysBase64[key] = *encryptedText
	}
	//roottoken
	encryptedText, err := c.EncryptText(initResponse.RootToken)
	if err != nil {
		logger.Error.Printf("EncryptInitResponse|Unable to encrypt one of the RootToken")
		return nil, err
	}
	encryptedInitResponseJSON.RootToken = *encryptedText
	return encryptedInitResponseJSON, nil

}

//EncryptText encrypts given plaintext
func (c *Crypto) EncryptText(plainText string) (*string, error) {
	encryptedBytes, err := rsa.EncryptOAEP(c.HashFun, c.RandomIOReader, c.PublicKey, []byte(plainText), nil)
	if err != nil {
		logger.Error.Printf("encryptText|Error while encrypting initResponse\n")
		return nil, err
	}
	encryptedString := string(encryptedBytes)
	return &encryptedString, nil
}

//DecryptInitResponse decrypts all the fields of json
func (c *Crypto) DecryptInitResponse(encryptedResponse *vaultinterface.InitResponse) (*vaultinterface.InitResponse, error) {
	decryptedInitResponseJSON := &vaultinterface.InitResponse{
		Keys:       make([]string, len(encryptedResponse.Keys)),
		KeysBase64: make([]string, len(encryptedResponse.KeysBase64)),
	}
	//Keys
	for key, value := range encryptedResponse.Keys {
		encryptedText, err := c.DecryptText(value)
		if err != nil {
			logger.Error.Printf("DecryptInitResponse|Unable to decrypt one of the Key %d\n", key)
			return nil, err
		}
		decryptedInitResponseJSON.Keys[key] = *encryptedText
	}
	//base64
	for key, value := range encryptedResponse.KeysBase64 {
		encryptedText, err := c.DecryptText(value)
		if err != nil {
			logger.Error.Printf("DecryptInitResponse|Unable to decrypt one of the KeysBase64 %d", key)
			return nil, err
		}
		decryptedInitResponseJSON.KeysBase64[key] = *encryptedText
	}
	// roottoken
	encryptedText, err := c.DecryptText(encryptedResponse.RootToken)
	if err != nil {
		logger.Error.Printf("DecryptInitResponse|Unable to decrypt RootToken\n ")
		return nil, err
	}
	decryptedInitResponseJSON.RootToken = *encryptedText
	return decryptedInitResponseJSON, nil

}

//DecryptText decryptes given ciphertext
func (c *Crypto) DecryptText(cipherText string) (*string, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, []byte(cipherText), nil)
	if err != nil {
		logger.Error.Printf("decryptText|Error while decrypting init response\n")
		return nil, err
	}
	decryptedString := string(decryptedBytes)
	return &decryptedString, nil
}
