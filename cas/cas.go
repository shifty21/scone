package cas

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
	"gopkg.in/yaml.v2"
)

var (
	//ErrConfigNotFound in case cas config is not loaded
	ErrConfigNotFound = errors.New("CAS configuration not found")
)

// getInjectionFileFromSessionYAML iterates over image and replace injection file content
// if it exists, otherwise simply adds it and return updated image
func updateSecretFromSession(secrets []Secret, secret Secret) []Secret {
	for k, v := range secrets {
		if v.Name == secret.Name {
			var temp Secret
			temp.Name = secret.Name
			temp.Export = secret.Export
			temp.ExportPublic = secret.ExportPublic
			temp.Kind = secret.Kind
			temp.Value = secret.Value
			secrets[k] = temp
			return secrets
		}
	}
	secrets = append(secrets, secret)
	return secrets
}

// UpdateCASSession reads secret from sessionfile add secrets and post it to CAS
// and update session file with udpated hash
func UpdateCASSession(config *config.CAS, secrets []Secret) error {
	if config == nil {
		return ErrConfigNotFound
	}
	session, err := GetCASSession(config)
	if err != nil {
		return fmt.Errorf("[ERR] while getting session from cas %w", err)
	}
	if session.Secrets != nil {
		if secrets != nil {
			for _, v := range secrets {
				log.Printf("Adding secret %v", v)
				session.Secrets = updateSecretFromSession(session.Secrets, v)
			}
		}
	} else {
		session.Secrets = secrets
	}

	updateHash, err := PUTCASSession(config, session)
	if err != nil {
		log.Printf("[ERR] Updating cas session %v", err)
		return err
	}
	//Update config file instead
	session.Predecessor = *updateHash

	//Update pred hash
	err = StorePredecessorHash(config.GetPredecessorHashFile(), *updateHash)
	if err != nil {
		return fmt.Errorf("[ERR] writing updated session hash %w", err)
	}

	//Remove secrets and store the file
	session.Secrets = nil
	err = StoreUpdatedSession(config.GetSessionFile(), session)
	if err != nil {
		log.Printf("[ERR] writing updated session %v", err)
	}
	log.Println("CAS session updated successfully")
	return nil
}

// Use atomic update in case there is parallel template rendering
// getPredecessorHash from give filepath
func getPredecessorHash(filePath string) (*string, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("[ERR] Unable to open file %v", err)
		return nil, err
	}
	var hash PredecessorHash
	err = yaml.Unmarshal(file, &hash)
	if err != nil {
		log.Printf("[ERR] while unmarshalling session file %v", err)
		return nil, err
	}
	return &hash.PredecessorHash, nil
}

// GetCASSession gets session file from cas and adds predecessor hash
// from predecessorhashfile
func GetCASSession(config *config.CAS) (*SessionYAML, error) {
	previousHash, err := getPredecessorHash(config.GetPredecessorHashFile())
	if err != nil {
		return nil, fmt.Errorf("[ERR] while getting predecessor hash at %v, error :%w", config.GetPredecessorHashFile(), err)
	}
	// log.Printf("Loading certificate from %v and key from %v", *config.Certificate, *config.Key)
	cer, err := tls.LoadX509KeyPair(config.GetCertificate(), config.GetKey())
	if err != nil {
		return nil, fmt.Errorf("[ERR] while loading keypair %w", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	var url = config.GetCASURL() + "/" + config.GetSessionName()
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("[ERR] Error getting session information from CAS server, CAS get call [%s] %s \n", url, err)
		return nil, err
	}
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("[ERR] Unable to verify session, CAS get call [%s], please check session key and certificate used are the same which were used during client registration", url)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("[ERR] No session named %v", config.GetSessionName())
	}
	if resp.StatusCode != 200 {
		var errRequest FailRequest
		err = json.NewDecoder(resp.Body).Decode(&errRequest)
		if err != nil {
			return nil, fmt.Errorf("[ERR] Unable to decode request %w", err)
		}
		return nil, fmt.Errorf("[ERR] Response code %v, response body [%s]", resp.StatusCode, errRequest.Message)
	}
	var response SessionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {

		return nil, fmt.Errorf("[ERR] Unable to decode request body %w", err)
	}
	var sessionYAML SessionYAML
	// sessionYAML := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(response.Session), &sessionYAML)
	if err != nil {
		return nil, fmt.Errorf("[ERR] while parsing session body %w", err)
	}
	sessionYAML.Predecessor = *previousHash
	return &sessionYAML, nil
}

