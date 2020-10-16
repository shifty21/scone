{{ with secret "secret/hello" }}
{{ .Data.hashicorp }}
{{ end }}