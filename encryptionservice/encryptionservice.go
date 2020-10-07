package encryptionservice

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	"github.com/shifty21/scone/logger"
	"github.com/urfave/negroni"
)

//Run start encryptionservice
func Run(config *config.Configuration, crypto *crypto.Crypto) {
	router := mux.NewRouter()
	service := NewEncryptionService(crypto)
	router.Handle("/ping", http.HandlerFunc(PingHandler)).Methods("GET")
	router.Handle("/encrypt", EncryptHandler(nil, service)).Methods("POST")
	router.Handle("/decrypt", DecryptHandler(nil, service)).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	server := negroni.New(negroni.NewRecovery())
	server.UseHandlerFunc(router.ServeHTTP)
	portInfo := ":" + strconv.Itoa(config.GetEncryptionServiceConfig().Port())
	s := &http.Server{
		Addr:           portInfo,
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Error.Fatal(s.ListenAndServe())
}
