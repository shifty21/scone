package utils

import "fmt"

//InitStatus gets the initialized status for vault
type InitStatus struct {
	Initialized bool `json:"initialized"`
}

// InitResponse holds a Vault init response.
type InitResponse struct {
	Keys               []string `json:"keys"`
	KeysBase64         []string `json:"keys_base64"`
	RecoveryKeys       []string `json:"recovery_keys"`
	RecoveryKeysBase64 []string `json:"recovery_keys_base64"`
	RootToken          string   `json:"root_token"`
}

// UnsealResponse holds a Vault init response.
type UnsealResponse struct {
	Type        string `json:"type"`
	Initialized bool   `json:"initialized"`
	Sealed      bool   `json:"sealed"`
	T           int    `json:"t"`
	N           int    `json:"n"`
	Progress    int    `json:"progress"`

	Nonce        string `json:"nonce"`
	Version      string `json:"version"`
	Migration    bool   `json:"migration"`
	ClusterName  string `json:"cluster_name"`
	ClusterID    string `json:"cluster_id"`
	RecoverySeal bool   `json:"recovery_seal"`
	StorageType  string `json:"storage_type"`
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

// InitRequest holds a Vault init request.
type InitRequest struct {
	PGPKeys           []string `json:"pgp_keys"`
	SecretShares      int      `json:"secret_shares"`
	SecretThreshold   int      `json:"secret_threshold"`
	RecoveryShares    int      `json:"recovery_shares"`
	RecoveryThreshold int      `json:"recovery_threshold"`
	RecoveryPGPKeys   []string `json:"recovery_pgp_keys"`
	RootTokenPGPKey   string   `json:"root_token_pgp_key"`
}

// UnsealRequest holds a Vault unseal request.
type UnsealRequest struct {
	Key   string `json:"key"`
	Reset bool   `json:"reset"`
}

// InitializedUnsealRequest holds a Vault unseal request.
type InitializedUnsealRequest struct {
	Key string `json:"key"`
}
