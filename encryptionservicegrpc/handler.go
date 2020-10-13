package encryptionservicegrpc

import (
	"context"
	"fmt"

	pb "github.com/shifty21/scone/sconecryptoproto"
)

type SconeCryptoGRPC struct {
	pb.UnimplementedCryptoServiceServer
	service Service
}

// //Encrypt handles encryption
func (s *SconeCryptoGRPC) Encrypt(ctx context.Context, request *pb.BytePacket) (*pb.BytePacket, error) {

	encrypted, err := s.service.Encrypt(ctx, request.Value)
	fmt.Printf("Encrypt|Plain text : %v\n", string(request.Value))
	if err != nil {
		fmt.Printf("Encrypt|Error while encrypting data %v", err)
		return nil, err
	}
	return &pb.BytePacket{Value: encrypted}, nil
}

//Decrypt handles encryption
func (s *SconeCryptoGRPC) Decrypt(ctx context.Context, request *pb.BytePacket) (*pb.BytePacket, error) {
	decrypted, err := s.service.Decrypt(ctx, request.Value)
	if err != nil {
		fmt.Printf("Decrypt|Error while encrypting data %v", err)
		return nil, err
	}
	fmt.Printf("Decrypt|Plain text : %v\n", string(decrypted))
	return &pb.BytePacket{Value: decrypted}, nil
}
