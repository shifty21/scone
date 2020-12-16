package config

// CAS is the Min/Max duration used by the Watcher
type CAS struct {
	//key for vault-init session
	key string
	//url for cas
	url string
	//certificate for vault-init session
	certificate string
	//sessionName of vault-init
	sessionName string
	//sessionFile file
	sessionFile string
	//exportToSession exports secret to particular session
	exportToSession string
	//predecessorHashFile for predecessor cas hash
	predecessorHashFile string
}

//GetKey return session api
func (v *CAS) GetKey() string {
	return v.key
}

//GetURL return session api
func (v *CAS) GetURL() string {
	return v.url
}

//GetCertificate return session api
func (v *CAS) GetCertificate() string {
	return v.certificate
}

//GetSessionName return session api
func (v *CAS) GetSessionName() string {
	return v.sessionName
}

//GetPredecessorHashFile return session api
func (v *CAS) GetPredecessorHashFile() string {
	return v.predecessorHashFile
}

//GetSessionFile return session api
func (v *CAS) GetSessionFile() string {
	return v.sessionFile
}

//GetExportToSessionName export secret to particular session
func (v *CAS) GetExportToSessionName() string {
	return v.exportToSession
}

//LoadCASConfig loads values from viper
func LoadCASConfig() *CAS {
	return &CAS{
		key:                 getStringOrPanic("cas.key"),
		url:                 getStringOrPanic("cas.url"),
		certificate:         getStringOrPanic("cas.certificate"),
		sessionName:         getStringOrPanic("cas.session_name"),
		sessionFile:         getStringOrPanic("cas.session_file"),
		exportToSession:     getStringOrPanic("cas.export_to_session"),
		predecessorHashFile: getStringOrPanic("cas.predecessor_hash_file"),
	}

}
