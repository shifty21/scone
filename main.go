package main

import (
	"log"
	"os"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/encryptiongrpc"
	"github.com/shifty21/scone/encryptionhttp"
	"github.com/shifty21/scone/vaultcryptoinit"
	"github.com/shifty21/scone/vaultinit"
	"github.com/shifty21/scone/vaultinterface"
)

func main() {
	config := config.ConfigureAllInterfaces()
	crypto, err := crypto.InitCrypto(config.GetCryptoConfig())
	if err != nil {
		log.Println("Error while initializing crypto module, Exiting")
		os.Exit(1)
	}
	forever := make(chan struct{})
	initType := "http"
	switch initType {
	case "vanilla":

		vault := vaultinterface.Initialize(config, crypto)
		go vault.Run(vaultinit.EncryptKeyFun, vaultinit.ProcessKeyFun)
		<-forever
	case "cas":
		vault := vaultinterface.Initialize(config, crypto)
		go vault.Run(vaultcryptoinit.EncryptKeyFun, vaultcryptoinit.ProcessKeyFun)
		<-forever
	case "http":
		encryptionhttp.Run(config, crypto)
	default:
		grpcServer, err := encryptiongrpc.NewGRPCService(
			encryptiongrpc.Config(config),
			encryptiongrpc.CryptoService(crypto),
			encryptiongrpc.EnableTLS(true),
		)
		if err != nil {
			log.Printf("Error while creating grpc service")
			os.Exit(1)
		}
		grpcServer.Run()
	}

}
