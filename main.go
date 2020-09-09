package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/shifty21/scone/logger"
	"github.com/shifty21/scone/vaultinitshamir"
	"github.com/shifty21/scone/vaultinterface"
)

func main() {
	logger.InitLogger()

	publicKeyPath := "experiments/keypairs/mykey.pub"
	privateKeyPath := "experiments/keypairs/mykey.pem"
	config := &vaultinitshamir.Config{}

	config.PublicKeyPath = &publicKeyPath
	config.PrivateKeyPath = &privateKeyPath
	err := vaultinitshamir.InitShamirInterface(config)
	if err != nil {
		logger.Info.Printf("Error while initializing shamirInterface")
		os.Exit(-1)
	}
	// vaultinterface.Initialize()
	// vaultinterface.Run(vaultinitshamir.EncryptKeyFun, vaultinitshamir.ProcessKeyFun)

	initResponseByte, err := ioutil.ReadFile("initResponse.json")
	var initResponseJSON vaultinterface.InitResponse
	err = json.Unmarshal(initResponseByte, &initResponseJSON)
	if err != nil {
		logger.Info.Printf("main|Error while saving file to json %v\n", err)
	}
	logger.Info.Printf("Unencrypted Response %v", initResponseJSON.GoString())
	err = vaultinitshamir.EncryptKeyFun(&initResponseJSON)
	if err != nil {
		logger.Info.Printf("main|Error while encrypting %v\n", err)
	}
	// logger.Info.Printf("Encrypted Response %v", vaultinterface.GetInitResponse())
	decryptedRespose, err := vaultinitshamir.ProcessKeyFun()
	if err != nil {
		logger.Info.Printf("main|Error while decrypting %v\n", err)
	}
	logger.Info.Printf("Unencrypted Response %v", decryptedRespose.GoString())
}
