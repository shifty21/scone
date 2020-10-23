package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/encryptiongrpc"
	"github.com/shifty21/scone/encryptionhttp"
	"github.com/shifty21/scone/vaultcryptoinit"
	"github.com/shifty21/scone/vaultgpginit"
	"github.com/shifty21/scone/vaultinit"
	"github.com/shifty21/scone/vaultinterface"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Println("Please specify one of vanilla, cas, grpc")
		os.Exit(0)
	}
	if args[0] == "--version" || args[0] == "version" {
		fmt.Println("Vault Initializer Version - 0.01")
		os.Exit(0)
	}
	config := config.ConfigureAllInterfaces()
	crypto, err := crypto.InitCrypto(config.GetCryptoConfig())
	if err != nil {
		log.Printf("Error while initializing crypto module, Exiting %v", err)
		os.Exit(1)
	}
	forever := make(chan struct{})

	initType := args[0]
	switch initType {
	case "vanilla":
		log.Println("Starting vanilla vault initialization")
		vault := vaultinterface.Initialize(config, crypto)
		go vault.Run(vaultinit.EncryptKeyFun, vaultinit.ProcessKeyFun)
		<-forever
	case "scone":
		log.Println("Starting scone based vault initialization")
		vault := vaultinterface.Initialize(config, crypto)
		go vault.Run(vaultcryptoinit.EncryptKeyFun, vaultcryptoinit.ProcessKeyFun)
		<-forever
	case "pgp":
		log.Println("Starting pgp based vault initialization")
		// crypto, err := gpgcrypto.InitGPGCrypto(config.GetGPGCryptoConfig())
		// if err != nil {
		// 	fmt.Printf("Error while initializing gpgcrypto %v: ", err)
		// }
		// gpgcrypto.Test(crypto)
		vault := vaultinterface.Initialize(config, crypto)
		vault.AutoInitilization = true
		go vault.Run(vaultgpginit.EncryptKeyFun, vaultgpginit.ProcessKeyFun)
		<-forever
	case "auto":
		log.Println("Starting pgp based vault initialization")
		vault := vaultinterface.Initialize(config, crypto)
		vault.AutoInitilization = true
		go vault.Run(vaultcryptoinit.EncryptKeyFun, vaultcryptoinit.ProcessKeyFun)
		<-forever
	case "http":
		encryptionhttp.Run(config, crypto)
	case "grpc":
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
	default:
		log.Println("Please specify one of vanilla, cas, grpc")
		os.Exit(0)
	}

}
