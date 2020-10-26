package encryptionhttp

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/rsacrypto"
	"github.com/urfave/negroni"
)

//HTTPServer server
type HTTPServer struct {
	Config      *config.Configuration
	SconeCrypto *rsacrypto.Crypto
	SignalCh    chan os.Signal
}

//HTTPServerOption interface for setting HTTPServer config
type HTTPServerOption func(*HTTPServer)

//SetConfig sets config
func SetConfig(config *config.Configuration) HTTPServerOption {
	return func(h *HTTPServer) {
		h.Config = config
	}
}

//SetSconeCrypto sets config
func SetSconeCrypto(Crypto *rsacrypto.Crypto) HTTPServerOption {
	return func(h *HTTPServer) {
		h.SconeCrypto = Crypto
	}
}

//Run start encryptionhttp
func Run(option ...HTTPServerOption) {
	log.Printf("%vStarting http server", encryptionserviceLog)
	s := &HTTPServer{
		SignalCh: make(chan os.Signal),
	}
	signal.Notify(s.SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	for _, o := range option {
		o(s)
	}
	sconecrypto, err := rsacrypto.InitCrypto(s.Config.GetCryptoConfig())
	if err != nil {
		log.Printf("%vError while initializing crypto module, Exiting %v", encryptionserviceLog, err)
		os.Exit(1)
	}
	s.SconeCrypto = sconecrypto
	router := mux.NewRouter()
	service := NewEncryptionhttp(s.SconeCrypto)
	router.Handle("/ping", http.HandlerFunc(PingHandler)).Methods("GET")
	router.Handle("/encrypt", EncryptHandler(nil, service)).Methods("POST")
	router.Handle("/decrypt", DecryptHandler(nil, service)).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	server := negroni.New(negroni.NewRecovery())
	server.UseHandlerFunc(router.ServeHTTP)
	portInfo := ":" + strconv.Itoa(s.Config.GetencryptionhttpConfig().Port())
	httpserver := &http.Server{
		Addr:           portInfo,
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		<-s.SignalCh
		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
		err := httpserver.Shutdown(ctxShutDown)
		if err != nil {
			log.Printf("%vServer Shutdown Failed:%v", encryptionserviceLog, err)
		}
		if err == http.ErrServerClosed {
			log.Printf("%vServer closed successfully", encryptionserviceLog)
		}

	}()
	log.Fatal(httpserver.ListenAndServe())
}
