package pki

type RoleRequest struct {
	Name           string `json:"name,omitempty"`
	TTL            string `json:"ttl,omitempty"`
	MaxTTL         string `json:"max_ttl,omitempty"`
	AllowLocalhost bool   `json:"allow_localhost,omitempty"`

	// used with allow_bare_domains and allow_subdomains
	AllowedDomains                []string `json:"allowed_domains,omitempty"`
	AllowedDomainsTemplate        bool     `json:"allowed_domains_template,omitempty"`
	AllowBareDomains              bool     `json:"allow_bare_domains,omitempty"`
	AllowSubdomains               bool     `json:"allow_subdomains,omitempty"`
	AllowGlobDomains              bool     `json:"allow_glob_domains,omitempty"`
	AllowAnyName                  bool     `json:"allow_any_name,omitempty"`
	EnforceHostnames              bool     `json:"enforce_hostnames,omitempty"`
	AllowIPSans                   bool     `json:"allow_ip_sans,omitempty"`
	AllowedURISans                string   `json:"allowed_uri_sans,omitempty"`
	AllowedOtherSans              string   `json:"allowed_other_sans,omitempty"`
	ServerFlag                    bool     `json:"server_flag,omitempty"`
	ClientFlag                    bool     `json:"client_flag,omitempty"`
	CodeSigningFlag               bool     `json:"code_signing_flag,omitempty"`
	EmailProtectionFlag           bool     `json:"email_protection_flag,omitempty"`
	KeyType                       string   `json:"key_type,omitempty"`
	KeyBits                       int      `json:"key_bits,omitempty"`
	KeyUsage                      []string `json:"key_usage,omitempty"`
	ExtKeyUsage                   []string `json:"ext_key_usage,omitempty"`
	ExtKeyUsageOids               string   `json:"ext_key_usage_oids,omitempty"`
	UseCSRCommonName              bool     `json:"use_csr_common_name,omitempty"`
	UseCSRSans                    bool     `json:"use_csr_sans,omitempty"`
	OU                            string   `json:"ou,omitempty"`
	Organization                  string   `json:"organization,omitempty"`
	Country                       string   `json:"country,omitempty"`
	Locality                      string   `json:"locality,omitempty"`
	Province                      string   `json:"province,omitempty"`
	StreetAddress                 string   `json:"street_address,omitempty"`
	PostalCode                    string   `json:"postal_code,omitempty"`
	SerialNumber                  string   `json:"serial_number,omitempty"`
	GenerateLease                 bool     `json:"generate_lease,omitempty"`
	NoStore                       bool     `json:"no_store,omitempty"`
	RequireCN                     bool     `json:"require_cn,omitempty"`
	PolicyIdentifiers             []string `json:"policy_identifiers,omitempty"`
	BasicConstraintsValidForNonCA bool     `json:"basic_constraints_valid_for_non_ca,omitempty"`
	NotBeforeDuration             string   `json:"not_before_duration,omitempty"`
}

type GenerateCertificateRequest struct {
	Name              string `json:"json,omitempty"`
	CommonName        string `json:"common_name,omitempty"`
	AltNames          string `json:"alt_names,omitempty"`
	IPSans            string `json:"ip_sans,omitempty"`
	URISans           string `json:"uri_sans,omitempty"`
	OtherSans         string `json:"other_sans,omitempty"`
	TTL               string `json:"ttl,omitempty"`
	Format            string `json:"format,omitempty"`
	PrivateKeyFormat  string `json:"private_key_format,omitempty"`
	ExcludeCNFromSans bool   `json:"exclude_cn_from_sans,omitempty"`
}

type GenerateRootRequest struct {
	Type       string `json:"type,omitempty"`
	CommonName string `json:"common_name,omitempty"`
	AltNames   string `json:"alt_names,omitempty"`
	IPSans     string `json:"ip_sans,omitempty"`
	URISans    string `json:"uri_sans,omitempty"`
	OtherSans  string `json:"other_sans,omitempty"`
	TTL        string `json:"ttl,omitempty"`
	//pem, der or pem_bundle
	Format string `json:"format,omitempty"`
	// default der other than pkcs8
	PrivateKeyFormat string `json:"private_key_format,omitempty"`
	//rsa or ec
	KeyType string `json:"key_type,omitempty"`
	//default to 2048 is using ec type then change to 224,256,384 or 521
	KeyBits int `json:"key_bits,omitempty"`
	//default to -1
	MaxPathLength       int    `json:"max_path_length,omitempty"`
	ExcludeCNFromSans   bool   `json:"exclude_cn_from_sans,omitempty"`
	PermittedDNSDomains string `json:"permitted_dns_domains,omitempty"`
	OU                  string `json:"ou,omitempty"`
	Organization        string `json:"organization,omitempty"`
	Country             string `json:"country,omitempty"`
	Locality            string `json:"locality,omitempty"`
	Province            string `json:"province,omitempty"`
	StreetAddress       string `json:"street_address,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	SerialNumber        string `json:"serial_number,omitempty"`
}

type SecretEngine struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	//default_lease_ttl, max_lease_ttl, force_no_cache
	Config map[string]string `json:"config,omitempty"`
}

type SetURLRequest struct {
	IssuingCertificates   []string `json:"issuing_certificates,omitempty"`
	CRLDistributionPoints []string `json:"crl_distribution_points,omitempty"`
	OCSPServers           []string `json:"ocsp_servers,omitempty"`
}

type Policy struct {
	Policy string `json:"policy,omitempty"`
}
