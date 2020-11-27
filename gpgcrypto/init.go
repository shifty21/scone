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
	PublicKey  []string
	PrivateKey []string
	PassPhrase []byte
}

//InitGPGCrypto initializes variables need to shamir key based vault initialization
func InitGPGCrypto(config *config.GPGCrypto) (*Crypto, error) {
	log.Println("InitGPGCrypt package")
	crypto := &Crypto{
		PublicKey:  make([]string, 1),
		PrivateKey: make([]string, 1),
	}
	log.Printf("Initializing crypto package with public key [%v] and private key [%v]", config.PublicKeyPath(), config.PrivateKeyPath())
	if config.PublicKeyPath() != "" {
		publicKey, err := readKeyFile(config.PublicKeyPath())
		if err != nil {
			return nil, err
		}
		crypto.PublicKey[0] = publicKey
	}
	if config.PrivateKeyPath() != "" {
		privateKey, err := readKeyFile(config.PrivateKeyPath())
		if err != nil {
			return nil, err
		}
		crypto.PrivateKey[0] = privateKey
	}
	crypto.PassPhrase = []byte(config.PassPhrase())
	return crypto, nil
}

func readKeyFile(filename string) (string, error) {
	// fmt.Printf("ReadKeyFile| filename %v\n", filename)
	// open private key
	_, err := os.Stat(filename)
	if err != nil {
		return "", fmt.Errorf("readKeyFile|File error %w", err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("GetRSAPublicKey|Error reading public key, path %w", err)

	}
	// fmt.Printf("ReadKeyFile| filename %v bytes read %v\n", filename, string(data))
	return string(data), nil
}

//Test tests the gpgcrypto package
func Test(crypto *Crypto) {

	var testdata = make([][]byte, 1)
	testdata[0] = []byte("samuraijack")

	_, encryptedKeys, err := EncryptShares(testdata, crypto.PublicKey)
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
	bytebuf, err := DecryptBytes(encryptedKey, crypto.PrivateKey[0], crypto.PassPhrase)
	if err != nil {
		fmt.Printf("Error decrypting: %v", err)
	}
	// decodedString, err := base64.StdEncoding.DecodeString(bytebuf.String())
	// if err != nil {
	// 	fmt.Printf("Error decrypting: %v", err)
	// }
	fmt.Printf("Decrypted data %v\n", bytebuf.String())
}
