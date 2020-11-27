package config

import (
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

//Vault configration
type Vault struct {
	address       string
	checkInterval string
	Shares        *Shares
}

//Address gives vault address
func (v *Vault) Address() string {
	return v.address
}

//CheckInterval return check interval, time to check for vault initialization
func (v *Vault) CheckInterval() time.Duration {
	i, err := strconv.Atoi(v.checkInterval)
	if err != nil {
		log.Printf("Error while converting configured checkinterval to integer %v, setting default to 5 Second ", err)
		return time.Duration(5) * time.Second
	}
	return time.Duration(i) * time.Second
}

//Shares define all the shared needed for vault initialization
type Shares struct {
	//RecoveryShares default value 1
	RecoveryShares int `mapstructure:"recovery_shares"`
	//RecoveryThreshold default value 1
	RecoveryThreshold int `mapstructure:"recovery_threshold"`
	//SecretShares default value 1
	SecretShares int `mapstructure:"secret_shares"`
	//SecretThreshold default value 1
	SecretThreshold int `mapstructure:"secret_threshold"`
}

//LoadVaultConfig loads values from viper
func LoadVaultConfig() *Vault {
	shares := &Shares{}
	err := viper.UnmarshalKey("shares", shares)
	if err != nil {
		log.Printf("Error loading share config, using default shares")

	}
	if shares.RecoveryShares == 0 &&
		shares.RecoveryThreshold == 0 &&
		shares.SecretShares == 0 &&
		shares.SecretThreshold == 0 {
		shares.RecoveryShares = 1
		shares.RecoveryThreshold = 1
		shares.SecretShares = 1
		shares.SecretThreshold = 1
	}
	log.Printf("Shares %v", shares)
	return &Vault{
		address:       getStringOrPanic("vault_address"),
		checkInterval: getStringOrPanic("check_interval"),
		Shares:        shares,
	}
}
