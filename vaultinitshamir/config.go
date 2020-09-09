package vaultinitshamir

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	// ErrConfigStringEmpty is the error returned when CAS is specified as an empty
	// string.
	ErrConfigStringEmpty = errors.New("VaultInitShamir: required fields cannot be empty, please check if crypto_public_key, and crypto_private_key are provided")
	//ErrConfigInvalidFormat invalid format
	ErrConfigInvalidFormat = errors.New("VaultInitShamir: invalid format")

	CryptoVar *Crypto
)

// Config for vaultinitshamir module
type Config struct {
	//PublicKeyPath for encryption
	PublicKeyPath *string
	//PrivateKeyPath for decryption, will be populated by CAS session
	PrivateKeyPath *string
}

// DefaultConfig is the default configuration.
func DefaultConfig() *Config {
	return &Config{}
}

// Finalize ensures there no nil pointers.
func (c *Config) Finalize() {
	if c.PublicKeyPath == nil {
		c.PublicKeyPath = StringFromEnv([]string{
			"CRYPTO_PUBLIC_KEY",
		}, "")
	}

	if c.PrivateKeyPath == nil {
		c.PrivateKeyPath = StringFromEnv([]string{
			"CRYPTO_PRIVATE_KEY",
		}, "")
	}

}

// GoString defines the printable version of this struct.
func (c *Config) GoString() string {
	if c == nil {
		return "(*Config)(nil)"
	}

	return fmt.Sprintf("&Config{"+
		"PublicKeyPath:%v, "+
		"PrivateKeyPath:%v, "+
		"}",
		c.PublicKeyPath,
		c.PrivateKeyPath,
	)
}

// ParseConfig parses a string of the format `minimum(:maximum)` into a
// Config.
func ParseConfig(s string) (*Config, error) {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return nil, ErrConfigStringEmpty
	}
	parts := strings.Split(s, ":")
	var PublicKeyPath, PriateKeyPath string
	switch len(parts) {
	case 4:
		PublicKeyPath, PriateKeyPath = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	default:
		return nil, ErrConfigInvalidFormat
	}
	var c Config
	c.PublicKeyPath = &PublicKeyPath
	c.PrivateKeyPath = &PriateKeyPath
	return &c, nil
}

// Set sets the value in the format min[:max] for a CAS timer.
func (c *Config) Set(value string) error {
	CAS, err := ParseConfig(value)
	if err != nil {
		return err
	}
	c.PublicKeyPath = CAS.PublicKeyPath
	c.PrivateKeyPath = CAS.PrivateKeyPath
	return nil
}

// String returns the string format for this CAS variable
func (c *Config) String() string {
	return fmt.Sprintf("%v:%v", c.PublicKeyPath, c.PrivateKeyPath)
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
