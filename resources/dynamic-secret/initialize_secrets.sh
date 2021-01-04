
#update hash

# export VAULT_ADDR=http://127.0.0.1:8200
# ./vault login s.YpPpLXLypRaGuNWQTEOQtiZM
./vault secrets enable database

./vault write database/roles/demo-client \
    db_name=admin \
    creation_statements='{ "db": "admin", "roles": [{ "role": "readWrite" }, {"role": "readWrite", "db": "dev"},{"role": "readWrite", "db": "test"}, {"role": "read", "db": "production"}] }' \
    default_ttl="2m" \
    max_ttl="6m"

./vault write database/config/admin \
    plugin_name=mongodb-database-plugin \
    allowed_roles="demo-client" \
    connection_url="mongodb://{{username}}:{{password}}@mongodb:27017/admin" \
    username="myUserAdmin" \
    password="abc123"

#for policy
./vault policy write demo-client /root/go/bin/resources/dynamic-secret/credential-policy.hcl

#for address template 
# vault secrets enable -path=secret/ kv
# vault kv put secret/hello hashicorp="101 2nd St"