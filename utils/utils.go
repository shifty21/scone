package utils

import "fmt"

//InitStatus gets the initialized status for vault
type InitStatus struct {
	Initialized bool `json:"initialized"`
}

// InitResponse holds a Vault init response.
type InitResponse struct {
	Keys       []string `json:"keys"`
	KeysBase64 []string `json:"keys_base64"`
	RootToken  string   `json:"root_token"`
}

// GoString defines the printable version of this struct.
func (i *InitResponse) GoString() string {
	if i == nil {
		return "(*InitResponse)(nil)"
	}

	return fmt.Sprintf("&CASConfig{"+
		"Keys:%v, "+
		"KeysBase64:%v, "+
		"RootToken:%v,"+
		"}",
		i.Keys,
		i.KeysBase64,
		i.RootToken,
	)
}

// UnsealResponse holds a Vault unseal response.
type UnsealResponse struct {
	Sealed   bool `json:"sealed"`
	T        int  `json:"t"`
	N        int  `json:"n"`
	Progress int  `json:"progress"`
}

// InitRequest holds a Vault init request.
type InitRequest struct {
	SecretShares    int `json:"secret_shares"`
	SecretThreshold int `json:"secret_threshold"`
}

// UnsealRequest holds a Vault unseal request.
type UnsealRequest struct {
	Key   string `json:"key"`
	Reset bool   `json:"reset"`
}
