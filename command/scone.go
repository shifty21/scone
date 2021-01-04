package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/vaultinterface"
	"github.com/shifty21/scone/vaultsconeinit"
)

//Scone struct
type Scone struct {
	*RunOptions
}

//Run starts a scone based intialization
func (s *Scone) Run(args []string) int {
	log.Println("Initializing Vault with scone based encryption of vault initresponse")
	s.Vault.EncryptKeyFun = vaultsconeinit.EncryptKeyFun
	s.Vault.ProcessKeyFun = vaultsconeinit.ProcessKeyFun
	s.options = append(s.options, vaultinterface.EnableGPGEncryption())
	s.options = append(s.options, vaultinterface.SetGPGCryptoConfig(s.config.GetGPGCryptoConfig()))
	s.options = append(s.options, vaultinterface.SconeCryptoConfig(s.config.GetCryptoConfig()))
	s.options = append(s.options, vaultinterface.SetVaultConfig(s.config.GetVaultConfig()))
	s.options = append(s.options, vaultinterface.SetCASConfig(s.config.GetCASConfig()))
	err := s.Vault.Finalize(s.options...)
	if err != nil {
		log.Printf("Error finalizing vaultinterface %v", err)
		return 1
	}
	s.Vault.Run()
	return 0
}

//Help provides help for scone based initialization
func (s *Scone) Help() string {
	helpText := `

	Usage: vault_init scone
	
		Starts a scone based initialization
	
		  $ vault_init scone
		  
	`

	return strings.TrimSpace(helpText)
}

//Synopsis for scone based vault initialization
func (s *Scone) Synopsis() string {
	return "Initialize vault with gpg based encryption"
}
