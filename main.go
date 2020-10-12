package main

import (
	"fmt"
	"os"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/encryptionservicegrpc"
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
	// encryptionservice.Run(config, crypto)

	// vaultinterface.Initialize()
	// vaultinterface.Run(vaultinitshamir.EncryptKeyFun, vaultinitshamir.ProcessKeyFun)

	grpcServer, err := encryptionservicegrpc.NewGRPCService(
		encryptionservicegrpc.Config(config),
		encryptionservicegrpc.CryptoService(crypto),
		encryptionservicegrpc.EnableTLS(true),
	)
	if err != nil {
		fmt.Printf("Error while creating grpc service")
		os.Exit(1)
	}
	grpcServer.Run()
}
