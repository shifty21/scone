reload_signal = "SIGHUP"
kill_signal = "SIGINT"
max_stale = "4m"
log_level = "info"
vault {
    address = "http://127.0.0.1:8200"
    token = "s.oABAsfib14wT318Sfxl8hLfk"
    renew_token = false
}
cas {
get_session_api = "https://localhost:8081/session"
cert = "resources/demo-client/conf/client.crt"
key = "resources/demo-client/conf/client-key.key"
session_name = "demo-client"
session_file = "resources/demo-client/session.yml"
test_updated_session_file = "resources/demo-client/test_updated_session.yml"
predecessor_hash_file = "resources/demo-client/predecessor_hash.yaml"
}

template {
    source      = "resources/consul-template/templates/config.yml.tpl"
    destination = "resources/consul-template/templates/config.yml"
  error_on_missing_key = true
  command = "go run demo-client/main.go "
  command_timeout = "0s"
}