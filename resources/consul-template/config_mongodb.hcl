reload_signal = "SIGHUP"
kill_signal = "SIGINT"
max_stale = "4m"
log_level = "info"
vault {
    address = "http://vault:8200"
    renew_token = false
}
cas {
        get_session_api = "https://cas:8081/session"
        cert = "/root/go/bin/resources/mongodb/conf/client.crt"
        key = "/root/go/bin/resources/mongodb/conf/client-key.key"
        session_name = "mongodb"
        session_file = "/root/go/bin/resources/mongodb/session.yml"
        test_updated_session_file = "/root/go/bin/resources/mongodb/mongodb_updated_session.yml"
        predecessor_hash_file = "/root/go/bin/resources/mongodb/predecessor_hash.yaml"
}

template {
  source      = "/root/go/bin/resources/consul-template/templates/mongodb-ca-cert.tpl"
  destination = "/root/go/bin/resources/consul-template/templates/mongodb-ca.crt"
  perms       = "0600"
}

template {
  source      = "/root/go/bin/resources/consul-template/templates/mongodb-pem-key.tpl"
  destination = "/root/go/bin/resources/consul-template/templates/mongodb.pem"
}