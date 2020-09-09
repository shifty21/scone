package vaultinitcas

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/logger"
)

// AuthVaultByCAS authenticates application with CAS.
func AuthVaultByCAS(casConfig *VaultCASConfig) bool {
	cer, err := tls.LoadX509KeyPair(*casConfig.Certificate, *casConfig.Key)
	if err != nil {
		logger.Error.Printf("Error while loading keypair %s", err)
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
		logger.Error.Printf("Error getting session information from CAS server, CAS get call [%s] %s \n", url, err)
		return false
	}
	if resp.StatusCode != 200 {
		logger.Error.Printf("Unable to verify session, CAS get call [%s] body [%s] error := %s \n", url, resp.Body, err)
		return false
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	logger.Info.Printf("Value of response %v ", bodyString)
	return true
}
