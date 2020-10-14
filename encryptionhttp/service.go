package encryptionhttp

import (
	"context"
	"log"

	"github.com/shifty21/scone/crypto"
)

const (
	logName = "Encryption.Service|"
)

//Service interface
type Service interface {
	EncryptData(ctx context.Context, data string) (*string, error)
	DecryptData(ctx context.Context, data string) (*string, error)
}

//ServiceImpl struct
type ServiceImpl struct {
	crypto *crypto.Crypto
}

//Newencryptionhttp for checking flying status in a particular location
func Newencryptionhttp(crypto *crypto.Crypto) Service {
	return &ServiceImpl{
		crypto: crypto,
	}
}

//EncryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) EncryptData(ctx context.Context, data string) (*string, error) {
	encryptedData, _ := s.crypto.EncryptString(data)
	log.Printf("Encrypted Data %v \n", *encryptedData)
	decryptedData, _ := s.crypto.DecryptString(*encryptedData)
	log.Printf("Decrypted data Data %v \n", *decryptedData)
	return encryptedData, nil
}

//DecryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) DecryptData(ctx context.Context, data string) (*string, error) {
	return s.crypto.DecryptString(data)
}
