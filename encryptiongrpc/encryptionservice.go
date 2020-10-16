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
	"github.com/shifty21/scone/crypto"
	pb "github.com/shifty21/scone/encryptiongrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//Options contains optional parameters for GRPCServer
type Options struct {
	EnableTLS     bool
	TLSConfig     *tls.Config
	Config        *config.Configuration
	CryptoService *crypto.Crypto
}

//Option function interface for options
type Option func(o *Options)

//EnableTLS for starting tls based grpc server
func EnableTLS(b bool) Option {
	return func(o *Options) {
		o.EnableTLS = b
	}
}

//TLSConfig gets tls config
func TLSConfig(c *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = c
	}
}

//Config sets config
func Config(c *config.Configuration) Option {
	return func(o *Options) {
		o.Config = c
	}
}

//CryptoService sets cryptoservice
func CryptoService(c *crypto.Crypto) Option {
	return func(o *Options) {
		o.CryptoService = c
	}
}

//GRPCServer for encryption service
type GRPCServer struct {
	Opts       Options
	SignalChan chan os.Signal
	Server     *grpc.Server
}

//NewGRPCService creates new GRPCServer
func NewGRPCService(opts ...Option) (*GRPCServer, error) {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	SignalCh := make(chan os.Signal)
	signal.Notify(SignalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)

	if options.EnableTLS {
		log.Println("Loading tls cred")
		tlsconfig, err := LoadTLSCredentials(options.Config)
		if err != nil {
			return nil, fmt.Errorf("Error while generating certificate config %w", err)
		}
		options.TLSConfig = tlsconfig
	}

	s := &GRPCServer{
		Opts:       options,
		SignalChan: SignalCh,
	}
	if options.EnableTLS {
		s.Server = grpc.NewServer(grpc.Creds(credentials.NewTLS(options.TLSConfig)))
	} else {
		log.Println("Enabling grpc server")
		s.Server = grpc.NewServer()
	}
	return s, nil
}

//Run start encryptionhttp
func (s *GRPCServer) Run(opts ...Options) {
	log.Println("Starting grpc server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Opts.Config.GetGRPCServiceConfig().Port()))
	if err != nil {
		log.Printf("grpcTest|failed to listen: %v", err)
		os.Exit(1)
	}

	service := Newencryptionhttp(s.Opts.CryptoService)
	pb.RegisterCryptoServiceServer(s.Server, &SconeCryptoGRPC{service: service})
	go func() {
		select {
		case <-s.SignalChan:
			log.Printf("Exiting server")
			s.Server.GracefulStop()
		}
	}()
	s.Server.Serve(lis)
}

//LoadTLSCredentials loads tls cert and key
func LoadTLSCredentials(config *config.Configuration) (*tls.Config, error) {
	// Load server's certificate and private key
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
