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
	// cas.GetCASSession(config.Config.GetCASConfig())
	// return 0
	session, err := cas.GetSessionFromYaml(c.config.GetCASConfig())
	if err != nil {
		log.Printf("Error while getting session %v", err)
		return 1
	}
	//add predecessor hash and secrets to session
	if c.config.GetCASConfig().GetPredecessorHash() == "empty" {
		log.Println("No Predecessof hash provided registring for first time")
		updateHash, err := cas.UpdateSession(c.config.GetCASConfig(), session)
		if err != nil {
			log.Printf("Error Updating cas session %v", err)
			return 1
		}
		session.Predecessor = *updateHash
		err = cas.UpdateSessionFile(c.config.GetCASConfig(), session)
		if err != nil {
			log.Printf("Error wwriting updated session %v", err)
		}
		log.Println("CAS session updated successfully")
		return 0
	}
	session.Predecessor = c.config.GetCASConfig().GetPredecessorHash()
	session = cas.AddSecrets(session, cas.Secret{
		Name: "RootToken5", Kind: "ascii",
		Value: "2nM9epXhYKhrlMf6b5gFy41xmQNEqx5", Export: []cas.ExportTo{{Session: c.config.GetCASConfig().GetExportToSessionName()}}})
	session = cas.AddSecrets(session, cas.Secret{Name: "RootToken6", Kind: "ascii", ExportPublic: true, Value: "JackHammer"})
	log.Printf("Updated Session %v", session)

	// cas.GetSession(c.config.GetCASConfig())
	updateHash, err := cas.UpdateSession(c.config.GetCASConfig(), session)
	if err != nil {
		log.Printf("Error Updating cas session %v", err)
		return 1
	}
	//Update config file instead
	// c.config.GetCASConfig().SetPredecessorHash(*updateHash)
	session.Predecessor = *updateHash
	err = cas.UpdateSessionFile(c.config.GetCASConfig(), session)
	if err != nil {
		log.Printf("Error wwriting updated session %v", err)
	}
	log.Println("CAS session updated successfully")
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
