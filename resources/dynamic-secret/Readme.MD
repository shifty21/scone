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


mongod  --sslMode requireSSL --sslPEMKeyFile  mongodb.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 --sslCAFile mongodb-ca.crt --setParameter authenticationMechanisms=PLAIN,MONGODB-X509

create user
mongo --ssl --host localhost --sslCAFile mongodb-ca.crt --sslPEMKeyFile mongodb-client.pem sample.js

sample.js
db.getSiblingDB("\$external").runCommand(
    {
        createUser:"CN=jack",
        roles: [{"role" : "root","db" : "admin"}]
});

to test

db.getSiblingDB("$external").auth(
     {
         mechanism: "MONGODB-X509",
         user: "CN=jack"
     }
     );

mongod  --sslMode requireSSL --sslPEMKeyFile  mongodb.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 --sslCAFile mongodb-ca.crt --setParameter authenticationMechanisms=PLAIN,MONGODB-X509 --auth



mongod --auth  --sslMode requireSSL --sslPEMKeyFile  mongodb.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 
--sslCAFile mongodb-ca.crt  --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509 --sslAllowInvalidCertificates


mongod --auth  --tlsMode requireTLS --certificateKeyFile  mongodb.pem --CAFile mongodb-ca.crt --allowConnectionsWithoutCertificates --dbpath /usr/local/var/mongodb --bind_ip 127.0.0.1 --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509 

mongod --tlsMode requireTLS --tlsCertificateKeyFile  mongodb.pem --tlsCAFile mongodb-ca.crt --tlsAllowConnectionsWithoutCertificates --dbpath /usr/local/var/mongodb --bind_ip 127.0.0.1 --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509

mongo --ssl --host hostname.example.com --sslPEMKeyFile /etc/ssl/client.pem --sslCAFile /etc/ssl/ca.pem