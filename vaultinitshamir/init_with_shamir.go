package vaultinitshamir

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"hash"
	"io"
	"io/ioutil"

	"github.com/shifty21/scone/logger"
)

//VaultInitShamir struct to store shamir encryption related stuff
type VaultInitShamir struct {
	HashFun        hash.Hash
	RandomIOReader io.Reader
	PublicKey      *rsa.PublicKey
	Label          []byte
	PrivateKey     *rsa.PrivateKey
}

//InitShamirInterface initializes variables need to shamir key based vault initialization
func InitShamirInterface() (*VaultInitShamir, error) {
	v := &VaultInitShamir{}
	v.HashFun = sha512.New()
	v.RandomIOReader = rand.Reader
	err := v.GetRSAPublicKey("experiments/keypairs/mykey.pub")
	if err != nil {
		logger.Error.Printf("InitShamirInterface|Error while loading Public key\n")
		return nil, err
	}
	err = v.GetRSAPrivateKey("experiments/keypairs/mykey.pem")
	if err != nil {
		logger.Error.Printf("InitShamirInterface|Error while loading Private key %v\n", err)
		return nil, err
	}
	return v, nil
}

//GetRSAPublicKey loads PUBLIC KEY based pem block
func (v *VaultInitShamir) GetRSAPublicKey(publicKeyPath string) error {
	publicPEM, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		logger.Error.Printf("GetRSAPublicKey|Error reading public key, path %v\n", publicKeyPath)
		return err
	}
	pemBlock, _ := pem.Decode([]byte(publicPEM))
	if pemBlock == nil && pemBlock.Type != "PUBLIC KEY" {
		logger.Error.Printf("GetRSAPublicKey|RSA public key is of the wrong type %v\n", err)
		return err
	}
	var pubKey interface{}
	if pubKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes); err != nil {
		logger.Error.Printf("GetRSAPublicKey|Unable to parse RSA public key%v\n", err)
		return err
	}
	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		logger.Info.Printf("GetRSAPublicKey|RSA public key")
		v.PublicKey = pub
	default:
		logger.Error.Printf("GetRSAPublicKey|Unrecognized public key")
		return errors.New("GetRSAPublicKey|Unrecognized public key. Handles on Public key with PEM block PUBLIC KEY")
	}

	return nil
}

//GetRSAPrivateKey loads RSA PRIVATE KEY based pem block
func (v *VaultInitShamir) GetRSAPrivateKey(privateKeyPath string) error {
	priv, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		logger.Error.Printf("GetRSAPrivateKey|Unable to read pvt key %v\n", err)
		return err
	}
	privatePEM, _ := pem.Decode(priv)
	if privatePEM.Type != "RSA PRIVATE KEY" {
		logger.Error.Printf("GetRSAPrivateKey|RSA private key is of the wrong type %v\n", err)
		return errors.New("GetRSAPrivateKey|PEM block of type RSA PRIVATE KEY not found")
	}
	v.PrivateKey, err = x509.ParsePKCS1PrivateKey(privatePEM.Bytes)
	if err != nil {
		logger.Error.Printf("GetRSAPrivateKey|Error while parsing private key")
	}
	return nil

}
