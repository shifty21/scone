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
    token = "s.HYTNTvxIQVHXb1Ma7asHqTMv"
    renew_token = false
}

cas {
get_session_api = "https://cas:8081/session/"
cert = "/root/go/bin/resources/demo-client/conf/client.crt"
key = "/root/go/bin/resources/demo-client/conf/client-key.key"
session_name = "demo-client"
}

template {
    source      = "/root/go/bin/resources/consul-template/find_address.tpl"
    destination = "/root/go/bin/resources/consul-template/hashicorp_address.txt"
}