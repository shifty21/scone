package command

import (
	"log"
	"strings"

	"github.com/shifty21/scone/cas"
)

//CAS struct
type CAS struct {
	*RunOptions
}

//Run starts initialization process based on auto-unseal
func (c *CAS) Run(args []string) int {
	// cas.GetCASSession(c.config.GetCASConfig())
	// return 0
	err := cas.UpdateCASSession(c.config.GetCASConfig(), nil)
	if err != nil {
		log.Printf("Error while Posting session to cas %v", err)
		return 1
	}
	return 0
}

//Help provides help for auto initialization
func (c *CAS) Help() string {
	helpText := `

	Usage: vault_init cas
	
		Gets a cas session according to supplied configurations

		  $ vault_init cas
		  
	`

	return strings.TrimSpace(helpText)
}

//Synopsis for auto
func (c *CAS) Synopsis() string {
	return "Perform CAS related operations"
}
