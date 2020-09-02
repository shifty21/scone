package vaultinitcas

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	// ErrCASStringEmpty is the error returned when CAS is specified as an empty
	// string.
	ErrCASStringEmpty = errors.New("CAS: required fields cannot be empty, please check if get_session_api, cert, key and session_name are provided")
	//ErrCASInvalidFormat invalid format
	ErrCASInvalidFormat = errors.New("CAS: invalid format")
	//CASCONFIG variable
	CASCONFIG CASConfig
)

// CASConfig is the Min/Max duration used by the Watcher
type CASConfig struct {
	GetSessionAPI *string
	Certificate   *string
	Key           *string
	SessionName   *string
}

// DefaultCASConfig is the default configuration.
func DefaultCASConfig() *CASConfig {
	return &CASConfig{}
}

// Copy returns a deep copy of this configuration.
func (c *CASConfig) Copy() *CASConfig {
	if c == nil {
		return nil
	}

	var o CASConfig
	o.GetSessionAPI = c.GetSessionAPI
	o.Certificate = c.Certificate
	o.Key = c.Key
	o.SessionName = c.SessionName
	return &o
}

// Finalize ensures there no nil pointers.
func (c *CASConfig) Finalize() {
	if c.GetSessionAPI == nil {
		c.GetSessionAPI = StringFromEnv([]string{
			"CAS_GET_SESSION",
		}, "")
	}

	if c.Certificate == nil {
		c.Certificate = StringFromEnv([]string{
			"CLIENT_CERT",
		}, "")
	}

	if c.Key == nil {
		c.Key = StringFromEnv([]string{
			"CLIENT_KEY",
		}, "")
	}

	if c.SessionName == nil {
		c.SessionName = StringFromEnv([]string{
			"SESSION_NAME",
		}, "")
	}

}

// GoString defines the printable version of this struct.
func (c *CASConfig) GoString() string {
	if c == nil {
		return "(*CASConfig)(nil)"
	}

	return fmt.Sprintf("&CASConfig{"+
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

// ParseCASConfig parses a string of the format `minimum(:maximum)` into a
// CASConfig.
func ParseCASConfig(s string) (*CASConfig, error) {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return nil, ErrCASStringEmpty
	}

	parts := strings.Split(s, ":")
	var GetSessionAPI, Certificate, Key, SessionName string
	switch len(parts) {
	case 4:
		GetSessionAPI, Certificate, Key, SessionName = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2]), strings.TrimSpace(parts[3])
	default:
		return nil, ErrCASInvalidFormat
	}

	var c CASConfig
	c.GetSessionAPI = &GetSessionAPI
	c.Certificate = &Certificate
	c.Key = &Key
	c.SessionName = &SessionName

	return &c, nil
}

// Set sets the value in the format min[:max] for a CAS timer.
func (c *CASConfig) Set(value string) error {
	CAS, err := ParseCASConfig(value)
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
func (c *CASConfig) String() string {
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
