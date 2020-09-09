package vaultinitcas

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	// ErrVaultCASStringEmpty is the error returned when CAS is specified as an empty
	// string.
	ErrVaultCASStringEmpty = errors.New("VaultInitCAS: required fields cannot be empty, please check if get_session_api, cert, key and session_name are provided")
	//ErrVaultCASInvalidFormat invalid format
	ErrVaultCASInvalidFormat = errors.New("VaultInitCAS: invalid format")
	//VCASConfig variable
	VCASConfig VaultCASConfig
)

// VaultCASConfig is the Min/Max duration used by the Watcher
type VaultCASConfig struct {
	GetSessionAPI *string
	Certificate   *string
	Key           *string
	SessionName   *string
}

// DefaultVaultCASConfig is the default configuration.
func DefaultVaultCASConfig() *VaultCASConfig {
	return &VaultCASConfig{}
}

// Finalize ensures there no nil pointers.
func (c *VaultCASConfig) Finalize() {
	if c.GetSessionAPI == nil {
		c.GetSessionAPI = StringFromEnv([]string{
			"VAULT_CAS_GET_SESSION",
		}, "")
	}

	if c.Certificate == nil {
		c.Certificate = StringFromEnv([]string{
			"VAULT_CLIENT_CERT",
		}, "")
	}

	if c.Key == nil {
		c.Key = StringFromEnv([]string{
			"VAULT_CLIENT_KEY",
		}, "")
	}

	if c.SessionName == nil {
		c.SessionName = StringFromEnv([]string{
			"VAULT_SESSION_NAME",
		}, "")
	}

}

// GoString defines the printable version of this struct.
func (c *VaultCASConfig) GoString() string {
	if c == nil {
		return "(*VaultCASConfig)(nil)"
	}

	return fmt.Sprintf("&VaultCASConfig{"+
		"GetSessionAPI:%v, "+
		"Certificate:%v, "+
		"Key:%v,"+
		"SessionName:%v"+
		"}",
		c.GetSessionAPI,
		c.Certificate,
		c.Key,
		c.SessionName,
	)
}

// ParseVaultCASConfig parses a string of the format `minimum(:maximum)` into a
// VaultCASConfig.
func ParseVaultCASConfig(s string) (*VaultCASConfig, error) {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return nil, ErrVaultCASStringEmpty
	}
	parts := strings.Split(s, ":")
	var GetSessionAPI, Certificate, Key, SessionName string
	switch len(parts) {
	case 4:
		GetSessionAPI, Certificate, Key, SessionName = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2]), strings.TrimSpace(parts[3])
	default:
		return nil, ErrVaultCASInvalidFormat
	}
	var c VaultCASConfig
	c.GetSessionAPI = &GetSessionAPI
	c.Certificate = &Certificate
	c.Key = &Key
	c.SessionName = &SessionName
	return &c, nil
}

// Set sets the value in the format min[:max] for a CAS timer.
func (c *VaultCASConfig) Set(value string) error {
	CAS, err := ParseVaultCASConfig(value)
	if err != nil {
		return err
	}
	c.GetSessionAPI = CAS.GetSessionAPI
	c.Certificate = CAS.Certificate
	c.Key = CAS.Key
	c.SessionName = CAS.SessionName
	return nil
}

// String returns the string format for this CAS variable
func (c *VaultCASConfig) String() string {
	return fmt.Sprintf("%v:%v:%v:%v", c.GetSessionAPI, c.Certificate, c.Key, c.SessionName)
}

//StringFromEnv extracts string from env
func StringFromEnv(list []string, def string) *string {
	for _, s := range list {
		if v := os.Getenv(s); v != "" {
			trimmed := strings.TrimSpace(v)
			return &trimmed
		}
	}
	return &def
}
