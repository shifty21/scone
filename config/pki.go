package config

import (
	"log"

	"github.com/spf13/viper"
)

//PKIEngine configration
type PKIEngine struct {
	SecretEngineURL        string `mapstructure:"secret_engine_url"`
	PolicyFile             string `mapstructure:"policy_file"`
	Name                   string `mapstructure:"name"`
	RoleURL                string `mapstructure:"role_url"`
	PolicyURL              string `mapstructure:"policy_url"`
	GenerateCertificateURL string `mapstructure:"generate_certificate_url"`
	//internal or exported type
	GenerateRootURL       string `mapstructure:"generate_root_url"`
	SetConfigURL          string `mapstructure:"set_config_url"`
	IssuingCertificates   string `mapstructure:"issuing_certificates"`
	CrlDistributionPoints string `mapstructure:"crl_distribution_points"`
	VaultToken            string
	VaultURL              string `mapstructure:"vault_url"`
	RootCACertLocation    string `mapstructure:"root_ca_cert_location"`
}

//GetRoleURL get api for creating pki role
func (p *PKIEngine) GetRoleURL() string {
	return p.VaultURL + p.RoleURL + p.Name
}

//GetPolicyURL get api for creating pki role
func (p *PKIEngine) GetPolicyURL() string {
	return p.VaultURL + p.PolicyURL + p.Name
}

//GetGenerateCertificateURL get api for creating certificate
func (p *PKIEngine) GetGenerateCertificateURL() string {
	return p.VaultURL + p.GenerateCertificateURL + p.Name
}

//GetGenerateRootCertificateURL get api for creating root CA certificate
func (p *PKIEngine) GetGenerateRootCertificateURL() string {
	return p.VaultURL + p.GenerateRootURL
}

//GetSetConfigURL get api for setting urls for crl distribution and issuing_ca
func (p *PKIEngine) GetSetConfigURL() string {
	return p.VaultURL + p.SetConfigURL
}

//GetSecretEngineURL get api for setting urls for enabling pki secret engine
func (p *PKIEngine) GetSecretEngineURL() string {
	return p.VaultURL + p.SecretEngineURL
}

//GetIssuingCertificates get api for setting crl_distribution_points url in pki config
func (p *PKIEngine) GetIssuingCertificates() string {
	return p.VaultURL + p.IssuingCertificates
}

//GetCRLDistributionPointURL get api for setting crl_distribution_points url in pki config
func (p *PKIEngine) GetCRLDistributionPointURL() string {
	return p.VaultURL + p.CrlDistributionPoints
}

//LoadPKIConfig for pki engine
func LoadPKIConfig() *PKIEngine {
	pkiEngine := &PKIEngine{}
	err := viper.UnmarshalKey("vault_pki_config", pkiEngine)
	if err != nil {
		log.Printf("Error loading share config, using default shares")
	}
	pkiEngine.VaultToken = getString("vault_token_env")
	log.Printf("Loaded configuration %v", pkiEngine)
	return pkiEngine
}
