storage "consul" {
  address = "consul:8500"
  path    = "vault"
}
listener "tcp" {
  address     = "0.0.0.0:8200"
  tls_disable = 1
}
seal "scone" {
  vault_init_grpc_tls_cert = "/root/go/bin/resources/vault-init/grpc/cert/ca.cert"
  vault_init_grpc_tls_key  = "/root/go/bin/resources/vault-init/grpc/cert/ca.key"
  vault_init_grpc_address = "127.0.0.1:8082"
  vault_init_grpc_tls_enable = true
}

disable_mlock = true

api_addr = "http://127.0.0.1:8200"
cluster_addr = "https://127.0.0.1:8201"
ui = true
