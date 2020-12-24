package cas

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

//RegisterSession contains session related objects
type RegisterSession struct {
	Session        *SessionYAML
	ENV            string `yaml:"env"`
	Key            string `yaml:"key"`
	Certificate    string `yaml:"certificate"`
	Command        string `yaml:"command"`
	SessionFileLoc string `yaml:"session_file_loc"`
}

//Register contains session related objects
type Register struct {
	CASAddress string             `yaml:"cas_address"`
	Sessions   []*RegisterSession `yaml:"sessions"`
}

//LoadRegisterSessionConfig loads all the session that needs to be registered
func LoadRegisterSessionConfig(filepath string) (*Register, error) {
	// log.Printf("Session file %v\n", *config.SessionFile)
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Unable to open file %w", err)
	}
	register := &Register{}
	err = yaml.Unmarshal(fileContent, register)
	if err != nil {
		return nil, fmt.Errorf("[ERR] while unmarshalling session file %w", err)
	}
	return register, nil
}

//GetMREnclave runs a command and gets session
func GetMREnclave(command string, env string) error {
	cmd := exec.Command(command)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	additionalEnv := strings.Split(env, " ")
	newEnv := os.Environ()
	for i := range additionalEnv {
		newEnv = append(newEnv, additionalEnv[i])
	}
	cmd.Env = newEnv
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	// Enclave hash:
	fmt.Printf("\nout:========\n%s\nerr:========\n%s\n", outStr, errStr)
	return nil
}

// RegisterCASSession reads secret from sessionfile and post it to CAS
// and update session file with udpated hash
func RegisterCASSession(config *Register) error {
	if config == nil {
		return ErrConfigNotFound
	}
	for x := range config.Sessions {
		log.Printf("Sessions to be registered %v", config.Sessions[x])
		session, err := GetSessionFromYaml(config.Sessions[x].SessionFileLoc)
		if err != nil {
			return fmt.Errorf("[ERR] while getting session from cas %w", err)
		}
		config.Sessions[x].Session = session
		updateHash, err := POSTCASSession(config.CASAddress, config.Sessions[x], session)
		if err != nil {
			log.Printf("[ERR] Updating cas session %v", err)
			return err
		}
		//Update config file instead
		session.Predecessor = *updateHash
		//Update pred hash
		err = StoreUpdatedSession(config.Sessions[x].SessionFileLoc, session)
		if err != nil {
			return fmt.Errorf("[ERR] writing updated session %w", err)
		}
		log.Printf("CAS session updated successfully %v", config.Sessions[x].Session.Name)
	}
	return nil
}

//POSTCASSession posts the session provided to cas with specified cert and key
func POSTCASSession(casAddress string, config *RegisterSession, session *SessionYAML) (*string, error) {
	// log.Printf("Loading certificate from %v and key from %v", *config.Certificate, *config.Key)
	cer, err := tls.LoadX509KeyPair(config.Certificate, config.Key)
	if err != nil {
		return nil, fmt.Errorf("[ERR] while loading keypair: %w", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	//marshall session and send post request to CAS
	marshalled, err := yaml.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("[ERR] marshalling session: %w", err)
	}
	log.Printf("Marshalled session %v", string(marshalled))
	var url = casAddress
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, fmt.Errorf("[ERR] getting session information from CAS server %s, error: %w", url, err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERR] making post request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("[ERR] Forbidden for: %v", url)
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("[ERR] No session named: %v", session.Name)
	}
	//Parse failure and sucess response individually
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		var errRequest FailRequest
		err = json.NewDecoder(resp.Body).Decode(&errRequest)
		if err != nil {
			return nil, fmt.Errorf("[ERR] Unable to decode FailRequest: %w", err)
		}
		return nil, fmt.Errorf("[ERR] Response code [%v], Response body: %v", resp.StatusCode, errRequest.Message)
	}
	var response PostResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Unable to decode request body: %w", err)
	}
	//Only in development
	session.Services[0].MREnclaves[0] = response.Hash
	err = StoreUpdatedSession(config.SessionFileLoc, session)
	if err != nil {
		log.Printf("[ERR] writing updated session %v", err)
	}
	return &response.Hash, nil
}
