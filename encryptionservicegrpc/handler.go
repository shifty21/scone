package encryptionservicegrpc

import (
	"context"
	"fmt"

	pb "github.com/shifty21/scone/sconecryptogrpc"
)

type SconeCryptoGRPC struct {
	pb.UnimplementedCryptoServiceServer
	service Service
}

//EncryptHandler handles encryption
func (s *SconeCryptoGRPC) Encrypt(ctx context.Context, request *pb.StringPacket) (*pb.StringPacket, error) {
	encrypted, err := s.service.EncryptData(ctx, request.Value)
	fmt.Printf("Encrypt|Plain text : %v\n", request.Value)
	if err != nil {
		fmt.Printf("Encrypt|Error while encrypting data %v", err)
		return nil, err
	}
	return &pb.StringPacket{Value: *encrypted}, nil
}

//DecryptHandler handles encryption
func (s *SconeCryptoGRPC) Decrypt(ctx context.Context, request *pb.StringPacket) (*pb.StringPacket, error) {
	decrypted, err := s.service.DecryptData(ctx, request.Value)
	if err != nil {
		fmt.Printf("Decrypt|Error while encrypting data %v", err)
		return nil, err
	}
	fmt.Printf("Decrypt|Plain text : %v\n", decrypted)
	return &pb.StringPacket{Value: *decrypted}, nil
}

// //EncryptHandler handles encryption
func (s *SconeCryptoGRPC) EncryptByte(ctx context.Context, request *pb.BytePacket) (*pb.BytePacket, error) {

	encrypted, err := s.service.EncryptDataByte(ctx, request.Value)
	fmt.Printf("EncryptByte|Plain text : %v\n", string(request.Value))
	if err != nil {
		fmt.Printf("EncryptByte|Error while encrypting data %v", err)
		return nil, err
	}
	return &pb.BytePacket{Value: encrypted}, nil
}

//DecryptHandler handles encryption
func (s *SconeCryptoGRPC) DecryptByte(ctx context.Context, request *pb.BytePacket) (*pb.BytePacket, error) {
	decrypted, err := s.service.DecryptDataByte(ctx, request.Value)
	if err != nil {
		fmt.Printf("EncryptionServiceGRPC|Error while encrypting data %v", err)
		return nil, err
	}
	fmt.Printf("DecryptByte|Plain text : %v\n", string(decrypted))
	return &pb.BytePacket{Value: decrypted}, nil
}
