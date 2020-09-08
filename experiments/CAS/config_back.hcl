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
    token = "s.qSPYXpulHuFgEd2dxmp3o1zs"
    renew_token = false
}
cas {
get_session_api = "https://127.0.0.1:8081/session/"
cert = "/Users/abhishekchoudhary/CAS/conf/client_back.crt"
key = "/Users/abhishekchoudhary/CAS/conf/client-key_back.key"
session_name = "blender"
}

template {
    source      = "/Users/abhishekchoudhary/CAS/find_address.tpl"
    destination = "/Users/abhishekchoudhary/CAS/hashicorp_address.txt"
}
