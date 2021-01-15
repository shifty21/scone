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

	Register cas sesssion of the binary mentioned in the register_sessions.yaml file.
	For configuration only provide the directory and make sure the configuration filename 
	is register_session.yaml

	  $ vault_init register resource/vault-init/
	  
`

	return strings.TrimSpace(helpText)
}

//Synopsis for gRPC
func (r *GenerateCertificate) Synopsis() string {
	return "Register cas sessions"
}
