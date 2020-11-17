package cas

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
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
	var url = config.GetURL() + config.GetSessionName()
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
	log.Printf("[INFO] Value of response %v ", getRequest.Session)
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
	var url = config.GetURL() + config.GetSessionName()
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("[ERR] Error getting session information from CAS server, CAS get call [%s] %s \n", url, err)
		return false
	}
	if resp.StatusCode == 403 {
		log.Printf("[ERR] Unable \n", url)
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
