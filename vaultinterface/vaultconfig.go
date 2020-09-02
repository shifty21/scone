package vaultinterface

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/cloudkms/v1"
)

var (
	vaultAddr             string
	checkIntervalDuration time.Duration
	gcsBucketName         string
	httpClient            http.Client

	kmsService *cloudkms.Service
	kmsKeyID   string

	storageClient *storage.Client
	signalCh      chan os.Signal
	stop          func()
	//InitResp from vault
	InitResp InitResponse

	userAgent = fmt.Sprintf("vault-init/0.1.0 (%s)", runtime.Version())
)

//GetConfig reads config from env variables or set them to default values
func GetConfig() {
	log.Println("Starting the vault-init service...")
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
		log.Printf("Shutting down")
		// kmsCtxCancel()
		// storageCtxCancel()
		os.Exit(0)
	}
}
