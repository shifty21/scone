package vaultinitcas

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

// AuthCAS authenticates application with CAS.
func AuthCAS(casConfig *CASConfig) bool {
	cer, err := tls.LoadX509KeyPair(*casConfig.Certificate, *casConfig.Key)
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
	var url = *casConfig.GetSessionAPI + *casConfig.SessionName
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("[ERR] Error getting session information from CAS server, CAS get call [%s] %s \n", url, err)
		return false
	}
	if resp.StatusCode != 200 {
		log.Printf("[ERR] Unable to verify session, CAS get call [%s] body [%s] error := %s \n", url, resp.Body, err)
		return false
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Printf("[INFO] Value of response %v ", bodyString)
	return true
}
