package encryptionservicegrpc

import (
	"context"
	"fmt"

	"github.com/shifty21/scone/crypto"
)

//Service interface
type Service interface {
	Encrypt(ctx context.Context, data []byte) ([]byte, error)
	Decrypt(ctx context.Context, data []byte) ([]byte, error)
}

//ServiceImpl struct
type ServiceImpl struct {
	crypto *crypto.Crypto
}

//NewEncryptionService for checking flying status in a particular location
func NewEncryptionService(crypto *crypto.Crypto) Service {
	// fmt.Println("NewEncryptionService")
	return &ServiceImpl{
		crypto: crypto,
	}
}

//Encrypt gets weather from darksky if its not present in cache
func (s *ServiceImpl) Encrypt(ctx context.Context, data []byte) ([]byte, error) {
	encryptedData, _ := s.crypto.EncryptBytes(data)
	decryptedData, _ := s.crypto.DecryptByte(encryptedData)
	fmt.Printf("EncryptHandler|Plain text :%v\n", string(decryptedData))
	return encryptedData, nil
}

//Decrypt gets weather from darksky if its not present in cache
func (s *ServiceImpl) Decrypt(ctx context.Context, data []byte) ([]byte, error) {
	return s.crypto.DecryptByte(data)
}
