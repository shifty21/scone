package dynamicsecret

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
)

//SetupDynamicSecret creates database config
func SetupDynamicSecret(conf *config.DynamicSecret) error {
	// //enable database secret engine
	// secretEngine := &SecretEngine{
	// 	Type: "database",
	// }
	// log.Printf("Enabling secret engine with data %v", secretEngine)
	// data, err := json.Marshal(secretEngine)
	// if err != nil {
	// 	return fmt.Errorf("[ERR] marshalling secret engine request %v", err)
	// }

	// err = MakeRequest(conf.GetSecretEngineURL(), data, conf.VaultToken)
	// if err != nil {
	// 	return fmt.Errorf("[ERR] while creating secret enting %v", err)
	// }
	// // //Create database role
	// roleRequest := &RoleRequest{
	// 	DBName:            conf.Database,
	// 	CreationStatement: conf.CreationStatement,
	// 	DefaultTTL:        conf.DefaultTTL,
	// 	MaxTTL:            conf.MaxTTL,
	// }
	// log.Printf("Making role request with data %v", roleRequest)
	// data, err = json.Marshal(roleRequest)
	// if err != nil {
	// 	return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	// }
	// err = MakeRequest(conf.GetRoleURL(), data, conf.VaultToken)
	// if err != nil {
	// 	return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	// }
	// configure mongodb secret engine
	configRequest := &Request{
		PluginName:    conf.PluginName,
		AllowedRoles:  conf.AllowdRoles,
		ConnectionURL: conf.GetMongoDBConnectionURL(),
	}
	tlscertificateKey, err := ioutil.ReadFile(conf.TLSCertificateKey)
	if err != nil {
		return fmt.Errorf("Error while reading certificate file %v", err)
	}
	cacert, err := ioutil.ReadFile(conf.TLSCA)
	if err != nil {
		return fmt.Errorf("Error while reading ca certificate file %v", err)
	}
	configRequest.TLSCA = string(cacert)
	configRequest.TLSCertificateKey = string(tlscertificateKey)
	log.Printf("Making config request with data %v", configRequest)
	data, err := json.Marshal(configRequest)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	err = MakeRequest(conf.GetConfigURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}

	//Create policy
	fileContent, err := ioutil.ReadFile(conf.PolicyFile)
	if err != nil {
		return fmt.Errorf("[ERR] Unable to open policy file at %v, err %w", conf.PolicyFile, err)
	}
	policy := &Policy{
		Policy: string(fileContent),
	}
	log.Printf("Making Policy request with data %v", string(fileContent))
	data, err = json.Marshal(policy)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	err = MakeRequest(conf.GetPolicyURL(), data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}
	return nil
}

//MakeRequest makes requesf for give url and request payload
func MakeRequest(url string, data []byte, token string) error {
	client := &http.Client{}
	r := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return fmt.Errorf("[ERR] creating post request %s, error: %w", url, err)
	}
	log.Printf("MakeRequest|sending post session %v with token %v", url, token)
	req.Header.Set("X-Vault-Token", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("MakeRequest|Error posting session")
		return fmt.Errorf("[ERR] making post request: %w", err)
	}
	defer resp.Body.Close()

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while read response %v", err)
	}
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		return fmt.Errorf("Error while making request %v", string(responseData))
	}
	log.Printf("MakeRequest|Sent database config response %v, with status code %v", resp, resp.StatusCode)

	return nil
}
