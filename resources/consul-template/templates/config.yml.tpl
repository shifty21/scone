{{- with secret "database/creds/demo-client" }}
LeaseID : "{{ .Data}}"
username: "{{ .Data.username }}"
password: "{{ .Data.password }}"
database: "admin"
address: mongodb
{{- end }}