//PUTCASSession posts the session provided to cas with specified cert and key
func PUTCASSession(config *config.CAS, session *SessionYAML) (*string, error) {
	// log.Printf("Loading certificate from %v and key from %v", *config.Certificate, *config.Key)
	cer, err := tls.LoadX509KeyPair(config.GetCertificate(), config.GetKey())
	if err != nil {
		return nil, fmt.Errorf("[ERR] while loading keypair: %w", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	//marshall session and send post request to CAS
	marshalled, err := yaml.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("[ERR] marshalling session: %w", err)
	}
	log.Printf("Marshalled session %v", string(marshalled))
	var url = config.GetCASURL()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, fmt.Errorf("[ERR] getting session information from CAS server %s, error: %w", url, err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERR] making post request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("[ERR] Forbidden for: %v", url)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("[ERR] No session named: %v", config.GetSessionName())
	}
	//Parse failure and sucess response individually
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		var errRequest FailRequest
		err = json.NewDecoder(resp.Body).Decode(&errRequest)
		if err != nil {
			return nil, fmt.Errorf("[ERR] Unable to decode FailRequest: %w", err)
		}
		return nil, fmt.Errorf("[ERR] Response code [%v], Response body: %v", resp.StatusCode, errRequest.Message)
	}
	var response PostResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Unable to decode request body: %w", err)
	}
	return &response.Hash, nil
}

//GetSessionFromYaml updates session file
func GetSessionFromYaml(filePath string) (*SessionYAML, error) {
	// log.Printf("Session file %v\n", *config.SessionFile)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Unable to open file %w", err)
	}
	var sessionFile SessionYAML
	err = yaml.Unmarshal(file, &sessionFile)
	if err != nil {
		return nil, fmt.Errorf("[ERR] while unmarshalling session file %w", err)
	}
	return &sessionFile, nil
}

// StoreUpdatedSession writes session file back with updated session
// This is only for demo purpose, remove in production
func StoreUpdatedSession(fileLocation string, session *SessionYAML) error {
	marshalledSession, err := yaml.Marshal(session)
	if err != nil {
		return fmt.Errorf("[ERR] while Marhalling session yaml %w", err)
	}
	err = ioutil.WriteFile(fileLocation, marshalledSession, 0644)
	if err != nil {
		return fmt.Errorf("[ERR] while writing session file %w", err)
	}
	return err
}

// StorePredecessorHash predecessor hash with updated value from cas
func StorePredecessorHash(filePath string, hash string) error {
	marshalledSession, err := yaml.Marshal(&PredecessorHash{PredecessorHash: hash})
	if err != nil {
		return fmt.Errorf("[ERR] while Marhalling PredecessorHash yaml %w", err)
	}
	err = ioutil.WriteFile(filePath, marshalledSession, 0644)
	if err != nil {
		return fmt.Errorf("[ERR] while writing Predecessor hash to file %w", err)
	}
	return nil
}

// getInjectionFileFromSessionYAML iterates over image and replace injection file content
// if it exists, otherwise simply adds it and return updated image
func getInjectionFileFromSessionYAML(image Image, injectFile *InjectionFile) Image {
	for k, v := range image.InjectionFiles {
		if v.Path == injectFile.Path {
			var temp InjectionFile
			// log.Printf("Replacing content of %v with value %v", v.Content, injectFile.Content)
			temp.Path = injectFile.Path
			temp.Content = injectFile.Content
			image.InjectionFiles[k] = temp
			return image
		}
	}
	image.InjectionFiles = append(image.InjectionFiles, *injectFile)
	return image
}
