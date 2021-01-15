{{- /* yet-cert.tpl */ -}}
{{ with secret "pki/issue/nginx" "common_name=tu-dresden.de" "format=pem" "ttl=60m" }}
{{ .Data.private_key }}
{{ end }}