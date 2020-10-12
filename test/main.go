package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/shifty21/scone/sconecryptogrpc"
	"google.golang.org/grpc"
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
	// fmt.Printf("encrypted message %v", encryptedResult)
	var responsebody []byte
	resp.Body.Read(responsebody)
	marshalledmessage, err = json.Marshal(responsebody)
	if err != nil {
		log.Fatalln(err)
	}
	respdecrypt, err := http.Post("http://localhost:8083/decrypt", "application/json", bytes.NewBuffer(marshalledmessage))
	if err != nil {
		log.Fatalln(err)
	}

	var resultdecrypt RequestByte

	json.NewDecoder(respdecrypt.Body).Decode(&resultdecrypt)

	log.Println(resultdecrypt.Data)
}

func grpcTest(conn *grpc.ClientConn) {
	req := pb.StringPacket{Value: "jackhammer"}
	client := pb.NewCryptoServiceClient(conn)
	response, err := client.Encrypt(context.Background(), &req)
	if err != nil {
		fmt.Printf("grpcTest|Error while making encrypt request %v", err)
	}
	fmt.Printf("grpcTest|encrypted response %v", response.Value)

	decryptedResponse, err := client.Decrypt(context.Background(), response)
	if err != nil {
		fmt.Printf("grpcTest|Error while making decrypt request %v", err)
	}
	fmt.Printf("grpcTest|decrypted response %v", decryptedResponse.Value)
	conn.Close()
}

func main() {
	targetURL := "localhost:8083"
	conn, err := grpc.Dial(targetURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Test.Main|Error while connecting to grpc ")
	}
	defer conn.Close()
	grpcTest(conn)
}

//Request request json to get the flying status
type RequestByte struct {
	Data string `json:"data"`
}
