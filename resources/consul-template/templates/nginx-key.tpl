{{- /* yet-cert.tpl */ -}}
{{ with secret "pki/issue/scone" "common_name=tu-dresden.de" "ip_sans=127.0.0.1" "ttl=380m" }}
{{ .Data.private_key }}
{{ end }}