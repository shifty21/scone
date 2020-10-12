package config

//EncryptionService configration
type EncryptionGRPC struct {
	//certificate for tls grpc
	certificate string
	//key for tls grpc
	key  string
	port int
}

//Certificate for grpc tls
func (c *EncryptionGRPC) Certificate() string {
	return c.certificate
}

//Key for grpc tls
func (c *EncryptionGRPC) Key() string {
	return c.certificate
}

//Key for grpc tls
func (c *EncryptionGRPC) Port() int {
	return c.port
}

//LoadEncryptionGRPCConfig for grpc server
func LoadEncryptionGRPCConfig() *EncryptionGRPC {
	return &EncryptionGRPC{
		certificate: getStringOrPanic("grpc_encryption_server.cert"),
		key:         getStringOrPanic("grpc_encryption_server.key"),
		port:        getIntOrPanic("grpc_encryption_server.port"),
	}

}
