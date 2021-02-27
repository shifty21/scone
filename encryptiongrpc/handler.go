package encryptiongrpc

import (
	"context"
	"fmt"

	pb "github.com/shifty21/scone/encryptiongrpc/proto"
)

//SconeCryptoGRPC implements encryption grpc server
type SconeCryptoGRPC struct {
	pb.UnimplementedCryptoServiceServer
	service Service
}

//Encrypt handles encryption
func (s *SconeCryptoGRPC) Encrypt(ctx context.Context, request *pb.BytePacket) (*pb.BytePacket, error) {

	encrypted, err := s.service.Encrypt(ctx, request.Value)
	// log.Printf("%vPlain text: %v\n", handlerLog, string(request.Value))
	if err != nil {
		return nil, fmt.Errorf("%vError while encrypting data %v", handlerLog, err)
	}
	return &pb.BytePacket{Value: encrypted}, nil
}

//Decrypt handles encryption
func (s *SconeCryptoGRPC) Decrypt(ctx context.Context, request *pb.BytePacket) (*pb.BytePacket, error) {
	decrypted, err := s.service.Decrypt(ctx, request.Value)
	if err != nil {
		return nil, fmt.Errorf("%vError while decrypting data %v", handlerLog, err)
	}
	// log.Printf("%v|Plain text: %v\n", handlerLog, string(decrypted))
	return &pb.BytePacket{Value: decrypted}, nil
}
