package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

//Configuration store all the configs
type Configuration map[string]interface{}

//Config map contains all the configuration
var Config Configuration

func fatalGetString(key string) string {
	checkKey(key)
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func getString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func getFeature(key string) bool {
	v, err := strconv.ParseBool(fatalGetString(key))
	if err != nil {
		return false
	}
	return v
}

func checkKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("%s key is not set", key)
	}
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func getIntOrPanic(key string) int {
	checkKey(key)
	v, err := strconv.Atoi(viper.GetString(key))
	panicIfErrorForKey(err, key)
	return v
}

func panicIfErrorForKey(err error, key string) {
	if err != nil {
		panic(fmt.Errorf("Could not parse key: %s. Error: %v", key, err))
	}
}

//ConfigureAllInterfaces configures vaultinitcas, vaultinitshamir configurations
func ConfigureAllInterfaces() *Configuration {
	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("resources/vault-init/")
	viper.AddConfigPath("/resources/vault-init/")
	viper.ReadInConfig()
	Config = Configuration{}
	Config["vault_config"] = LoadVaultConfig()
	Config["cas_config"] = LoadVaultCASConfig()
	Config["crypto_config"] = LoadCryptoConfig()
	Config["encryption_service"] = LoadCryptoConfig()
	return &Config
}

//GetVaultConfig return vault related config
func (c *Configuration) GetVaultConfig() *Vault {
	return Config["vault_config"].(*Vault)
}

//GetCASConfig returns cas related config
func (c *Configuration) GetCASConfig() *VaultCAS {
	return Config["cas_config"].(*VaultCAS)
}

//GetCryptoConfig returns cas related config
func (c *Configuration) GetCryptoConfig() *Crypto {
	return Config["crypto_config"].(*Crypto)
}

//GetEncryptionServiceConfig returns cas related config
func (c *Configuration) GetEncryptionServiceConfig() *EncryptionService {
	return Config["encryption_service"].(*EncryptionService)
}
