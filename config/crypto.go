package config

//Crypto configration
type Crypto struct {
	//PublicKeyPath for encryption
	publicKeyPath string
	//PrivateKeyPath for decryption, will be populated by CAS session
	privateKeyPath string
}

//PublicKeyPath gives vault address
func (c *Crypto) PublicKeyPath() string {
	return c.publicKeyPath
}

//PrivateKeyPath gives vault address
func (c *Crypto) PrivateKeyPath() string {
	return c.privateKeyPath
}

//LoadCryptoConfig loads values from viper
func LoadCryptoConfig() *Crypto {
	return &Crypto{
		publicKeyPath:  getStringOrPanic("public_key_path"),
		privateKeyPath: getStringOrPanic("private_key_path"),
	}

}
