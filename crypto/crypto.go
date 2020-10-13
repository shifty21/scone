package crypto

import (
	"crypto/rsa"

	"github.com/shifty21/scone/logger"
)

//EncryptString encrypts given plaintext
func (c *Crypto) EncryptString(plainText string) (*string, error) {
	encryptedBytes, err := rsa.EncryptOAEP(c.HashFun, c.RandomIOReader, c.PublicKey, []byte(plainText), nil)
	if err != nil {
		logger.Error.Printf("EncryptString|Error while encrypting %v\n", err)
		return nil, err
	}
	encryptedString := string(encryptedBytes)
	return &encryptedString, nil
}

//EncryptBytes encrypts given plaintext
func (c *Crypto) EncryptBytes(plainText []byte) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(c.HashFun, c.RandomIOReader, c.PublicKey, plainText, nil)
	if err != nil {
		logger.Error.Printf("EncryptBytes|Error while encrypting %v\n", err)
		return nil, err
	}
	// encryptedString := string(encryptedBytes)
	return encryptedBytes, nil
}

//DecryptString decryptes given ciphertext
func (c *Crypto) DecryptString(cipherText string) (*string, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, []byte(cipherText), nil)
	if err != nil {
		logger.Error.Printf("DecryptString|Error while decrypting %v\n", err)
		return nil, err
	}
	decryptedString := string(decryptedBytes)
	return &decryptedString, nil
}

//DecryptByte decryptes given ciphertext
func (c *Crypto) DecryptByte(cipherText []byte) ([]byte, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, cipherText, nil)
	if err != nil {
		logger.Error.Printf("DecryptByte|Error while decrypting %v\n", err)
		return nil, err
	}
	return decryptedBytes, nil
}
