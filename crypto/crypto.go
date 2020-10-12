package crypto

import (
	"crypto/rsa"

	"github.com/shifty21/scone/logger"
)

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

//EncryptTextBytes encrypts given plaintext
func (c *Crypto) EncryptTextBytes(plainText []byte) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(c.HashFun, c.RandomIOReader, c.PublicKey, plainText, nil)
	if err != nil {
		logger.Error.Printf("encryptText|Error while encrypting initResponse\n")
		return nil, err
	}
	// encryptedString := string(encryptedBytes)
	return encryptedBytes, nil
}

//DecryptText decryptes given ciphertext
func (c *Crypto) DecryptText(cipherText string) (*string, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, []byte(cipherText), nil)
	if err != nil {
		logger.Error.Printf("decryptText|Error while decrypting init response %v\n", err)
		return nil, err
	}
	decryptedString := string(decryptedBytes)
	return &decryptedString, nil
}

//DecryptTextByte decryptes given ciphertext
func (c *Crypto) DecryptTextByte(cipherText []byte) ([]byte, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, cipherText, nil)
	if err != nil {
		logger.Error.Printf("decryptText|Error while decrypting init response %v\n", err)
		return nil, err
	}
	return decryptedBytes, nil
}
