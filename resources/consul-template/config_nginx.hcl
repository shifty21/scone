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
cert = "/root/go/bin/resources/nginx/conf/client.crt"
key = "/root/go/bin/resources/nginx/conf/client-key.key"
session_name = "nginx"
session_file = "/root/go/bin/resources/nginx/session.yml"
test_updated_session_file = "/root/go/bin/resources/nginx/test_updated_session.yml"
predecessor_hash_file = "/root/go/bin/resources/nginx/predecessor_hash.yaml"
}

template {
  source      = "/root/go/bin/resources/consul-template/templates/nginx-cert.tpl"
  destination = "/root/go/bin/resources/consul-template/templates/nginx.crt"
  perms       = "0600"
}

template {
  source      = "/root/go/bin/resources/consul-template/templates/nginx-key.tpl"
  destination = "/root/go/bin/resources/consul-template/templates/nginx.key"
}
