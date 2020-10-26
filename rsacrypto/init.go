package rsacrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/utils"
)

//Crypto struct to store crypto related stuff
type Crypto struct {
	HashFun        hash.Hash
	RandomIOReader io.Reader
	PublicKey      *rsa.PublicKey
	Label          []byte
	PrivateKey     *rsa.PrivateKey
}

//InitResponse stores initresponse
var InitResponse *utils.InitResponse

//InitCrypto initializes variables need to shamir key based vault initialization
func InitCrypto(config *config.Crypto) (*Crypto, error) {
	log.Printf("Initializing crypto package with public key [%v] and private key [%v]", config.PublicKeyPath(), config.PrivateKeyPath())
	c := &Crypto{}
	c.HashFun = sha512.New()
	c.RandomIOReader = rand.Reader
	err := c.GetRSAPublicKey(config.PublicKeyPath())
	if err != nil {
		return nil, fmt.Errorf("InitCrypto|Error while loading Public key %w", err)
	}
	err = c.GetRSAPrivateKey(config.PrivateKeyPath())
	if err != nil {
		return nil, fmt.Errorf("InitCrypto|Error while loading Private key %w", err)
	}
	return c, nil
}

//GetRSAPublicKey loads PUBLIC KEY based pem block
func (c *Crypto) GetRSAPublicKey(publicKeyPath string) error {
	publicPEM, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("GetRSAPublicKey|Error reading public key, path %w", err)

	}
	pemBlock, _ := pem.Decode([]byte(publicPEM))
	if pemBlock == nil && pemBlock.Type != "PUBLIC KEY" {
		return fmt.Errorf("GetRSAPublicKey|RSA public key is of the wrong type %w", err)
	}
	var pubKey interface{}
	if pubKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes); err != nil {
		return fmt.Errorf("GetRSAPublicKey|Unable to parse RSA public key %w", err)
	}
	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		log.Println("GetRSAPublicKey|RSA public key")
		c.PublicKey = pub
	default:
		log.Println("GetRSAPublicKey|Unrecognized public key")
		return errors.New("GetRSAPublicKey|Unrecognized public key. Handles on Public key with PEM block PUBLIC KEY")
	}

	return nil
}

//GetRSAPrivateKey loads RSA PRIVATE KEY based pem block
func (c *Crypto) GetRSAPrivateKey(privateKeyPath string) error {
	priv, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return fmt.Errorf("GetRSAPrivateKey|Unable to read pvt key %w", err)
	}
	privatePEM, _ := pem.Decode(priv)
	if privatePEM.Type != "RSA PRIVATE KEY" {
		return fmt.Errorf("GetRSAPrivateKey|RSA private key is of the wrong type %w", err)
	}
	c.PrivateKey, err = x509.ParsePKCS1PrivateKey(privatePEM.Bytes)
	if err != nil {
		return fmt.Errorf("GetRSAPrivateKey|Error while parsing private key %w", err)
	}
	return nil

}
