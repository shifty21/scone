package cas

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
	"gopkg.in/yaml.v2"
)

// GetCASSession authenticates application with CAS.
func GetCASSession(config *config.CAS) bool {
	log.Printf("Loading certificate from %v and key from %v", config.GetCertificate(), config.GetKey())
	cer, err := tls.LoadX509KeyPair(config.GetCertificate(), config.GetKey())
	if err != nil {
		log.Printf("Error while loading keypair %s", err)
		return false
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	var url = config.GetURL() + "/" + config.GetSessionName()
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("[ERR] Error getting session information from CAS server, CAS get call [%s] %s \n", url, err)
		return false
	}
	if resp.StatusCode == 403 {
		log.Printf("[ERR] Unable to verify session, CAS get call [%s], please check session key and certificate used are the same which were used during client registration \n", url)
		return false
	}

	if resp.StatusCode == 404 {
		log.Printf("[ERR] No session named %v \n", config.GetSessionName())
		return false
	}
	if resp.StatusCode != 200 {
		var errRequest FailRequest
		err = json.NewDecoder(resp.Body).Decode(&errRequest)
		if err != nil {
			log.Printf("[ERR] Unable to decode request %v", err)
			return false
		}
		log.Printf("[ERR] Response code %v, response body [%s]\n", resp.StatusCode, errRequest.Message)
		return false
	}
	var response SessionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Printf("[ERR] Unable to decode request body %v", err)
		return false
	}
	// log.Printf("[INFO] Value of response %v ", getRequest.Session)
	var sessionYAML SessionYAML
	// sessionYAML := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(response.Session), &sessionYAML)
	if err != nil {
		log.Printf("Error while parsing session body %v", err)
	}
	log.Printf("Session %v", sessionYAML)
	return true
}

//UpdateSession authenticates application with CAS.
func UpdateSession(config *config.CAS, session *SessionYAML) (*string, error) {
	log.Printf("Loading certificate from %v and key from %v", config.GetCertificate(), config.GetKey())
	cer, err := tls.LoadX509KeyPair(config.GetCertificate(), config.GetKey())
	if err != nil {
		return nil, fmt.Errorf("[Err] while loading keypair: %w", err)
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
		return nil, fmt.Errorf("[Err] marshalling session: %w", err)

	}
	var url = config.GetURL()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, fmt.Errorf("[Err] getting session information from CAS server %s, error: %w", url, err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[Err] making post request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("[ERR] Forbidden for: %w", url)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("[ERR] No session named: %w", config.GetSessionName())
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
	//Update session file with updated hash
	return &response.Hash, nil
}

//GetSessionFromYaml updates session file
func GetSessionFromYaml(config *config.CAS) (*SessionYAML, error) {
	log.Printf("Session file %v\n", config.GetSessionFile())
	file, err := ioutil.ReadFile(config.GetSessionFile())
	if err != nil {
		log.Printf("[ERR] Unable to open file %v", err)
		return nil, err
	}
	var sessionFile SessionYAML
	err = yaml.Unmarshal(file, &sessionFile)
	if err != nil {
		log.Printf("Error whil unmarshalling session file %v", err)
		return nil, err
	}
	return &sessionFile, nil
}

//UpdateSessionFile writes session file back with updated session
func UpdateSessionFile(config *config.CAS, session *SessionYAML) error {
	marshalledSession, err := yaml.Marshal(session)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(config.GetSessionFile(), marshalledSession, 0644)
	return err
}

//AddSecrets to session
func AddSecrets(session *SessionYAML, secret Secret) *SessionYAML {
	session.Secrets = append(session.Secrets, secret)
	return session
}
