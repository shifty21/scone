{{- /* yet-cert.tpl */ -}}
{{ with secret "pki/issue/nginx" "common_name=localhost" "format=pem" "ttl=380m" }}
{{ .Data.certificate }}
{{ .Data.private_key }}
{{ end }}