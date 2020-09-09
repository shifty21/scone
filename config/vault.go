package config

import (
	"strconv"
	"time"

	"github.com/shifty21/scone/logger"
)

//Vault configration
type Vault struct {
	address       string
	checkInterval string
}

//Address gives vault address
func (v *Vault) Address() string {
	return v.address
}

//CheckInterval return check interval, time to check for vault initialization
func (v *Vault) CheckInterval() time.Duration {
	i, err := strconv.Atoi(v.checkInterval)
	if err != nil {
		logger.Error.Printf("Error while converting configured checkinterval to integer %v", err)
	}
	return time.Duration(i) * time.Second
}

//LoadVaultConfig loads values from viper
func LoadVaultConfig() *Vault {
	return &Vault{
		address:       getStringOrPanic("vault_address"),
		checkInterval: getStringOrPanic("check_interval"),
	}

}
