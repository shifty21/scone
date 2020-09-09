package config

// VaultCAS is the Min/Max duration used by the Watcher
type VaultCAS struct {
	sessionAPI  string
	certificate string
	key         string
	sessionName string
}

//GetSessionAPI return session api
func (v *VaultCAS) GetSessionAPI() string {
	return v.sessionAPI
}

//GetCertificate return session api
func (v *VaultCAS) GetCertificate() string {
	return v.certificate
}

//GetKey return session api
func (v *VaultCAS) GetKey() string {
	return v.key
}

//GetSessionName return session api
func (v *VaultCAS) GetSessionName() string {
	return v.sessionName
}

//LoadVaultCASConfig loads values from viper
func LoadVaultCASConfig() *VaultCAS {
	return &VaultCAS{
		sessionAPI:  getStringOrPanic("session_api"),
		certificate: getStringOrPanic("certificate"),
		key:         getStringOrPanic("key"),
		sessionName: getStringOrPanic("session_name"),
	}

}
