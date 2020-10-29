package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/vaultautoinit"
	"github.com/shifty21/scone/vaultinterface"
)

//Auto struct
type Auto struct {
	*RunOptions
}

//Run starts initialization process based on auto-unseal
func (a *Auto) Run(args []string) int {
	log.Println("Initializing Vault with auto-seal")
	a.Vault.EncryptKeyFun = vaultautoinit.EncryptKeyFun
	a.Vault.ProcessKeyFun = vaultautoinit.ProcessKeyFun
	a.options = append(a.options, vaultinterface.EnableGPGEncryption())
	a.options = append(a.options, vaultinterface.SetGPGCryptoConfig(a.config.GetGPGCryptoConfig()))
	a.options = append(a.options, vaultinterface.EnableAutoInitialization())
	a.options = append(a.options, vaultinterface.SetConfig(a.config.GetVaultConfig()))
	err := a.Vault.Finalize(a.options...)
	if err != nil {
		log.Printf("Error finalizing vaultinterface %v", err)
		return 1
	}
	a.Vault.Run()
	return 0
}

//Help provides help for auto initialization
func (a *Auto) Help() string {
	helpText := `

	Usage: vault_init auto
	
		Starts vault intialization based on recovery keys

		  $ vault_init auto
		  
	`

	return strings.TrimSpace(helpText)
}

//Synopsis for auto
func (a *Auto) Synopsis() string {
	return "Perform Vault initialization based on auto-unseal"
}
