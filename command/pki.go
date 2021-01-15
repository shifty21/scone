package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/pki"
)

//PKISecret struct
type PKISecret struct {
	*RunOptions
}

//Run runs a gRPC server
func (p *PKISecret) Run(args []string) int {
	log.Printf("Setting up dynamic secret")
	err := pki.SetupRootCA(p.config.GetPKIConfig())
	if err != nil {
		log.Printf("Error setting up dynamic secret %v", err)
		return 1
	}
	log.Println("Completed")
	return 0
}

//Help provides help for gRPC command
func (p *PKISecret) Help() string {
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
func (p *PKISecret) Synopsis() string {
	return "Register cas sessions"
}
