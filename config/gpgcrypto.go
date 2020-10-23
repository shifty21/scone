package config

//GPGCrypto configration
type GPGCrypto struct {
	//PublicKeyPath for encryption
	publicKeyPath string
	//PrivateKeyPath for decryption, will be populated by CAS session
	privateKeyPath string
}

//PublicKeyPath gives vault address
func (c *GPGCrypto) PublicKeyPath() string {
	return c.publicKeyPath
}

//PrivateKeyPath gives vault address
func (c *GPGCrypto) PrivateKeyPath() string {
	return c.privateKeyPath
}

//LoadGPGCryptoConfig loads values from viper
func LoadGPGCryptoConfig() *GPGCrypto {
	return &GPGCrypto{
		publicKeyPath:  getStringOrPanic("gpgcrypto.public_key_path"),
		privateKeyPath: getStringOrPanic("gpgcrypto.private_key_path"),
	}

}
