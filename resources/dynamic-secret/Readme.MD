### dynamic-secret

Read dynamic secret role <readonly>
vault read database/creds/readonly

curl --header "X-Vault-Token: $VAULT_TOKEN" --request GET $VAULT_ADDR/v1/database/creds/demo-client | jq

via CURL
curl \
    --header "X-Vault-Token: ..." \
    --request POST \
    --data @payload.json \
    http://127.0.0.1:8200/v1/database/config/admin

{
  "plugin_name": "mysql-database-plugin",
  "allowed_roles": "readonly",
  "connection_url": "{{username}}:{{password}}@tcp(127.0.0.1:3306)/",
  "username": "vaultuser",
  "password": "secretpassword"
}