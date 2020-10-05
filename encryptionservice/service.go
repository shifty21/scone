package encryptionservice

import (
	"context"

	"github.com/shifty21/scone/vaultinitcrypto"
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
	crypto *vaultinitcrypto.VaultInitCrypto
}

//NewEncryptionService for checking flying status in a particular location
func NewEncryptionService(crypto *vaultinitcrypto.VaultInitCrypto) Service {
	return &ServiceImpl{
		crypto: crypto,
	}
}

//EncryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) EncryptData(ctx context.Context, data string) (*string, error) {
	return s.crypto.EncryptText(data)
}

//DecryptData gets weather from darksky if its not present in cache
func (s *ServiceImpl) DecryptData(ctx context.Context, data string) (*string, error) {
	return s.crypto.DecryptText(data)
}
