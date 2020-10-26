package encryptiongrpc

import (
	"context"
	"log"

	"github.com/shifty21/scone/rsacrypto"
)

const (
	serviceLog           = "EncryptionGRPC.Service|"
	handlerLog           = "EncryptionGRPC.Handler|"
	encryptionServiceLog = "EncryptionGRPC.EncryptionSerice|"
)

//Service interface
type Service interface {
	Encrypt(ctx context.Context, data []byte) ([]byte, error)
	Decrypt(ctx context.Context, data []byte) ([]byte, error)
}

//ServiceImpl struct
type ServiceImpl struct {
	crypto *rsacrypto.Crypto
}

//NewEncryptionhttp for checking flying status in a particular location
func NewEncryptionhttp(crypto *rsacrypto.Crypto) Service {
	// log.Println("Newencryptionhttp")
	return &ServiceImpl{
		crypto: crypto,
	}
}

//Encrypt gets weather from darksky if its not present in cache
func (s *ServiceImpl) Encrypt(ctx context.Context, data []byte) ([]byte, error) {
	encryptedData, _ := s.crypto.EncryptBytes(data)
	decryptedData, _ := s.crypto.DecryptByte(encryptedData)
	log.Printf("%vPlain text :%v\n", serviceLog, string(decryptedData))
	return encryptedData, nil
}

//Decrypt gets weather from darksky if its not present in cache
func (s *ServiceImpl) Decrypt(ctx context.Context, data []byte) ([]byte, error) {
	return s.crypto.DecryptByte(data)
}
