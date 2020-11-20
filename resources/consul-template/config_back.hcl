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
    token = "s.nBNlTDf7RFmUw78W3EMosbDo"
    renew_token = false
}
cas {
get_session_api = "https://localhost:8081/session"
cert = "resources/consul-template/conf/client.crt"
key = "resources/consul-template/conf/client-key.key"
session_name = "consul-template"
session_file = "resources/consul-template/session.yml"
export_to_session = "demo-client"
predecessor_hash = "977ee24c3e6039dd862c75546f0d285da11b812a12b26f40fdc75096f2e8bd2f"
}

template {
    source      = "resources/consul-template/find_address.tpl"
    destination = "resources/consul-template/hashicorp_address.txt"
}