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
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"gopkg.in/yaml.v2"
)

//RegisterSession contains session related objects
type RegisterSession struct {
	Session        *SessionYAML
	ENV            string `yaml:"env"`
	Key            string `yaml:"key"`
	Certificate    string `yaml:"certificate"`
	Command        string `yaml:"command"`
	Parameter      string `yaml:"parameter"`
	SessionFileLoc string `yaml:"session_file_loc"`
}

// the constant values below are valid for x86_64

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

// the constant values below are valid for x86_64
const (
	mfdCloexec  = 0x0001
	memfdCreate = 319
)

func runFromMemory(filePath string, args []string, env []string) string {
	log.Printf("Filepath %v with paramter %v env variables %v", filePath, args, env)
	fdName := "" // *string cannot be initialized
	fd, _, _ := syscall.Syscall(memfdCreate, uintptr(unsafe.Pointer(&fdName)), uintptr(mfdCloexec), 0)

	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading filepath %v error %v", filePath, err)
	}

	_, err = syscall.Write(int(fd), buffer)
	if err != nil {
		log.Printf("Error writing  %v", err)
	}
	fdPath := fmt.Sprintf("/proc/self/fd/%d", fd)
	envvar := os.Environ()
	for x := range env {
		log.Printf("Appending env %v", env[x])
		envvar = append(envvar, env[x])
	}
	envvar = append(envvar)
	argsCombined := []string{"scone"}
	for x := range args {
		argsCombined = append(argsCombined, args[x])
	}
	// syscall.ForkExec()
	log.Printf("Envvar %v", envvar)
	log.Printf("args %v", argsCombined)
	// var stdin, stdout, stder bytes.Buffer
	r, w, _ := os.Pipe()

	execSpec := &syscall.ProcAttr{
		Env:   envvar,
		Files: []uintptr{w.Fd(), w.Fd(), w.Fd(), fd},
	}
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)

		outC <- buf.String()
	}()
	fork, err := syscall.ForkExec(fdPath, argsCombined, execSpec)
	// go syscall.Exec(fdPath, argsCombined, envvar)
	if err != nil {
		log.Printf("Error Forking %v", err)
	}
	log.Printf("start new process success, pid %d.", fork)
	// back to normal state
	w.Close()
	// os.Stdout = old // restoring the real stdout
	out := <-outC

	// reading our temp stdout
	reg, _ := regexp.Compile("Enclave hash: [a-z0-9]*\\n")
	matched := reg.FindString(out)
	log.Printf("Regexoutput %v", matched)

	return matched
	// syscall.Read()
}

//GetMREnclave runs a command and gets session
func GetMREnclave(command string, env string, parameter string) (string, error) {
	hash := runFromMemory(command, strings.Split(parameter, " "), strings.Split(env, " "))
	if hash == "" {
		return "", fmt.Errorf("Error while getting enclave hash, please \n" +
			"check binary is present and its configuration is correct")
	}
	return strings.Split(hash, " ")[1], nil
}

// RegisterCASSession reads secret from sessionfile and post it to CAS
// and update session file with udpated hash
func RegisterCASSession(config *Register) error {
	if config == nil {
		return ErrConfigNotFound
	}
	for x := range config.Sessions {
		log.Printf("Sessions to be registered %v", config.Sessions[x].Session)
		session, err := GetSessionFromYaml(config.Sessions[x].SessionFileLoc)
		if err != nil {
			log.Printf("[ERR] while getting session from cas %v", err)
			continue
		}
		config.Sessions[x].Session = session
		hash, err := GetMREnclave(config.Sessions[x].Command, config.Sessions[x].ENV, config.Sessions[x].Parameter)
		if err != nil {
			log.Printf("[ERR] Getting MREnclave for %v", config.Sessions[x].Session.Name)
			continue
		}
		config.Sessions[x].Session.Services[0].MREnclaves = []string{hash}
		updateHash, err := POSTCASSession(config.CASAddress, config.Sessions[x], session)
		if err != nil {
			log.Printf("[ERR] Updating cas session %v", err)
			continue
		}
		//Update config file instead
		session.Predecessor = *updateHash
		//Update pred hash
		err = StoreUpdatedSession(config.Sessions[x].SessionFileLoc, session)
		if err != nil {
			log.Printf("[ERR] writing updated session %v", err)
			continue
		}
		log.Printf("CAS session updated successfully %v", config.Sessions[x].Session.Name)
	}
	log.Printf("Registered all sessions")
	return nil
}

//POSTCASSession posts the session provided to cas with specified cert and key
func POSTCASSession(casAddress string, config *RegisterSession, session *SessionYAML) (*string, error) {
	log.Printf("POSTCASSession| registring %v", config.Session.Name)
	cer, err := tls.LoadX509KeyPair(config.Certificate, config.Key)
	if err != nil {
		log.Printf("POSTCASSession|Error loading certificate and key")
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
		log.Printf("POSTCASSession|Error marshalling session")
		return nil, fmt.Errorf("[ERR] marshalling session: %w", err)
	}
	log.Printf("Marshalled session %v", string(marshalled))
	var url = casAddress
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))
	if err != nil {
		log.Printf("POSTCASSession|Error creating post req")
		return nil, fmt.Errorf("[ERR] getting session information from CAS server %s, error: %w", url, err)
	}
	log.Printf("POSTCASSession|sending post session %v", url)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("POSTCASSession|Error posting session")
		return nil, fmt.Errorf("[ERR] making post request: %w", err)
	}
	log.Printf("POSTCASSession|Sent posting session")
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
	session.Predecessor = response.Hash
	err = StoreUpdatedSession(config.SessionFileLoc, session)
	if err != nil {
		log.Printf("[ERR] writing updated session %v", err)
	}
	return &response.Hash, nil
}
