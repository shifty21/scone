reload_signal = "SIGHUP"
kill_signal = "SIGINT"
max_stale = "4m"
log_level = "info"
vault {
    address = "http://localhost:8200"
    token = "s.pM50W8ai7dRZsSGIsdMx1dMp"
    renew_token = false
}
cas {
  get_session_api = "https://cas:8081/session"
  cert = "/root/go/bin/resources/vault-init/conf/client.crt"
  key = "/root/go/bin/resources/vault-init/conf/client-key.key"
  session_name = "vault-init-dynamicsecret"
  session_file = "/root/go/bin/resources/vault-init/session_dynamic.yml"
  test_updated_session_file = "/root/go/bin/resources/vault-init/dynamicsecret_updated_session.yml"
  predecessor_hash_file = "/root/go/bin/resources/vault-init/dynamicsecret_predecessor_hash.yaml"
}

template {
  source      = "resources/consul-template/templates/client-ca-cert.tpl"
  destination = "resources/consul-template/templates/client-ca.crt"
  perms       = "0600"
}

template {
  source      = "resources/consul-template/templates/client-pem-key.tpl"
  destination = "resources/consul-template/templates/client.pem"
}