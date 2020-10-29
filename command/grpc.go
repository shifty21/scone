package command

import (
	"fmt"
	"strings"

	"github.com/shifty21/scone/encryptiongrpc"
)

//GRPC struct
type GRPC struct {
	*RunOptions
}

//Run runs a gRPC server
func (g *GRPC) Run(args []string) int {
	grpcServer, err := encryptiongrpc.NewGRPCService(
		encryptiongrpc.SetConfig(g.config),
		encryptiongrpc.EnableTLS(true),
	)
	if err != nil {
		fmt.Printf("Error while creating grpc service: %v", err)
		return 1
	}
	grpcServer.Run()
	return 0
}

//Help provides help for gRPC command
func (g *GRPC) Help() string {
	helpText := `

Usage: vault_init grpc

	Starts a gRPC server that provides endpoints to encrypt and decrypt text

	  $ vault_init grpc
	  
`

	return strings.TrimSpace(helpText)
}

//Synopsis for gRPC
func (g *GRPC) Synopsis() string {
	return "Starts a gRPC server"
}
