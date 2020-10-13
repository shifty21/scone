package encryptionservicegrpc

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/shifty21/scone/config"
	"github.com/shifty21/scone/crypto"
	pb "github.com/shifty21/scone/sconecryptoproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Options struct {
	EnableTLS     bool
	TLSConfig     *tls.Config
	Config        *config.Configuration
	CryptoService *crypto.Crypto
}

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

type GRPCServer struct {
	Opts       Options
	SignalChan chan os.Signal
	Server     *grpc.Server
}

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
		fmt.Println("Loading tls cred")
		tlsconfig, err := LoadTLSCredentials(options.Config)
		if err != nil {
			fmt.Errorf("Error while generating certificate config")
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
		fmt.Println("Enabling grpc server")
		s.Server = grpc.NewServer()
	}
	return s, nil
}

//Run start encryptionservice
func (s *GRPCServer) Run(opts ...Options) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Opts.Config.GetGRPCServiceConfig().Port()))
	if err != nil {
		fmt.Printf("grpcTest|failed to listen: %v", err)
		os.Exit(1)
	}

	service := NewEncryptionService(s.Opts.CryptoService)
	pb.RegisterCryptoServiceServer(s.Server, &SconeCryptoGRPC{service: service})
	go func() {
		select {
		case <-s.SignalChan:
			fmt.Printf("Exiting server")
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
