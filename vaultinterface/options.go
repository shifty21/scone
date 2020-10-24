package vaultinterface

import (
	"log"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/gpgcrypto"
)

//Options for vault interface
type Options struct {
	VaultConfig        *config.Vault
	CASConfig          *config.VaultCAS
	InitializationType string
	SconeCryptoConfig  *config.Crypto
	SconeCrypto        *crypto.Crypto

	EnableGPGEncryption bool
	GPGCryptoConfig     *config.GPGCrypto
	GPGCrypto           *gpgcrypto.Crypto

	IsAutoInitilization bool
}

//Option function interface
type Option func(*Options)

//EnableAutoInitialization enable in case one needs to initialize vault via recovery keys
func EnableAutoInitialization() Option {
	return func(o *Options) {
		log.Println("Enable autoinitialization")
		o.IsAutoInitilization = true
	}
}

//EnableGPGEncryption enables pgp encryption for providing pgp keys to vault for intialization
//along with decryption of init response from vault
func EnableGPGEncryption() Option {
	return func(o *Options) {
		log.Println("Enabling GPGEncryption")
		o.EnableGPGEncryption = true
	}
}

//SetGPGCryptoConfig sets gpg crypto that provides gpg based crypto api
func SetGPGCryptoConfig(gpgCryptConfig *config.GPGCrypto) Option {
	return func(o *Options) {
		log.Println("Setting GPGEncryption")
		o.GPGCryptoConfig = gpgCryptConfig
	}
}

//SetCASConfig sets config that will verify the clients CAS config
func SetCASConfig(casConfig *config.VaultCAS) Option {
	return func(o *Options) {
		log.Println("Setting CASConfig")
		o.CASConfig = casConfig
	}
}

//SetVaultInitialiationType sets config that will verify the clients CAS config
func SetVaultInitialiationType(initialziationtype string) Option {
	return func(o *Options) {
		log.Println("Setting VaultInit Type")
		o.InitializationType = initialziationtype
	}
}

//SconeCryptoConfig sets config that will verify the clients CAS config
func SconeCryptoConfig(config *config.Crypto) Option {
	return func(o *Options) {
		log.Println("Setting SconeCryptoConfig")
		o.SconeCryptoConfig = config
	}
}
