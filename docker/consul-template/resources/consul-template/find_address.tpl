{{ with secret "secret/hello" }}
{{ .Data.data.hashicorp }}
{{ end }}
