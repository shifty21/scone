### dynamic-secret

Read dynamic secret role <readonly>
vault read database/creds/readonly

curl --header "X-Vault-Token: $VAULT_TOKEN" --request GET $VAULT_ADDR/v1/database/creds/demo-client | jq

via CURL
curl --header "X-Vault-Token: s.SCWsjXizDGNDmiPHrvP6Pclx" --request POST \
    --data @payload.json \
    http://127.0.0.1:8200/v1/database/config/admin

{
  "plugin_name": "mongodb-database-plugin",
  "allowed_roles": "demo-client",
  "connection_url": "mongodb://{{username}}:{{password}}@mongodb:27017/admin",
  "username": "myUserAdmin",
  "password": "R/$QNreid\=0"
}