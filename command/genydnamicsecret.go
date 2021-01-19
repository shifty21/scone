package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/dynamicsecret"
)

//GenerateSecret struct
type GenerateSecret struct {
	*RunOptions
}

//Run runs a gRPC server
func (r *GenerateSecret) Run(args []string) int {
	err := dynamicsecret.GenerateCreadentials(r.config.GetDynamiSecret())
	if err != nil {
		log.Printf("Error setting up dynamic secret %v", err)
		return 1
	}
	log.Println("Completed")
	return 0
}

//Help provides help for gRPC command
func (r *GenerateSecret) Help() string {
	helpText := `

Usage: vault_init register resource/vault-init/

	Generates database dynamic secret and outputs to terminal

	  $ vault_init register resource/vault-init/
	  
`

	return strings.TrimSpace(helpText)
}

//Synopsis for gRPC
func (r *GenerateSecret) Synopsis() string {
	return "Register cas sessions"
}