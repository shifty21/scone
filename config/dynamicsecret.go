package config

import (
	"log"

	"github.com/spf13/viper"
)

//DynamicSecret configration
type DynamicSecret struct {
	VaultURL           string `mapstructure:"vault_url"`
	TLSCertificateKey  string `mapstructure:"tls_certificate_key"`
	TLSCA              string `mapstructure:"tls_ca"`
	Database           string `mapstructure:"database"`
	PluginName         string `mapstructure:"plugin_name"`
	AllowdRoles        string `mapstructure:"allowed_roles"`
	UserName           string `mapstructure:"username"`
	Password           string `mapstructure:"mongodb_password"`
	VaultToken         string `mapstructure:"vault_token_env"`
	CreationStatement  string `mapstructure:"creation_statements"`
	MaxTTL             string `mapstructure:"max_ttl"`
	DefaultTTL         string `mapstructure:"default_ttl"`
	PolicyFile         string `mapstructure:"policy_file"`
	MongoConnectionURL string `mapstructure:"connection_url"`
	RoleURL            string `mapstructure:"role_url"`
	ConfigURL          string `mapstructure:"config_url"`
	PolicyURL          string `mapstructure:"policy_url"`
	SecretEngineURL    string `mapstructure:"secret_engine_url"`
	GenerateCredURL    string `mapstructure:"credential_url"`
}

//MongoDBConnectionURL gives api endpoint for setting dynamic creadential
func (d *DynamicSecret) GetMongoDBConnectionURL() string {
	return d.MongoConnectionURL
}

//ConfigURL gives api endpoint to create config
func (d *DynamicSecret) GetConfigURL() string {
	return d.VaultURL + d.ConfigURL + d.Database
}

//RoleURL gives api endpoint to setting up role
func (d *DynamicSecret) GetRoleURL() string {
	return d.VaultURL + d.RoleURL + d.AllowdRoles
}

//PolicyURL gives api endpoint to setting up policy
func (d *DynamicSecret) GetPolicyURL() string {
	return d.VaultURL + d.PolicyURL + d.AllowdRoles
}

//SecretEngineURL gives api endpoint to setting up database secret engine
func (d *DynamicSecret) GetSecretEngineURL() string {
	return d.VaultURL + d.SecretEngineURL
}

//DBCredentialURL gives api endpoint to setting up dynamic db credential
func (d *DynamicSecret) DBCredentialURL() string {
	return d.VaultURL + d.GenerateCredURL + d.AllowdRoles
}

//LoadDynamicSecretConfig for grpc server
func LoadDynamicSecretConfig() *DynamicSecret {
	dynamicSecret := &DynamicSecret{}
	err := viper.UnmarshalKey("vault_dynamic_secret", dynamicSecret)
	if err != nil {
		log.Printf("Error loading share config, using default shares")
	}
	dynamicSecret.Password = getString("mongodb_password")
	dynamicSecret.VaultToken = getString("vault_token_env")
	return dynamicSecret
}
