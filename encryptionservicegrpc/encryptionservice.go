package encryptionservicegrpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	pb "github.com/shifty21/scone/sconecryptogrpc"
	"google.golang.org/grpc"
)

//Run start encryptionservice
func Run(config *config.Configuration, crypto *crypto.Crypto) {
	SignalCh := make(chan os.Signal)
	signal.Notify(SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetEncryptionServiceConfig().Port()))
	if err != nil {
		fmt.Printf("grpcTest|failed to listen: %v", err)
		os.Exit(1)
	}
	// certs, err := loadTLSCredentials()
	// if err != nil {

	// }
	grpcServer := grpc.NewServer()
	// if certs != nil {
	// 	grpcServer = grpc.NewServer(grpc.Creds(certs))
	// }
	service := NewEncryptionService(crypto)
	pb.RegisterCryptoServiceServer(grpcServer, &SconeCryptoGRPC{service: service})
	go func() {
		<-SignalCh
		grpcServer.GracefulStop()
	}()
	grpcServer.Serve(lis)

}
