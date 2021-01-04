### dynamic-secret

Read dynamic secret role <readonly>
vault read database/creds/readonly

curl --header "X-Vault-Token: $VAULT_TOKEN" --request GET $VAULT_ADDR/v1/database/creds/demo-client | jq