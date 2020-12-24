package cas

import (
	"testing"

	"github.com/shifty21/scone/config"
)

const ConfigPath = "../testresources/vault-init/"
const FaultySessionPath = "../testresources/vault-init/session_fault.yaml"

func TestGetCASSession(t *testing.T) {
	testConfig := config.LoadConfig(ConfigPath).GetCASConfig()
	secret := Secret{Kind: "ascii", ExportPublic: false, Name: "VAULT_RESPONSE_ENCRYPTED", Value: "Test_Secret"}
	secrets := []Secret{secret}
	cases := []struct {
		name    string
		conf    *config.CAS
		err     error
		secrets []Secret
	}{
		{
			name:    "Working",
			conf:    testConfig,
			err:     nil,
			secrets: secrets,
		},
		{
			name: "Nil CAS config",
			conf: nil,
			err:  ErrConfigNotFound,
		},
	}

	for _, tc := range cases {
		err := UpdateCASSession(tc.conf, tc.secrets)
		if err != nil {
			t.Fatalf("Error posting Session %v", err)
		}
	}
}
