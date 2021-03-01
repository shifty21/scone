package pki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
)

func saveFile(data string, filePath string) error {
	return ioutil.WriteFile(filePath, []byte(data), 0644)
}

//GenerateRoot generates root certificate
func GenerateRoot(conf *config.PKIEngine) error {
	configRequest := &GenerateRootRequest{
		CommonName: conf.RootCommonName,
		TTL:        conf.RootMaxTTL,
	}
	// log.Printf("Making GenerateRootRequest CA request with data %v", configRequest)
	data, err := json.Marshal(configRequest)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	httpResponse, err := MakeRequest(conf.GetGenerateRootCertificateURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}
	rootCertificateResponse := &CertificateResponse{}
	err = json.Unmarshal(httpResponse, rootCertificateResponse)
	if err != nil {
		return fmt.Errorf("Error while unmarshalling GenerateRoot CA request %v", err)
	}
	//save certificate
	err = saveFile(rootCertificateResponse.Data.Certificate, conf.RootCACertLocation+"certificate.cert")
	if err != nil {
		log.Printf("Error while saving certificate %v", err)
	}
	//save issuing ca
	err = saveFile(rootCertificateResponse.Data.IssuingCA, conf.RootCACertLocation+"issuing_ca.cert")
	if err != nil {
		log.Printf("Error while saving issuing_ca cert %v", err)
	}
	//save private key
	err = saveFile(rootCertificateResponse.Data.PrivateKey, conf.RootCACertLocation+"private_key.key")
	if err != nil {
		log.Printf("Error while saving privatekey %v", err)
	}
	//save serial
	err = saveFile(rootCertificateResponse.Data.SerialNumber, conf.RootCACertLocation+"serial.txt")
	if err != nil {
		log.Printf("Error while saving serial number %v", err)
	}
	return nil
}

//SetupRootCA configures and setup rootca
func SetupRootCA(conf *config.PKIEngine) error {
	//Enable pki
	secretEngine := &SecretEngine{
		Type: "pki",
	}
	log.Printf("Enabling secret engine with data %v", secretEngine)
	data, err := json.Marshal(secretEngine)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling secret engine request %v", err)
	}
	_, err = MakeRequest(conf.GetSecretEngineURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating secret %v", err)
	}
	//Generate root ca
	err = GenerateRoot(conf)
	if err != nil {
		return fmt.Errorf("Error while generating root CA %v", err)
	}
	//create role
	roleRequest := &RoleRequest{
		AllowLocalhost: true,
		AllowAnyName:   true,
		MaxTTL:         conf.MaxTTL,
		KeyUsage:       conf.KeyUsage,
		ExtKeyUsage:    conf.ExtKeyUsage,
	}
	// log.Printf("Making Role request with data %v", roleRequest)
	data, err = json.Marshal(roleRequest)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	_, err = MakeRequest(conf.GetRoleURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}
	//set config url
	configURls := &SetURLRequest{
		IssuingCertificates:   []string{conf.GetIssuingCertificates()},
		CRLDistributionPoints: []string{conf.GetCRLDistributionPointURL()},
	}
	log.Printf("Making config urls request with data %v", configURls)
	data, err = json.Marshal(configURls)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	_, err = MakeRequest(conf.GetSetConfigURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}
	//acl policycreate
	fileContent, err := ioutil.ReadFile(conf.PolicyFile)
	if err != nil {
		return fmt.Errorf("[ERR] Unable to open policy file at %v, err %w", conf.PolicyFile, err)
	}
	policy := &Policy{
		Policy: string(fileContent),
	}
	// log.Printf("Making Policy request with data %v", string(fileContent))
	data, err = json.Marshal(policy)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	_, err = MakeRequest(conf.GetPolicyURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}
	return nil
}

//MakeRequest makes requesf for give url and request payload
func MakeRequest(url string, data []byte, token string) ([]byte, error) {
	client := &http.Client{}
	r := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return nil, fmt.Errorf("[ERR] creating post request %s, error: %w", url, err)
	}
	log.Printf("MakeRequest|sending post session %v with token %v", url, token)
	req.Header.Set("X-Vault-Token", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("MakeRequest|Error posting session")
		return nil, fmt.Errorf("[ERR] making post request: %w", err)
	}
	log.Printf("MakeRequest|Status code %v", resp.StatusCode)
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while read response %v", err)
	}
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error while making request %v", string(responseData))
	}
	log.Printf("MakeRequest|Sent database config, with status code %v", resp.StatusCode)
	return responseData, nil
}
