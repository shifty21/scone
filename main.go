package main

import (
	"github.com/ykhedar/scone/vaultinitcas"
	"github.com/ykhedar/scone/vaultinterface"
)

func main() {
	vaultinterface.GetConfig()
	vaultinitcas.CASCONFIG.Finalize()
	vaultinterface.Run(vaultinitcas.EncryptKeyFun, vaultinitcas.ProcessKeyFun)
}
