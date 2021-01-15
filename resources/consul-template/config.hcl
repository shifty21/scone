reload_signal = "SIGHUP"
kill_signal = "SIGINT"
max_stale = "4m"
log_level = "info"
vault {
    address = "http://vault:8200"
    token = "s.E8OoVqCo6FsUSrWlgIHc8vRa"
    renew_token = false
}
cas {
get_session_api = "https://cas:8081/session"
cert = "/root/go/bin/resources/demo-client/conf/client.crt"
key = "/root/go/bin/resources/demo-client/conf/client-key.key"
session_name = "demo-client"
session_file = "/root/go/bin/resources/demo-client/session.yml"
test_updated_session_file = "/root/go/bin/resources/demo-client/test_updated_session.yml"
predecessor_hash_file = "/root/go/bin/resources/demo-client/predecessor_hash.yaml"
}

// template {
//     source      = "/root/go/bin/resources/consul-template/templates/config.yml.tpl"
//     destination = "/root/go/bin/resources/consul-template/templates/config.yml"
// }

template {
  source      = "/root/go/bin/resources/consul-template/templates/yet-cert.tpl"
  destination = "/root/go/bin/resources/consul-template/templates/yet.crt"
  perms       = "0600"
}

template {
  source      = "/root/go/bin/resources/consul-template/templates/yet-key.tpl"
  destination = "/root/go/bin/resources/consul-template/templates/yet.key"
}