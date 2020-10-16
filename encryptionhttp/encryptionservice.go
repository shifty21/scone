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
	"github.com/shifty21/scone/crypto"
	"github.com/urfave/negroni"
)

//Run start encryptionhttp
func Run(config *config.Configuration, crypto *crypto.Crypto) {
	log.Println("Starting http server")
	SignalCh := make(chan os.Signal)
	signal.Notify(SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	router := mux.NewRouter()
	service := Newencryptionhttp(crypto)
	router.Handle("/ping", http.HandlerFunc(PingHandler)).Methods("GET")
	router.Handle("/encrypt", EncryptHandler(nil, service)).Methods("POST")
	router.Handle("/decrypt", DecryptHandler(nil, service)).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	server := negroni.New(negroni.NewRecovery())
	server.UseHandlerFunc(router.ServeHTTP)
	portInfo := ":" + strconv.Itoa(config.GetencryptionhttpConfig().Port())
	s := &http.Server{
		Addr:           portInfo,
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		<-SignalCh
		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
		err := s.Shutdown(ctxShutDown)
		if err != nil {
			log.Printf("Server Shutdown Failed:%+s", err)
		}
		if err == http.ErrServerClosed {
			log.Printf("Server closed successfully")
		}

	}()
	log.Fatal(s.ListenAndServe())
}
