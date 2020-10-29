package command

import (
	"fmt"
	"strings"
)

//Version struct
type Version struct {
	*RunOptions
}

//Run starts a scone based intialization
func (v *Version) Run(args []string) int {
	fmt.Fprint(v.Stdout, "Vault Initializer Version - 0.01")
	return 0
}

//Help provides help for scone based initialization
func (v *Version) Help() string {
	helpText := `

	Usage: vault_init version
	
		Get version of vault_init
	
		  $ vault_init version
		  
	`

	return strings.TrimSpace(helpText)
}

//Synopsis for scone based vault initialization
func (v *Version) Synopsis() string {
	return "Get version"
}
