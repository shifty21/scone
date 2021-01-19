package dynamicsecret

type Request struct {
	PluginName        string `json:"plugin_name"`
	AllowedRoles      string `json:"allowed_roles"`
	ConnectionURL     string `json:"connection_url"`
	UserName          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	TLSCertificateKey string `json:"tls_certificate_key,omitempty"`
	TLSCA             string `json:"tls_ca,omitempty"`
}

type RoleRequest struct {
	DBName            string `json:"db_name"`
	CreationStatement string `json:"creation_statements"`
	DefaultTTL        string `json:"default_ttl"`
	MaxTTL            string `json:"max_ttl"`
}

type SecretEngine struct {
	Type string `json:"type"`
}

type Policy struct {
	Policy string `json:"policy"`
}
type Response struct {
	Errors []string `json:"errors"`
}
