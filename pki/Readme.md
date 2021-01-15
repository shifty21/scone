## helping apis

read role

curl \
    --header "X-Vault-Token: ..." \
    http://127.0.0.1:8200/v1/pki/roles/my-role

Delete role 
curl \
    --header "X-Vault-Token: ..." \
    --request DELETE \
    http://127.0.0.1:8200/v1/pki/roles/my-role


Delete root

curl \
    --header "X-Vault-Token: ..." \
    --request DELETE \
    http://127.0.0.1:8200/v1/pki/root

Revoke certificate 
{
  "serial_number": "39:dd:2e..."
}
curl \
    --header "X-Vault-Token: ..." \
    --request POST \
    --data @payload.json \
    http://127.0.0.1:8200/v1/pki/revoke