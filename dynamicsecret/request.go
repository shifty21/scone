package dynamicsecret

type Request struct {
	PluginName    string `json:"plugin_name"`
	AllowedRoles  string `json:"allowed_roles"`
	ConnectionURL string `json:"connection_url"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
	// {
	// 	"plugin_name": "mysql-database-plugin",
	// 	"allowed_roles": "readonly",
	// 	"connection_url": "{{username}}:{{password}}@tcp(127.0.0.1:3306)/",
	// 	"username": "vaultuser",
	// 	"password": "secretpassword"
	//   }
}
