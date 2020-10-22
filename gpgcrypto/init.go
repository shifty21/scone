package gpgcrypto

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/utils"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

//Crypto struct to store crypto related stuff
type Crypto struct {
	HashFun        hash.Hash
	RandomIOReader io.Reader
	PublicKey      *packet.PublicKey
	Label          []byte
	PrivateKey     *packet.PrivateKey
}

//InitResponse stores initresponse
var InitResponse *utils.InitResponse

//InitCrypto initializes variables need to shamir key based vault initialization
func InitCrypto(config *config.Crypto) (*Crypto, error) {
	log.Printf("Initializing crypto package with public key [%v] and private key [%v]", config.PublicKeyPath(), config.PrivateKeyPath())
	c := &Crypto{}
	c.HashFun = sha512.New()
	c.RandomIOReader = rand.Reader

	return c, nil
}

func decodePrivateKey(filename string) *packet.PrivateKey {

	// open ascii armored private key
	in, err := os.Open(filename)
	fmt.Printf("Error opening private key: %s", err)
	defer in.Close()

	block, err := armor.Decode(in)
	fmt.Printf("Error decoding OpenPGP Armor: %s", err)

	if block.Type != openpgp.PrivateKeyType {
		fmt.Printf("Invalid private key file Error decoding private key")
	}

	reader := packet.NewReader(block.Body)
	pkt, err := reader.Next()
	fmt.Printf("Error reading private key")

	key, ok := pkt.(*packet.PrivateKey)
	if !ok {
		fmt.Errorf("private key Error parsing private key")
	}
	return key
}

func decodePublicKey(filename string) *packet.PublicKey {

	// open ascii armored public key
	in, err := os.Open(filename)
	fmt.Printf("Error opening public key: %s", err)
	defer in.Close()

	block, err := armor.Decode(in)
	fmt.Printf("Error decoding OpenPGP Armor: %s", err)

	if block.Type != openpgp.PublicKeyType {
		fmt.Errorf("Invalid private key file Error decoding private key")
	}

	reader := packet.NewReader(block.Body)
	pkt, err := reader.Next()
	fmt.Printf("Error reading private key")

	key, ok := pkt.(*packet.PublicKey)
	if !ok {
		fmt.Errorf("Error parsing public key")
	}
	return key
}
