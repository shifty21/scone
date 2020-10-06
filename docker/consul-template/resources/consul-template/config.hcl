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
    token = "s.kEqN1GSTPze6DKF43yomfr3b"
    renew_token = false
}
cas {
get_session_api = "https://127.0.0.1:8081/session/"
cert = "/resources/consul-template/conf/client.crt"
key = "/resources/consul-template/conf/client-key.key"
session_name = "client"
}

template {
    source      = "/resources/consul-template/find_address.tpl"
    destination = "/resources/consul-template/hashicorp_address.txt"
}
