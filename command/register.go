package command

import (
	"log"
	"os"
	"strings"

	"github.com/shifty21/scone/cas"
)

//Register struct
type Register struct {
	*RunOptions
}

//Run runs a gRPC server
func (r *Register) Run(args []string) int {
	file := os.Args[2]
	log.Printf("Loading configuraiton from %v", file)
	registerSessions, err := cas.LoadRegisterSessionConfig(file)
	if err != nil {
		log.Printf("Error while Posting session to cas %v", err)
		return 1
	}
	err = cas.RegisterCASSession(registerSessions)
	if err != nil {
		log.Printf("Error while registering session %v", err)
		return 1
	}
	return 0
}

//Help provides help for gRPC command
func (r *Register) Help() string {
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
func (r *Register) Synopsis() string {
	return "Register cas sessions"
}
