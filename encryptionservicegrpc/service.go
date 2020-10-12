package encryptionservicegrpc

import (
	"context"
	"fmt"

	"github.com/shifty21/scone/crypto"
)

//Service interface
type Service interface {
	EncryptData(ctx context.Context, data string) (*string, error)
	EncryptDataByte(ctx context.Context, data []byte) ([]byte, error)
	DecryptData(ctx context.Context, data string) (*string, error)
	DecryptDataByte(ctx context.Context, data []byte) ([]byte, error)
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

//EncryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) EncryptData(ctx context.Context, data string) (*string, error) {
	encryptedData, _ := s.crypto.EncryptText(data)
	fmt.Printf("Service.EncryptData|Encrypted Data %v \n", string(*encryptedData))
	decryptedData, _ := s.crypto.DecryptText(*encryptedData)
	fmt.Printf("Service.EncryptData|Decrypted data match %v \n", data == *decryptedData)
	return encryptedData, nil
}

//EncryptDataByte gets weather from darksky if its not present in cache
func (s *ServiceImpl) EncryptDataByte(ctx context.Context, data []byte) ([]byte, error) {
	encryptedData, _ := s.crypto.EncryptTextBytes(data)
	decryptedData, _ := s.crypto.DecryptTextByte(encryptedData)
	fmt.Printf("EncryptHandler|Plain text :%v\n", string(decryptedData))
	return encryptedData, nil
}

//DecryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) DecryptData(ctx context.Context, data string) (*string, error) {
	return s.crypto.DecryptText(data)
}

//DecryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) DecryptDataByte(ctx context.Context, data []byte) ([]byte, error) {
	return s.crypto.DecryptTextByte(data)
}
