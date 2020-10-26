package rsacrypto

import (
	"crypto/rsa"
	"fmt"
)

//EncryptString encrypts given plaintext
func (c *Crypto) EncryptString(plainText string) (*string, error) {
	encryptedBytes, err := rsa.EncryptOAEP(c.HashFun, c.RandomIOReader, c.PublicKey, []byte(plainText), nil)
	if err != nil {
		return nil, fmt.Errorf("EncryptString|Error while encrypting %w", err)
	}
	encryptedString := string(encryptedBytes)
	return &encryptedString, nil
}

//EncryptBytes encrypts given plaintext
func (c *Crypto) EncryptBytes(plainText []byte) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(c.HashFun, c.RandomIOReader, c.PublicKey, plainText, nil)
	if err != nil {
		return nil, fmt.Errorf("EncryptBytes|Error while encrypting %w", err)
	}
	// encryptedString := string(encryptedBytes)
	return encryptedBytes, nil
}

//DecryptString decryptes given ciphertext
func (c *Crypto) DecryptString(cipherText string) (*string, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, []byte(cipherText), nil)
	if err != nil {
		return nil, fmt.Errorf("DecryptString|Error while decrypting %w", err)
	}
	decryptedString := string(decryptedBytes)
	return &decryptedString, nil
}

//DecryptByte decryptes given ciphertext
func (c *Crypto) DecryptByte(cipherText []byte) ([]byte, error) {
	decryptedBytes, err := rsa.DecryptOAEP(c.HashFun, c.RandomIOReader, c.PrivateKey, cipherText, nil)
	if err != nil {
		return nil, fmt.Errorf("DecryptByte|Error while decrypting %w", err)
	}
	return decryptedBytes, nil
}
