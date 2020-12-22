package encryptiongrpc

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/shifty21/scone/config"
	pb "github.com/shifty21/scone/encryptiongrpc/proto"
	"github.com/shifty21/scone/rsacrypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//Option function interface for options
type Option func(o *GRPCServer)

//EnableTLS for starting tls based grpc server
func EnableTLS(b bool) Option {
	return func(o *GRPCServer) {
		o.EnableTLS = b
	}
}

//SetTLSConfig gets tls config
func SetTLSConfig(c *tls.Config) Option {
	return func(o *GRPCServer) {
		o.TLSConfig = c
	}
}

//SetConfig sets config
func SetConfig(c *config.Configuration) Option {
	return func(o *GRPCServer) {
		o.Config = c
	}
}

//CryptoService sets cryptoservice
func CryptoService(c *rsacrypto.Crypto) Option {
	return func(o *GRPCServer) {
		o.SconeCrypto = c
	}
}

//GRPCServer for encryption service
type GRPCServer struct {
	SignalChan  chan os.Signal
	Server      *grpc.Server
	EnableTLS   bool
	TLSConfig   *tls.Config
	Config      *config.Configuration
	SconeCrypto *rsacrypto.Crypto
}

//NewGRPCService creates new GRPCServer
func NewGRPCService(opts ...Option) (*GRPCServer, error) {
	SignalCh := make(chan os.Signal)
	signal.Notify(SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	s := &GRPCServer{
		SignalChan: SignalCh,
		EnableTLS:  false,
	}

	for _, o := range opts {
		o(s)
	}

	sconecrypto, err := rsacrypto.InitCrypto(s.Config.GetCryptoConfig())
	if err != nil {
		return nil, fmt.Errorf("%vError while initializing crypto module: %w", encryptionServiceLog, err)
	}
	s.SconeCrypto = sconecrypto
	if s.EnableTLS {
		log.Printf("%vLoading tls cred", encryptionServiceLog)
		tlsconfig, err := LoadTLSCredentials(s.Config)
		if err != nil {
			return nil, fmt.Errorf("%vError while generating certificate config %w", encryptionServiceLog, err)
		}
		s.Server = grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsconfig)))
	} else {
		log.Printf("%vEnabling grpc server", encryptionServiceLog)
		s.Server = grpc.NewServer()
	}
	return s, nil
}

//Run start encryptionhttp
func (s *GRPCServer) Run(opts ...Option) {
	log.Printf("%vStarting grpc server", encryptionServiceLog)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.GetGRPCServiceConfig().Port()))
	if err != nil {
		log.Printf("%vfailed to listen: %v", encryptionServiceLog, err)
		os.Exit(1)
	}

	service := NewEncryptionhttp(s.SconeCrypto)
	pb.RegisterCryptoServiceServer(s.Server, &SconeCryptoGRPC{service: service})
	go func() {
		select {
		case <-s.SignalChan:
			log.Printf("%vExiting server", encryptionServiceLog)
			s.Server.GracefulStop()
		}
	}()
	s.Server.Serve(lis)
}

//LoadTLSCredentials loads tls cert and key for grpc service
func LoadTLSCredentials(config *config.Configuration) (*tls.Config, error) {
	// Load server's certificate and private key
	log.Printf("LoadTLSCredentials| Certificate %v", config.GetGRPCServiceConfig().Certificate())
	log.Printf("LoadTLSCredentials| Key %v", config.GetGRPCServiceConfig().Key())
	serverCert, err := tls.LoadX509KeyPair(config.GetGRPCServiceConfig().Certificate(), config.GetGRPCServiceConfig().Key())
	if err != nil {
		return nil, err
	}
	// Create the credentials and return it
	tlsconfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return tlsconfig, nil
}
