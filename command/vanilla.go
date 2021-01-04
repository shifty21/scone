package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/vaultinterface"
	"github.com/shifty21/scone/vaultvanillainit"
)

//Vanilla struct
type Vanilla struct {
	*RunOptions
}

//Run initializes vault
func (v *Vanilla) Run(args []string) int {
	log.Println("Initializing Vault without encryption")
	v.Vault.EncryptKeyFun = vaultvanillainit.EncryptKeyFun
	v.Vault.ProcessKeyFun = vaultvanillainit.ProcessKeyFun

	v.options = append(v.options, vaultinterface.SetVaultConfig(v.config.GetVaultConfig()))
	v.options = append(v.options, vaultinterface.EnableVanillaInitialization())
	err := v.Vault.Finalize(v.options...)
	if err != nil {
		log.Printf("Error finalizing vaultinterface %v", err)
		return 1
	}
	v.Vault.Run()
	return 0
}

//Help provides help for vanilla command
func (v *Vanilla) Help() string {
	helpText := `

Usage: vault_init vanilla

Initializes vault without any encryption or decryption

	  $ vault_init vanilla
	  
`
	return strings.TrimSpace(helpText)
}

//Synopsis for gRPC
func (v *Vanilla) Synopsis() string {
	return "Initializes vault without any encryption or decryption"
}
