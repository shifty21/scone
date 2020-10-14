# Vault Initializer
This code base intializes vault instance with shamir based initialization. The intialization is done in 2 ways as of now. 
1. Authorize vault shamir key-share with default 5 keys and 3 as threshold
2. Authorize vault by CAS along with keyshare encryption. The decryption keys are provided by CAS if the initializer is authenticated by CAS.
3. Auto-Initialization is done by gRPC encryption service, which provides encryption and decryption endpoints on a TLS or non-TLS based gRPC connection.
4. For Auto-Initialization there as change go-kms-wrapping, which included scone wrapper. And minor addition in vault for scone wrapper.
5. Changes were made in consul-template to authenticate demo-client app, before rendering template. This ensures that consul-template renders for authenticated clients only.