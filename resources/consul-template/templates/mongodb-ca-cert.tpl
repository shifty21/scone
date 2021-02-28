{{- /* yet-cert.tpl */ -}}
{{ with secret "pki/issue/scone" "common_name=mongodb" "ip_sans=127.0.0.1" "ttl=380m" }}
{{ .Data.issuing_ca }}
{{ end }}