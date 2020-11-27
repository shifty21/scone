package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/shifty21/scone/encryptiongrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func grpcTest(conn *grpc.ClientConn) error {
	req := pb.BytePacket{Value: []byte("jackhammer")}
	client := pb.NewCryptoServiceClient(conn)
	response, err := client.Encrypt(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("grpcTest|Error while making encrypt request:%w", err)
	}
	decryptedResponse, err := client.Decrypt(context.Background(), response)
	if err != nil {
		return fmt.Errorf("grpcTest|Error while making decrypt request: %w", err)
	}
	log.Printf("grpcTest|Decrypted response %v\n", string(decryptedResponse.Value))
	conn.Close()
	return nil
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
	err = grpcTest(grpcConn)
	if err != nil {
		log.Printf("Error while Running grpcTest %v", err)
	}

}
func main() {
	// err := httpTest()
	// if err != nil {
	// 	log.Printf("Error running httpTest %v", err)
	// }
	grpcClient()
}

//RequestByte request json to get the flying status
type RequestByte struct {
	Data string `json:"data"`
}
