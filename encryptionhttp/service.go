package encryptionhttp

import (
	"context"
	"log"

	"github.com/shifty21/scone/rsacrypto"
)

const (
	encryptionserviceLog = "EncryptionHTTP.EncryptionService|"
	handlerLog           = "EncryptionHTTP.Handler|"
	serviceLog           = "Encryption.Service|"
)

//Service interface
type Service interface {
	EncryptData(ctx context.Context, data string) (*string, error)
	DecryptData(ctx context.Context, data string) (*string, error)
}

//ServiceImpl struct
type ServiceImpl struct {
	crypto *rsacrypto.Crypto
}

//NewEncryptionhttp for checking flying status in a particular location
func NewEncryptionhttp(crypto *rsacrypto.Crypto) Service {
	return &ServiceImpl{
		crypto: crypto,
	}
}

//EncryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) EncryptData(ctx context.Context, data string) (*string, error) {
	encryptedData, _ := s.crypto.EncryptString(data)
	log.Printf("%vEncrypted Data %v \n", serviceLog, *encryptedData)
	decryptedData, _ := s.crypto.DecryptString(*encryptedData)
	log.Printf("%vDecrypted data Data %v \n", serviceLog, *decryptedData)
	return encryptedData, nil
}

//DecryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) DecryptData(ctx context.Context, data string) (*string, error) {
	return s.crypto.DecryptString(data)
}
