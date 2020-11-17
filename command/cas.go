package command

import (
	"strings"

	"github.com/shifty21/scone/cas"
)

//CAS struct
type CAS struct {
	*RunOptions
}

//Run starts initialization process based on auto-unseal
func (a *CAS) Run(args []string) int {
	cas.GetSession(a.config.GetCASConfig())

	return 0
}

//Help provides help for auto initialization
func (a *CAS) Help() string {
	helpText := `

	Usage: vault_init cas
	
		Gets a cas session according to supplied configurations

		  $ vault_init cas
		  
	`

	return strings.TrimSpace(helpText)
}

//Synopsis for auto
func (a *CAS) Synopsis() string {
	return "Perform CAS related operations"
}
