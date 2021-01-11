package config

//DynamicSecret configration
type DynamicSecret struct {
	Database               string
	PluginName             string
	AllowdRoles            string
	ConnectionURL          string
	UserName               string
	Password               string
	VaultConfigURL         string
	VaultRoleURL           string
	VaultToken             string
	CreationStatement      string
	MaxTTL                 string
	DefaultTTL             string
	PolicyFile             string
	VaultCreatePolicyURL   string
	VaultSecretEngineURL   string
	VaultGenerateConfigURL string
}

//LoadDynamicSecretConfig for grpc server
func LoadDynamicSecretConfig() *DynamicSecret {
	return &DynamicSecret{
		Database:               getStringOrPanic("dynamic_secret.database"),
		PluginName:             getStringOrPanic("dynamic_secret.plugin_name"),
		AllowdRoles:            getStringOrPanic("dynamic_secret.allowed_roles"),
		ConnectionURL:          getStringOrPanic("dynamic_secret.connection_url"),
		UserName:               getStringOrPanic("dynamic_secret.username"),
		Password:               getString("mongodb_password"),
		VaultToken:             getString("VAULT_TOKEN"),
		VaultConfigURL:         getStringOrPanic("dynamic_secret.vault_config_url"),
		VaultRoleURL:           getStringOrPanic("dynamic_secret.vault_role_url"),
		CreationStatement:      getStringOrPanic("dynamic_secret.creation_statements"),
		MaxTTL:                 getStringOrPanic("dynamic_secret.max_ttl"),
		DefaultTTL:             getStringOrPanic("dynamic_secret.default_ttl"),
		PolicyFile:             getStringOrPanic("dynamic_secret.policy_file"),
		VaultCreatePolicyURL:   getStringOrPanic("dynamic_secret.vault_create_policy_url"),
		VaultSecretEngineURL:   getStringOrPanic("dynamic_secret.vault_secret_engine_url"),
		VaultGenerateConfigURL: getStringOrPanic("dynamic_secret.vault_generate_credential_url"),
	}
}
