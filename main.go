package main

import (
	"os"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/encryptionservice"
	"github.com/shifty21/scone/logger"
)

func main() {
	logger.InitLogger()
	config := config.ConfigureAllInterfaces()
	crypto, err := crypto.InitCrypto(config.GetCryptoConfig())
	if err != nil {
		logger.Error.Printf("Error while initializing crypto module, Exiting")
		os.Exit(1)
	}
	// vault := vaultinterface.Initialize(config, crypto)
	// go vault.Run(vaultinit.EncryptKeyFun, vaultinit.ProcessKeyFun)
	// go vault.Run(vaultcryptoinit.EncryptKeyFun, vaultcryptoinit.ProcessKeyFun)
	encryptionservice.Run(config, crypto)
	// vaultinterface.Initialize()
	// vaultinterface.Run(vaultinitshamir.EncryptKeyFun, vaultinitshamir.ProcessKeyFun)

}
