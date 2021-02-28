{{- /* yet-cert.tpl */ -}}
{{ with secret "pki/issue/nginx" "common_name=mongodb" "ip_sans=127.0.0.1" "ttl=380m" }}
{{ .Data.certificate }}
{{ .Data.private_key }}
{{ end }}