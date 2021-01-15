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

//GenerateCertificate generates a certificate and exports it to specific session
func GenerateCertificate(conf *config.PKIEngine) error {
	log.Printf("Generating certificate and key with token %v and url %v", conf.VaultToken, conf.GetGenerateCertificateURL())
	client := &http.Client{}
	configRequest := &GenerateRootRequest{
		CommonName: "127.0.0.1",
		Format:     "pem",
		TTL:        "10m",
	}
	requestData, err := json.Marshal(configRequest)
	if err != nil {
		return fmt.Errorf("Error while marshalling generate certificate request %v", err)
	}
	req, err := http.NewRequest("POST", conf.GetGenerateCertificateURL(), bytes.NewReader(requestData))
	if err != nil {
		return fmt.Errorf("[ERR] creating post request %s, error: %w", conf.GetGenerateCertificateURL(), err)
	}
	log.Printf("MakeRequest|sending post session %v with token %v", conf.GetGenerateCertificateURL(), conf.VaultToken)
	req.Header.Set("X-Vault-Token", conf.VaultToken)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[ERR] making post request: %w", err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while read response %v", err)
	}
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		return fmt.Errorf("Error while making request got status code %v, data %v", resp.StatusCode, string(responseData))
	}
	response := &CertificateResponse{}
	err = json.Unmarshal(responseData, response)
	if err != nil {
		log.Printf("Error while unmarshalling response %v, response body %v", err, string(responseData))
	}
	log.Printf("MakeRequest|response %v, with status code %v", response, resp.StatusCode)
	//export to cas session
	return nil
}
