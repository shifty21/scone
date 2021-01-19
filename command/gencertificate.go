package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/pki"
)

//GenerateCertificate struct
type GenerateCertificate struct {
	*RunOptions
}

//Run runs a gRPC server
func (r *GenerateCertificate) Run(args []string) int {
	err := pki.GenerateCertificate(r.config.GetPKIConfig())
	if err != nil {
		log.Printf("Error setting up generating certificate and key %v", err)
		return 1
	}
	log.Println("Completed")
	return 0
}

//Help provides help for gRPC command
func (r *GenerateCertificate) Help() string {
	helpText := `

Usage: vault_init register resource/vault-init/

	Generates pki certificates and exports them to specified sessions

	  $ vault_init generate_certificates resource/vault-init/
	  
`

	return strings.TrimSpace(helpText)
}

//Synopsis for gRPC
func (r *GenerateCertificate) Synopsis() string {
	return "Register cas sessions"
}
