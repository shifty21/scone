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
	//hash for cas session
	predecessorHash string
	//name of session to which secret will be exposed
	exportToSessionName string
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

//GetExportToSessionName return session api
func (v *CAS) GetExportToSessionName() string {
	return v.exportToSessionName
}

//GetSessionFile return session api
func (v *CAS) GetSessionFile() string {
	return v.sessionFile
}

//GetPredecessorHash return session api
func (v *CAS) GetPredecessorHash() string {
	return v.predecessorHash
}

//SetPredecessorHash return session api
func (v *CAS) SetPredecessorHash(hash string) {
	v.predecessorHash = hash
}

//LoadCASConfig loads values from viper
func LoadCASConfig() *CAS {
	return &CAS{
		key:                 getStringOrPanic("cas.key"),
		url:                 getStringOrPanic("cas.url"),
		certificate:         getStringOrPanic("cas.certificate"),
		sessionName:         getStringOrPanic("cas.session_name"),
		sessionFile:         getStringOrPanic("cas.session_file"),
		predecessorHash:     getStringOrPanic("cas.predecessor_hash"),
		exportToSessionName: getStringOrPanic("cas.export_to_Session"),
	}

}