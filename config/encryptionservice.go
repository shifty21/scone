package config

//EncryptionService configration
type EncryptionService struct {
	//PublicKeyPath for encryption
	port int
	//PrivateKeyPath for decryption, will be populated by CAS session
	privateKeyPath string
}

//Port gives vault address
func (c *EncryptionService) Port() int {
	return c.port
}

//LoadEncryptionServiceConfig loads values from viper
func LoadEncryptionServiceConfig() *EncryptionService {
	return &EncryptionService{
		port: getIntOrPanic("http_encryption_server.port"),
	}

}
