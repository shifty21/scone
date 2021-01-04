package config

import (
	"log"

	"github.com/spf13/viper"
)

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
	exportToSession []*string
	//predecessorHashFile for predecessor cas hash
	predecessorHashFile string
	//testUpdatedSessionFile for string updated session for dev purpose only
	testUpdatedSessionFile string
}

//GetKey return session api
func (v *CAS) GetKey() string {
	return v.key
}

//GetURL return session api
func (v *CAS) GetCASURL() string {
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
func (v *CAS) GetExportToSessionName() []*string {
	return v.exportToSession
}

//UpdatedSessionFileLoc export secret to particular session
func (v *CAS) UpdatedSessionFileLoc() string {
	return v.testUpdatedSessionFile
}

//LoadCASConfig loads values from viper
func LoadCASConfig() *CAS {
	var exportToSession []*string
	err := viper.UnmarshalKey("cas.export_to_session", &exportToSession)
	if err != nil {
		log.Fatalf("Error getting gpgcrypto config %v", err)
	}
	log.Printf("gpgcrypt keyset %v", exportToSession)
	return &CAS{
		key:                    getStringOrPanic("cas.key"),
		url:                    getStringOrPanic("cas.url"),
		certificate:            getStringOrPanic("cas.certificate"),
		sessionName:            getStringOrPanic("cas.session_name"),
		sessionFile:            getStringOrPanic("cas.session_file"),
		exportToSession:        exportToSession,
		testUpdatedSessionFile: getStringOrPanic("cas.test_updated_session_file"),
		predecessorHashFile:    getStringOrPanic("cas.predecessor_hash_file"),
	}
}
