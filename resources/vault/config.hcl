backend "inmem" {}

storage "consul" {
  address = "127.0.0.1:8500"
  path    = "vault"
}
listener "tcp" {
  address     = "127.0.0.1:8200"
  tls_disable = 1
}
seal "scone" {
  VAULT_INIT_GRPC_TLS_CERT = "/Users/abhishekchoudhary/.go/src/github.com/hashicorp/scone/encryptionservicegrpc/cert/ca-cert.pem"
	VAULT_INIT_GRPC_TLS_KEY  = "/Users/abhishekchoudhary/.go/src/github.com/hashicorp/scone/encryptionservicegrpc/cert/ca-key.pem"
	VAULT_INIT_GRPC_ADDRESS = "http://127.0.0.1:8083/"
}

disable_mlock = true

api_addr = "http://127.0.0.1:8200"
cluster_addr = "https://127.0.0.1:8201"
ui = true
