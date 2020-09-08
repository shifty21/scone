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

	// vaultinterface.GetConfig()
	// vaultinitcas.CASCONFIG.Finalize()
	// vaultinterface.Run(vaultinitcas.EncryptKeyFun, vaultinitcas.ProcessKeyFun)

	v, err := vaultinitshamir.InitShamirInterface()
	if err != nil {
		logger.Info.Printf("Error while initializing shamirInterface")
		os.Exit(-1)
	}
	initResponseByte, err := ioutil.ReadFile("initResponse.json")
	var initResponseJSON vaultinterface.InitResponse
	err = json.Unmarshal(initResponseByte, &initResponseJSON)
	if err != nil {
		logger.Info.Printf("main|Error while saving file to json %v\n", err)
	}
	// logger.Info.Printf("RootToken %v \n", initResponseJSON.RootToken)
	// encryptedToken, err := v.EncryptText(initResponseJSON.Keys[1])
	// if err != nil {
	// 	logger.Info.Printf("main|Error while encrypting %v\n", err)
	// }
	// logger.Info.Printf("EncryptedToken %v\n", *encryptedToken)

	// decryptedToken, err := v.DecryptText(*encryptedToken)
	// if err != nil {
	// 	logger.Info.Printf("main|Error while decrypting %v\n", err)
	// }
	// logger.Info.Printf("RootToken %v\n", *decryptedToken)
	// logger.Info.Printf("InitResponse %v\n", initResponseJSON)
	// encryptedResponse, err := v.EncryptInitResponse(&initResponseJSON)
	// if err != nil {
	// 	logger.Info.Printf("main|Error while encrypting %v\n", err)
	// }
	// // logger.Info.Printf("PrivateKey %v", v.PrivateKey)
	// decryptedResponse, err := v.DecryptInitResponse(encryptedResponse)
	// if err != nil {
	// 	logger.Info.Printf("main|Error while decrypting %v\n", err)
	// }
	// logger.Info.Printf("InitResponse %v\n", decryptedResponse)
	logger.Info.Printf("Unencrypted Response %v", initResponseJSON.GoString())
	encryptedResponse, err := v.EncryptKeyFun(&initResponseJSON)
	if err != nil {
		logger.Info.Printf("main|Error while encrypting %v\n", err)
	}
	decryptedRespose, err := v.ProcessKeyFun(encryptedResponse)
	if err != nil {
		logger.Info.Printf("main|Error while decrypting %v\n", err)
	}
	logger.Info.Printf("Unencrypted Response %v", decryptedRespose.GoString())
}
