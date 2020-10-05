package main

import (
	"os"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/encryptionservice"
	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/vaultinitcrypto"
	"github.com/shifty21/scone/vaultinterface"
)

func main() {
	logger.InitLogger()
	config := config.ConfigureAllInterfaces()

	crypto, err := vaultinitcrypto.InitShamirInterface(config.GetCryptoConfig())
	if err != nil {
		logger.Info.Printf("Error while initializing shamirInterface")
		os.Exit(-1)
	}
	vault := vaultinterface.Initialize(config, crypto)
	vault.Run(vaultinitcrypto.EncryptKeyFun, vaultinitcrypto.ProcessKeyFun)
	encryptionservice.Run(config, crypto)
	// vaultinterface.Initialize()
	// vaultinterface.Run(vaultinitshamir.EncryptKeyFun, vaultinitshamir.ProcessKeyFun)

}
