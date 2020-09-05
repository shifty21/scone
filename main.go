package main

import (
	"github.com/shifty21/scone/vaultinitcas"
	"github.com/shifty21/scone/vaultinterface"
)

func main() {
	vaultinterface.GetConfig()
	vaultinitcas.CASCONFIG.Finalize()
	vaultinterface.Run(vaultinitcas.EncryptKeyFun, vaultinitcas.ProcessKeyFun)
}
