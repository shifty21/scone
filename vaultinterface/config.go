package vaultinterface

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/shifty21/scone/logger"
)

var (
	vaultAddr             string
	checkIntervalDuration time.Duration
	httpClient            http.Client

	signalCh chan os.Signal
	stop     func()
	//InitResp from vault
	initResp *InitResponse
)

//SetInitResponse sets initResponse
func SetInitResponse(initResponse *InitResponse) {
	initResp = initResponse
}

//GetInitResponse gets initResponse
func GetInitResponse() (initResponse *InitResponse) {
	return initResp
}

//Initialize reads config from env variables or set them to default values
func Initialize() {
	logger.Info.Println("GetConfig|Starting the vault-init service...")
	if s := os.Getenv("VAULT_ADDR"); s != "" {
		vaultAddr = s
	} else {
		vaultAddr = "https://127.0.0.1:8200"
	}

	checkInterval := os.Getenv("CHECK_INTERVAL")
	if s := os.Getenv("CHECK_INTERVAL"); s == "" {
		checkInterval = "10"
	}

	i, err := strconv.Atoi(checkInterval)
	if err != nil {
		log.Fatalf("CHECK_INTERVAL is invalid: %s", err)
	}

	checkIntervalDuration = time.Duration(i) * time.Second

	//get connection to kms cloud and storage for key
	httpClient = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	signalCh = make(chan os.Signal)
	signal.Notify(signalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)

	stop = func() {
		logger.Error.Println("Shutting down")
		os.Exit(0)
	}
}
