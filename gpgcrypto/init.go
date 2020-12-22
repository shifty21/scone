package gpgcrypto

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shifty21/scone/config"
)

//Crypto struct to store crypto related stuff
type Crypto struct {
	PublicKey  string
	PrivateKey string
	PassPhrase []byte
}

func getKeySet(config *config.GPGCrypto, index int) (*Crypto, error) {
	crypto := &Crypto{}
	log.Printf("Initializing crypto package with public key [%v] and private key [%v]", config.GetGPGCrypto(index).GetPublicKeyPath(), config.GetGPGCrypto(index).GetPrivateKeyPath())
	if config.GetGPGCrypto(index).GetPublicKeyPath() != "" {
		publicKey, err := readKeyFile(config.GetGPGCrypto(index).GetPublicKeyPath())
		if err != nil {
			return nil, err
		}
		crypto.PublicKey = publicKey
	}
	if config.GetGPGCrypto(index).GetPrivateKeyPath() != "" {
		privateKey, err := readKeyFile(config.GetGPGCrypto(index).GetPrivateKeyPath())
		if err != nil {
			return nil, err
		}
		crypto.PrivateKey = privateKey
	}
	crypto.PassPhrase = []byte(config.GetGPGCrypto(index).GetPassPhrase())
	return crypto, nil
}

//InitGPGCrypto initializes variables need to shamir key based vault initialization
func InitGPGCrypto(config *config.GPGCrypto) ([]*Crypto, error) {
	log.Printf("InitGPGCrypt package, with %v keysets", config.GetTotalKeySets())
	var cryptoSet []*Crypto
	for i := 0; i < config.GetTotalKeySets(); i++ {
		keyset, err := getKeySet(config, i)
		if err != nil {
			return nil, fmt.Errorf("Error getting keyset number %v Error: %w", i, err)
		}
		cryptoSet = append(cryptoSet, keyset)
	}
	log.Printf("Total Keysets %v", len(cryptoSet))
	return cryptoSet, nil
}

func readKeyFile(filename string) (string, error) {

	// open private key
	_, err := os.Stat(filename)
	if err != nil {
		return "", fmt.Errorf("readKeyFile|File error %w", err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("GetRSAPublicKey|Error reading public key, path %w", err)

	}
	fmt.Printf("ReadKeyFile|Filename %v Data %v\n", filename, string(data))
	return string(data), nil
}

//Test tests the gpgcrypto package
func Test(crypto []*Crypto) {

	var testdata = make([][]byte, 1)
	testdata[0] = []byte("samuraijack")
	publicKeys := make([]string, len(crypto))
	for _, v := range crypto {
		publicKeys = append(publicKeys, v.PublicKey)
	}
	_, encryptedKeys, err := EncryptShares(testdata, publicKeys)
	if err != nil {
		fmt.Printf("Test|Error while encrypting keys %v", err)
	}
	// encryptedKey := string(encryptedKeys[0])
	encryptedKey := base64.StdEncoding.EncodeToString(encryptedKeys[0])
	// fmt.Printf("Encrypted data %v\n", encryptedKey)
	// var byteencrypted []byte
	// _, err = hex.Decode(encryptedKeys[0], byteencrypted)
	// if err != nil {
	// 	fmt.Printf("Error while decoding encrypted string %v", err)
	// }

	// encryptedKey := "wcDMA+1zkOtUz/mjAQwATwKaTjmXwd5Vhi3+wqGyXDVckwenoeLGPiYeDTWiovI+F3H9mdjeRGlbV24IggJDzC28yKJyHd+smdc0fhPM5ej1JLOnW/oXNU5+CssI4HRESw03daU9L6kD/ODqzE07rKgjnZOC9Z77GEND5rDCCfHhSUX0hqwg/OiDmuzA4yv+1OwNb2Y2ialxStqzJll4hHLw/wetwoUK6widS9SkrvS8T8Nv+vQ+hqwNp0p/tUPjxntVNIw7fhxAWByhr7I+O+NqHuc2ur/aNj8yJF6Lg4DRnMzW3o1XyC/+qtKE9fA9LemwDeJZxtXzfNLzDgiUtsAmZqhYehX5FxZjNyZlfyescwedVV9nXE82KiELCz4DgyaiHShfevn1Akl4jkl+eV161Gymd4r2c7qDGBXLDqeFc/VIQFECBT4B0bualtDqquzDMpGr0SehePajLlH1gx415MqRnDpJCmhv2So2Id/w/Rf+MYnQLRyYdA3tflnjSz2wKgo2Z3xAuo20ZSWD0uAB5A0yo1zRp6ucF4nkjtNXi4vhgNLgjODH4T8d4MHiJ+maguAz5i//LxjXeyosp3vLGNpartpmMjblHzdVc9OQAkokOf9ojgqhxovUZRLp8xZeytf1N3YKrv9oKcrZXaspnlXuygLg+OSKFpMAsKgpu+TA+NAyfQAW4jiem77hhqIA"
	bytebuf, err := DecryptBytes(encryptedKey, crypto[0].PrivateKey, crypto[0].PassPhrase)
	if err != nil {
		fmt.Printf("Error decrypting: %v", err)
	}
	// decodedString, err := base64.StdEncoding.DecodeString(bytebuf.String())
	// if err != nil {
	// 	fmt.Printf("Error decrypting: %v", err)
	// }
	fmt.Printf("Decrypted data %v\n", bytebuf.String())
}
