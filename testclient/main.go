package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pb "github.com/shifty21/scone/encryptiongrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func httpTest() {
	message := &RequestByte{
		Data: "jackhammer",
	}

	marshalledmessage, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://localhost:8083/encrypt", "application/json", bytes.NewBuffer(marshalledmessage))
	if err != nil {
		log.Fatalln(err)
	}
	// var encryptedResult RequestByte

	// json.NewDecoder(resp.Body).Decode(&encryptedResult)
	// decryptmessage := &RequestByte{
	// 	Data: encryptedResult.Data,
	// }
	// log.Printf("encrypted message %v", encryptedResult)
	var responsebody RequestByte
	err = json.NewDecoder(resp.Body).Decode(&responsebody)
	if err != nil {
		log.Println("Error while decoding request body %v", err)
		return
	}
	marshalledmessage, err = json.Marshal(responsebody)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Encrypted Response %v", responsebody)
	respdecrypt, err := http.Post("http://localhost:8083/decrypt", "application/json", bytes.NewBuffer(marshalledmessage))
	if err != nil {
		log.Fatalln(err)
	}

	var resultdecrypt RequestByte

	json.NewDecoder(respdecrypt.Body).Decode(&resultdecrypt)

	log.Printf("response %v", resultdecrypt.Data)
}

func grpcTest(conn *grpc.ClientConn) {
	req := pb.BytePacket{Value: []byte("jackhammer")}
	client := pb.NewCryptoServiceClient(conn)
	response, err := client.Encrypt(context.Background(), &req)
	if err != nil {
		log.Printf("grpcTest|Error while making encrypt request %v", err)
	}
	log.Printf("grpcTest|encrypted response %v\n", string(response.Value))

	decryptedResponse, err := client.Decrypt(context.Background(), response)
	if err != nil {
		log.Printf("grpcTest|Error while making decrypt request %v\n", err)
	}
	log.Printf("grpcTest|Decrypted response %v\n", string(decryptedResponse.Value))
	conn.Close()
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("../encryptiongrpc/cert/ca.cert")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	tlsconfig := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(tlsconfig), nil
}

func grpcClient() {
	targetURL := "localhost:8082"
	// config := config.ConfigureAllInterfaces()
	tlsserver := true
	var grpcConn *grpc.ClientConn
	var err error
	if tlsserver {
		log.Println("Secure connection")
		creds, err := loadTLSCredentials()
		if err != nil {
			log.Println("Error while loading tls cred")
			os.Exit(1)
		}

		// creds, err := credentials.NewTLS("../encryptiongrpc/cert/ca-cert.pem", nil)
		// if err != nil {
		// 	log.Printf("Error while loading creds %v", err)
		// }
		grpcConn, err = grpc.Dial(targetURL, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Printf("Test.Main|Error while connecting to grpc ")
			os.Exit(1)
		}
	} else {
		log.Println("Unsecure connection")
		grpcConn, err = grpc.Dial(targetURL, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Printf("Test.Main|Error while connecting to grpc ")
			os.Exit(1)
		}
	}
	defer grpcConn.Close()
	grpcTest(grpcConn)

}
func main() {
	httpTest()
}

//RequestByte request json to get the flying status
type RequestByte struct {
	Data string `json:"data"`
}
