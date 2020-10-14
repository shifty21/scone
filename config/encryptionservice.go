package config

//encryptionhttp configration
type encryptionhttp struct {
	//PublicKeyPath for encryption
	port int
	//PrivateKeyPath for decryption, will be populated by CAS session
	privateKeyPath string
}

//Port gives vault address
func (c *encryptionhttp) Port() int {
	return c.port
}

//LoadencryptionhttpConfig loads values from viper
func LoadencryptionhttpConfig() *encryptionhttp {
	return &encryptionhttp{
		port: getIntOrPanic("http_encryption_server.port"),
	}

}
