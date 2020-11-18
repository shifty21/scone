package cas

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
	"gopkg.in/yaml.v2"
)

// GetSession authenticates application with CAS.
func GetSession(config *config.CAS) bool {
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
	var getRequest GetRequest
	err = json.NewDecoder(resp.Body).Decode(&getRequest)
	if err != nil {
		log.Printf("[ERR] Unable to decode request body %v", err)
		return false
	}
	// log.Printf("[INFO] Value of response %v ", getRequest.Session)
	var sessionYAML SessionYAML
	// sessionYAML := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(getRequest.Session), &sessionYAML)
	if err != nil {
		log.Printf("Error while parsing session body %v", err)
	}
	log.Printf("Session %v", sessionYAML)
	return true
}

//PostSession authenticates application with CAS.
func PostSession(config *config.CAS) bool {
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
	session, err := GetUpdatedSession(config)
	if err != nil {
		log.Printf("Error while getting session %v", err)
		return false
	}
	marshalled, err := yaml.Marshal(session)
	if err != nil {
		log.Printf("Error marshalling session %v", err)
	}
	var url = config.GetURL()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		log.Printf("[ERR] Error getting session information from CAS server, CAS get call [%s] %s \n", url, err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making post request %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		log.Printf("[ERR] Forbidden for %v \n", url)
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
	var getRequest GetRequest
	err = json.NewDecoder(resp.Body).Decode(&getRequest)
	if err != nil {
		log.Printf("[ERR] Unable to decode request body %v", err)
		return false
	}
	log.Printf("[INFO] Value of response %v ", getRequest.Session)
	return true
}

//GetUpdatedSession updates session file
func GetUpdatedSession(config *config.CAS) (*SessionYAML, error) {
	log.Printf("Session file %v\n", config.GetSessionFile())
	file, err := ioutil.ReadFile(config.GetSessionFile())
	if err != nil {
		log.Printf("[ERR] Unable to open file %v", err)
		return nil, err
	}
	// log.Printf("Filecontent %v", file)
	var sessionFile SessionYAML
	err = yaml.Unmarshal(file, &sessionFile)
	if err != nil {
		log.Printf("Error whil unmarshalling session file %v", err)
		return nil, err
	}
	sessionFile.Predecessor = "c1dc09fc823912c6340cc157aabe58f0848989bfbbfc02812044fd28be1f2a27"
	secrets := Secrets{Name: "RootToken", Kind: "ascii", ExportPublic: true, Value: "2nM9epXhYKhrlMf6b5gFy41xmQNEqx5"}
	sessionFile.Secrets = append(sessionFile.Secrets, secrets)
	log.Printf("Updated Session %v", sessionFile)
	return &sessionFile, nil
}
