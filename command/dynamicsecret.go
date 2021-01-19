package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/dynamicsecret"
)

//DynamicSecret struct
type DynamicSecret struct {
	*RunOptions
}

//Run runs a gRPC server
func (r *DynamicSecret) Run(args []string) int {
	log.Printf("Setting up dynamic secret")
	err := dynamicsecret.SetupDynamicSecret(r.config.GetDynamiSecret())
	if err != nil {
		log.Printf("Error setting up dynamic secret %v", err)
		return 1
	}
	log.Println("Completed")
	return 0
}

//Help provides help for gRPC command
func (r *DynamicSecret) Help() string {
	helpText := `

Usage: vault_init register resource/vault-init/

	Sets up database dynamic secret

	  $ vault_init register resource/vault-init/
	  
`
	return strings.TrimSpace(helpText)
}

//Synopsis for gRPC
func (r *DynamicSecret) Synopsis() string {
	return "Register cas sessions"
}
