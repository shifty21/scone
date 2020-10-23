storage "raft" {
  path    = "./vault/data"
  node_id = "node1"
retry_join {
    leader_api_addr = "http://192.168.0.143:8200"
  }
}

listener "tcp" {
  address     = "192.168.0.8:8200"
  tls_disable = 1
}

disable_mlock = true

api_addr = "http://192.168.0.8:8200"
cluster_addr = "https://192.168.0.143:8200"
ui = true
