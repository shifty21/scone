reload_signal = "SIGHUP"
kill_signal = "SIGINT"
max_stale = "10m"
log_level = "info"
wait {
    min = "5s"
    max = "10s"
}
vault {
    address = "http://127.0.0.1:8200"
    token = "s.9dPa8pbheYLEeGByB7Gqvg0r"
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
    source      = "resources/demo-client/config.yaml.tpl"
    destination = "resources/demo-client/config.yaml"
    wait {
    min = "2s"
    max = "10s"
  }
  error_on_missing_key = true
}