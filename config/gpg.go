package config

import (
	"log"

	"github.com/spf13/viper"
)

//KeySet for particular gpg key
type KeySet struct {
	//PublicKeyPath for encryption
	PublicKeyPath string `mapstructure:"public_key_path"`
	//PrivateKeyPath for decryption, will be populated by CAS session
	PrivateKeyPath string `mapstructure:"private_key_path"`
	PassPhrase     string `mapstructure:"pass_phrase"`
}

//GPGCrypto configration
type GPGCrypto struct {
	keySet          []*KeySet
	numberOfKeySets int
}

//GetPublicKeyPath gives vault address
func (k *KeySet) GetPublicKeyPath() string {
	return k.PublicKeyPath
}

//GetPrivateKeyPath gives vault address
func (k *KeySet) GetPrivateKeyPath() string {
	return k.PrivateKeyPath
}

//GetPassPhrase gives vault address
func (k *KeySet) GetPassPhrase() string {
	return k.PassPhrase
}

//GetGPGCrypto gets particular keyset
func (g *GPGCrypto) GetGPGCrypto(i int) *KeySet {
	return g.keySet[i]
}

//GetTotalKeySets gets particular keyset
func (g *GPGCrypto) GetTotalKeySets() int {
	return g.numberOfKeySets
}

//LoadGPGCryptoConfig loads values from viper
func LoadGPGCryptoConfig() *GPGCrypto {
	var keysetslice []*KeySet
	err := viper.UnmarshalKey("gpgcrypto", &keysetslice)
	if err != nil {
		log.Fatalf("Error getting gpgcrypto config %v", err)
	}
	log.Printf("gpgcrypt object %v", len(keysetslice))
	gpgcrypto := &GPGCrypto{keySet: keysetslice, numberOfKeySets: len(keysetslice)}
	return gpgcrypto
}
