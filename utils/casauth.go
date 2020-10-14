package utils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
)

// AuthVaultByCAS authenticates application with CAS.
func AuthVaultByCAS(config *config.VaultCAS) error {
	cer, err := tls.LoadX509KeyPair(config.GetCertificate(), config.GetKey())
	if err != nil {
		return fmt.Errorf("AuthVaultByCAS|Error while loading keypair %w", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	var url = config.GetSessionAPI() + config.GetSessionName()
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("AuthVaultByCAS|Error getting session information from CAS server %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("AuthVaultByCAS|Unable to verify session, CAS get call [%s] body [%s] error := %s", url, resp.Body, err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("AuthVaultByCAS|Error while reading response body %w", err)
	}
	bodyString := string(bodyBytes)
	log.Printf("AuthVaultByCAS|CAS Response %v ", bodyString)
	return nil
}
