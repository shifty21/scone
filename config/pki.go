package config

import (
	"log"

	"github.com/spf13/viper"
)

//PKIEngine configration
type PKIEngine struct {
	VaultToken             string
	RootCommonName         string `mapstructure:"root_common_name"`
	SecretEngineURL        string `mapstructure:"secret_engine_url"`
	PolicyFile             string `mapstructure:"policy_file"`
	Name                   string `mapstructure:"name"`
	MountPoint             string `mapstructure:"mount_point"`
	RoleURL                string `mapstructure:"role_url"`
	PolicyURL              string `mapstructure:"policy_url"`
	GenerateCertificateURL string `mapstructure:"generate_certificate_url"`
	//internal or exported type
	RootMaxTTL            string   `mapstructure:"root_max_ttl"`
	MaxTTL                string   `mapstructure:"max_ttl"`
	DefaultTTL            string   `mapstructure:"default_ttl"`
	ExtKeyUsage           []string `mapstructure:"ext_key_usage"`
	KeyUsage              []string `mapstructure:"key_usage"`
	GenerateRootURL       string   `mapstructure:"generate_root_url"`
	SetConfigURL          string   `mapstructure:"set_config_url"`
	IssuingCertificates   string   `mapstructure:"issuing_certificates"`
	CrlDistributionPoints string   `mapstructure:"crl_distribution_points"`
	VaultURL              string   `mapstructure:"vault_url"`
	RootCACertLocation    string   `mapstructure:"root_ca_cert_location"`
	GeneratedCertLocation string   `mapstructure:"generated_cert_location"`
	CertificateTTL        string   `mapstructure:"certificate_ttl"`
	CertificateCommonName string   `mapstructure:"certificate_common_name"`
	CertificateIPSAN      string   `mapstructure:"certificate_ip_san"`
}

//GetRoleURL get api for creating pki role
func (p *PKIEngine) GetRoleURL() string {
	return p.VaultURL + p.MountPoint + p.RoleURL + p.Name
}

//GetPolicyURL get api for creating pki role
func (p *PKIEngine) GetPolicyURL() string {
	return p.VaultURL + p.PolicyURL + p.Name
}

//GetGenerateCertificateURL get api for creating certificate
func (p *PKIEngine) GetGenerateCertificateURL() string {
	return p.VaultURL + p.MountPoint + p.GenerateCertificateURL + p.Name
}

//GetGenerateRootCertificateURL get api for creating root CA certificate
func (p *PKIEngine) GetGenerateRootCertificateURL() string {
	return p.VaultURL + p.MountPoint + p.GenerateRootURL
}

//GetSetConfigURL get api for setting urls for crl distribution and issuing_ca
func (p *PKIEngine) GetSetConfigURL() string {
	return p.VaultURL + p.MountPoint + p.SetConfigURL
}

//GetSecretEngineURL get api for setting urls for enabling pki secret engine
func (p *PKIEngine) GetSecretEngineURL() string {
	return p.VaultURL + p.SecretEngineURL + p.MountPoint
}

//GetIssuingCertificates get api for setting crl_distribution_points url in pki config
func (p *PKIEngine) GetIssuingCertificates() string {
	return p.VaultURL + p.MountPoint + p.IssuingCertificates
}

//GetCRLDistributionPointURL get api for setting crl_distribution_points url in pki config
func (p *PKIEngine) GetCRLDistributionPointURL() string {
	return p.VaultURL + p.MountPoint + p.CrlDistributionPoints
}

//LoadPKIConfig for pki engine
func LoadPKIConfig() *PKIEngine {
	pkiEngine := &PKIEngine{}
	err := viper.UnmarshalKey("vault_pki_config", pkiEngine)
	if err != nil {
		log.Printf("Error loading share config, using default shares")
	}
	pkiEngine.VaultToken = getString("vault_token_env")
	// log.Printf("Loaded configuration %v", pkiEngine)
	return pkiEngine
}
