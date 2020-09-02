package vaultinterface

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
