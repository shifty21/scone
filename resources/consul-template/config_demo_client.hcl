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
cert = "/root/go/bin/resources/demo-client/conf/client.crt"
key = "/root/go/bin/resources/demo-client/conf/client-key.key"
session_name = "demo-client"
session_file = "/root/go/bin/resources/demo-client/session.yml"
test_updated_session_file = "/root/go/bin/resources/demo-client/test_updated_session.yml"
predecessor_hash_file = "/root/go/bin/resources/demo-client/predecessor_hash.yaml"
}

template {
    source      = "resources/consul-template/templates/config.yml.tpl"
    destination = "resources/consul-template/templates/config.yml"
}
