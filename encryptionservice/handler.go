package encryptionservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shifty21/scone/logger"
	"google.golang.org/grpc"
)

//EncryptHandler handles encryption
func EncryptHandler(conn *grpc.ClientConn, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			logger.Error.Println("Error while decoding request body")
			RespondWithCustomErrors(w, nil, http.StatusBadRequest)
			return
		}
		data, err := service.EncryptData(r.Context(), request.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		encryptedData := &Response{Data: *data}
		// logger.Info.Printf("Encrypted Data %v", encryptedData)

		marshalled, err := json.Marshal(encryptedData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(marshalled)
		return
	}
}

//DecryptHandler handles encryption
func DecryptHandler(conn *grpc.ClientConn, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			logger.Error.Println("Error while decoding request body")
			RespondWithCustomErrors(w, nil, http.StatusBadRequest)
			return
		}
		decryptedData, err := service.DecryptData(r.Context(), request.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var response Response
		response.Data = *decryptedData
		marshalled, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(marshalled)
		return

	}
}

//NotFoundHandler not found
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("No such route")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

//PingHandler for checking service status
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := "Success"
	marshalledResult, _ := json.Marshal(result)
	w.Write(marshalledResult)
}

//RespondWithCustomErrors in case of user defined error
func RespondWithCustomErrors(w http.ResponseWriter, errorBody interface{}, statusCode int) error {
	w.WriteHeader(statusCode)
	body, err := json.Marshal(errorBody)

	if err != nil {
		return err
	}
	w.Write(body)
	return nil
}
