{{- with secret "database/creds/demo-client" }}
username: "{{ .Data.username }}"
password: "{{ .Data.password }}"
database: "admin"
address: mongodb
{{- end }}
