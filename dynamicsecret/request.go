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

type Request struct {
	PluginName    string `json:"plugin_name"`
	AllowedRoles  string `json:"allowed_roles"`
	ConnectionURL string `json:"connection_url"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
}

type RoleRequest struct {
	DBName            string `json:"db_name"`
	CreationStatement string `json:"creation_statements"`
	DefaultTTL        string `json:"default_ttl"`
	MaxTTL            string `json:"max_ttl"`
}

type SecretEngine struct {
	Type string `json:"type"`
}

type Policy struct {
	Policy string `json:"policy"`
}

//CreateConfig craetes database config
func CreateConfig(conf *config.DynamicSecret) error {

	secretEngine := &SecretEngine{
		Type: "database",
	}
	log.Printf("Enabling secret engine with data %v", secretEngine)
	data, err := json.Marshal(secretEngine)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling secret engine request %v", err)
	}
	err = MakeRequest(conf.VaultSecretEngineURL, data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating secret enting %v", err)
	}

	roleRequest := &RoleRequest{
		DBName:            conf.Database,
		CreationStatement: conf.CreationStatement,
		DefaultTTL:        conf.DefaultTTL,
		MaxTTL:            conf.MaxTTL,
	}
	log.Printf("Making role request with data %v", roleRequest)
	data, err = json.Marshal(roleRequest)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	err = MakeRequest(conf.VaultRoleURL, data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}

	configRequest := &Request{
		PluginName:    conf.PluginName,
		AllowedRoles:  conf.AllowdRoles,
		ConnectionURL: conf.ConnectionURL,
		UserName:      conf.UserName,
		Password:      conf.Password,
	}
	log.Printf("Making config request with data %v", configRequest)
	data, err = json.Marshal(configRequest)
	if err != nil {
		return fmt.Errorf("[ERR] marshalling request for database config %v", err)
	}
	err = MakeRequest(conf.VaultConfigURL, data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}

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
	err = MakeRequest(conf.VaultRoleURL, data, conf.VaultToken)
	if err != nil {
		return fmt.Errorf("[ERR] while creating dynamic secret config %v", err)
	}
	return nil
}

type Response struct {
	Errors []string `json:"errors"`
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

	var responseData []byte
	_, err = resp.Body.Read(responseData)
	if err != nil {
		log.Printf("Error while read response %v", err)
	}
	response := &Response{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		log.Printf("Error while unmarshalling response %v", err)
	}
	log.Printf("MakeRequest|Sent database config response %v, with status code %v", resp, resp.StatusCode)

	return nil
}
