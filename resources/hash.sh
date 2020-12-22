SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault --version
export SCONE_CAS_ADDR=127.0.0.1
register vault-init
cd /root/go/bin/resources/vault-init && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
#update hash
consul kv delete -recurse vault/ 
demo client
cd /root/go/bin/resources/demo-client && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
export VAULT_ADDR=http://127.0.0.1:8200
vault login s.9dPa8pbheYLEeGByB7Gqvg0r
vault secrets enable database

vault write database/roles/demo-client \
    db_name=admin \
    creation_statements='{ "db": "admin", "roles": [{ "role": "readWrite" }, {"role": "readWrite", "db": "dev"},{"role": "readWrite", "db": "test"}, {"role": "read", "db": "production"}] }' \
    default_ttl="2m" \
    max_ttl="6m"

vault write database/config/admin \
    plugin_name=mongodb-database-plugin \
    allowed_roles="demo-client" \
    connection_url="mongodb://{{username}}:{{password}}@demo-client:27017/admin" \
    username="myUserAdmin" \
    password="abc123"

#for policy
vault policy write demo-client resources/dynamic-secret/credential-policy.hcl

#for address template 
vault secrets enable -path=secret/ kv
vault kv put secret/hello hashicorp="101 2nd St"


username: "v-root-demo-client-qTchT4AEAnL4NUJD32Ss-1608223221"
password: "A1a-9g3mzpNCGq8jqCLh"
database: "admin"

consul-template -config /root/go/bin/resources/consul-template/config.hcl -once
#update hash
cd /root/go/bin/resources/vault-init && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_scone.yml -X POST https://"$SCONE_CAS_ADDR":8081/session

#Get a cas session
curl -k -s --cert conf/client.crt --key conf/client-key.key https://$SCONE_CAS_ADDR:8081/session/blender
#https://github.com/emdem/consul-cluster-compose