package gpgcrypto

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/keybase/go-crypto/openpgp"
	"github.com/keybase/go-crypto/openpgp/packet"
)

// EncryptShares takes an ordered set of byte slices to encrypt and the
// corresponding base64-encoded public keys to encrypt them with, encrypts each
// byte slice with the corresponding public key.
//
// Note: There is no corresponding test function; this functionality is
// thoroughly tested in the init and rekey command unit tests
func EncryptShares(input [][]byte, pgpKeys []string) ([]string, [][]byte, error) {
	// fmt.Printf("EncryptShares| input %v, pgpKeys %v\n", input, pgpKeys)
	if len(input) != len(pgpKeys) {
		return nil, nil, fmt.Errorf("EncryptShares|Mismatch between number items to encrypt and number of PGP keys")
	}
	encryptedShares := make([][]byte, 0, len(pgpKeys))
	entities, err := GetEntities(pgpKeys)
	if err != nil {
		return nil, nil, err
	}
	for i, entity := range entities {
		ctBuf := bytes.NewBuffer(nil)
		pt, err := openpgp.Encrypt(ctBuf, []*openpgp.Entity{entity}, nil, nil, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("EncryptShares|Error setting up encryption for PGP message: %w", err)
		}
		_, err = pt.Write(input[i])
		if err != nil {
			return nil, nil, fmt.Errorf("EncryptShares|Error encrypting PGP message: %w", err)
		}
		pt.Close()
		encryptedShares = append(encryptedShares, ctBuf.Bytes())
	}

	fingerprints, err := GetFingerprints(nil, entities)
	if err != nil {
		return nil, nil, err
	}

	return fingerprints, encryptedShares, nil
}

// GetFingerprints takes in a list of openpgp Entities and returns the
// fingerprints. If entities is nil, it will instead parse both entities and
// fingerprints from the pgpKeys string slice.
func GetFingerprints(pgpKeys []string, entities []*openpgp.Entity) ([]string, error) {
	if entities == nil {
		var err error
		entities, err = GetEntities(pgpKeys)

		if err != nil {
			return nil, err
		}
	}
	ret := make([]string, 0, len(entities))
	for _, entity := range entities {
		ret = append(ret, fmt.Sprintf("%x", entity.PrimaryKey.Fingerprint))
	}
	return ret, nil
}

// GetEntities takes in a string array of base64-encoded PGP keys and returns
// the openpgp Entities
func GetEntities(pgpKeys []string) ([]*openpgp.Entity, error) {
	ret := make([]*openpgp.Entity, 0, len(pgpKeys))
	for _, keystring := range pgpKeys {
		data, err := base64.StdEncoding.DecodeString(keystring)
		if err != nil {
			return nil, fmt.Errorf("GetEntities|Error decoding given PGP key: %w", err)
		}
		entity, err := openpgp.ReadEntity(packet.NewReader(bytes.NewBuffer(data)))
		if err != nil {
			return nil, fmt.Errorf("GetEntities|Error parsing given PGP key: %w", err)
		}
		ret = append(ret, entity)
	}
	return ret, nil
}

// DecryptBytes takes in base64-encoded encrypted bytes and the base64-encoded
// private key and decrypts it. A bytes.Buffer is returned to allow the caller
// to do useful thing with it (get it as a []byte, get it as a string, use it
// as an io.Reader, etc), and also because this function doesn't know if what
// comes out is binary data or a string, so let the caller decide.
func DecryptBytes(encodedCrypt, privKey string, passphrasebyte []byte) (*bytes.Buffer, error) {
	// fmt.Printf("DecryptBytes|encodedCrypt %v \n privKey %v\n", encodedCrypt, privKey)
	privKeyBytes, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return nil, fmt.Errorf("DecryptBytes|Error decoding base64 private key: %w", err)
	}

	cryptBytes, err := base64.StdEncoding.DecodeString(encodedCrypt)
	if err != nil {
		return nil, fmt.Errorf("DecryptBytes|Error decoding base64 crypted bytes: %w", err)
	}
	// fmt.Printf("CryptBytes %v", string(cryptBytes))
	entity, err := openpgp.ReadEntity(packet.NewReader(bytes.NewBuffer(privKeyBytes)))
	if err != nil {
		return nil, fmt.Errorf("DecryptBytes|Error parsing private key: %w", err)
	}
	// fmt.Printf("Entities %v\n", entity.PrimaryKey)
	entityList := &openpgp.EntityList{entity}
	log.Printf("Decrypting private key using passphrase")
	err = entity.PrivateKey.Decrypt(passphrasebyte)
	if err != nil {
		return nil, fmt.Errorf("Error decrypting privatekey: %w", err)
	}
	for _, subkey := range entity.Subkeys {
		err := subkey.PrivateKey.Decrypt(passphrasebyte)
		if err != nil {
			return nil, fmt.Errorf("DecryptBytes|Error decrypting sub privatekey: %w", err)
		}
	}

	log.Printf("DecryptBytes|Finished decrypting private key using passphrase")
	md, err := openpgp.ReadMessage(bytes.NewBuffer(cryptBytes), entityList, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("DecryptBytes|Error decrypting the messages: %w", err)
	}

	ptBuf := bytes.NewBuffer(nil)
	ptBuf.ReadFrom(md.UnverifiedBody)

	return ptBuf, nil
}
