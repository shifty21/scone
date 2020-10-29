package main

import (
	"os"

	"github.com/shifty21/scone/command"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	log.Println("Please specify one of vanilla, cas, grpc")
	// 	os.Exit(0)
	// }
	// if args[0] == "--version" || args[0] == "version" {
	// 	log.Println("Vault Initializer Version - 0.01")
	// 	os.Exit(0)
	// }
	// config := config.LoadConfig()
	// switch args[0] {
	// case "http":
	// 	encryptionhttp.Run(
	// 		encryptionhttp.SetConfig(config),
	// 	)
	// case "grpc":
	// 	grpcServer, err := encryptiongrpc.NewGRPCService(
	// 		encryptiongrpc.SetConfig(config),
	// 		encryptiongrpc.EnableTLS(true),
	// 	)
	// 	if err != nil {
	// 		log.Printf("Error while creating grpc service")
	// 		os.Exit(1)
	// 	}
	// 	grpcServer.Run()
	// default:
	// 	v, err := vaultinterface.NewVaultInterface()
	// 	if err != nil {
	// 		fmt.Printf("Couldnt initialize vaultinterface %v", err)
	// 		os.Exit(0)
	// 	}
	// 	var options []vaultinterface.Option
	// 	forever := make(chan struct{})
	// 	switch args[0] {
	// 	case "scone":
	// 		log.Println("Initializing Vault with scone based encryption of vault initresponse")
	// 		v.EncryptKeyFun = vaultsconeinit.EncryptKeyFun
	// 		v.ProcessKeyFun = vaultsconeinit.ProcessKeyFun
	// 		options = append(options, vaultinterface.EnableGPGEncryption())
	// 		options = append(options, vaultinterface.SetGPGCryptoConfig(config.GetGPGCryptoConfig()))
	// 		options = append(options, vaultinterface.SconeCryptoConfig(config.GetCryptoConfig()))
	// 	case "auto":
	// 		log.Println("Initializing Vault with auto-seal")
	// 		v.EncryptKeyFun = vaultautoinit.EncryptKeyFun
	// 		v.ProcessKeyFun = vaultautoinit.ProcessKeyFun
	// 		options = append(options, vaultinterface.EnableGPGEncryption())
	// 		options = append(options, vaultinterface.SetGPGCryptoConfig(config.GetGPGCryptoConfig()))
	// 		options = append(options, vaultinterface.EnableAutoInitialization())
	// 	case "vanilla":
	// 		log.Println("Initializing Vault without encryption")
	// 		v.EncryptKeyFun = vaultinit.EncryptKeyFun
	// 		v.ProcessKeyFun = vaultinit.ProcessKeyFun
	// 	default:
	// 		log.Printf("Please specify gpg, scone, vanilla")
	// 		os.Exit(1)
	// 	}
	// 	options = append(options, vaultinterface.SetConfig(config.GetVaultConfig()))
	// 	err = v.Finalize(options...)
	// 	if err != nil {
	// 		log.Printf("Error finalizing vaultinterface %v", err)
	// 		os.Exit(1)
	// 	}
	// 	v.Run()
	// 	<-forever
	// }
}
