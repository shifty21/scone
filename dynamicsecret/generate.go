package dynamicsecret

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shifty21/scone/config"
)

//GenreateResponse for credentials
type GenreateResponse struct {
	Data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"data"`
}

//GenerateCreadentials generates dynamic secret for the config provided
func GenerateCreadentials(conf *config.DynamicSecret) error {
	log.Printf("Generating dynamic secret with token %v and url %v", conf.VaultToken, conf.VaultGenerateConfigURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.VaultGenerateConfigURL, nil)
	if err != nil {
		return fmt.Errorf("[ERR] creating post request %s, error: %w", conf.VaultGenerateConfigURL, err)
	}
	log.Printf("MakeRequest|sending post session %v with token %v", conf.VaultGenerateConfigURL, conf.VaultToken)
	req.Header.Set("X-Vault-Token", conf.VaultToken)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[ERR] making post request: %w", err)
	}
	defer resp.Body.Close()

	var responseData []byte
	_, err = resp.Body.Read(responseData)
	if err != nil {
		log.Printf("Error while read response %v", err)
	}
	response := &GenreateResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		log.Printf("Error while unmarshalling response %v, response body %v", err, string(responseData))
	}
	log.Printf("MakeRequest|Sent database config response %v, with status code %v", response, resp.StatusCode)
	return nil
}
