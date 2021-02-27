{{- /* yet-cert.tpl */ -}}
{{ with secret "pki/issue/nginx" "common_name=tu-dresden.de" "ip_sans=127.0.0.1" "ttl=200h" }}
{{ .Data.certificate }}
{{ .Data.issuing_ca }}
{{ end }}