start
export MONGO_INITDB_ROOT_USERNAME=root
export MONGO_INITDB_ROOT_PASSWORD=rootPassXXX
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb
login
mongo --port 27017  --authenticationDatabase "jack" -u "hammer" -p

mongo --port 27017  --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'root',pwd: 'p\"gcup\"7y4+v', roles: [{ role: 'root', db: 'admin' }] });"
p"gcup"7y4+v
["/usr/bin/mongo --port 27017","--eval \"db = db.getSiblingDB('admin'); db.createUser({ user: 'myUserAdmin', pwd: '$$SCONE::db_password$$', roles: [{ role: 'userAdminAnyDatabase', db: 'admin' }] });\"]

curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_mongodb.yml -X POST https://cas:8081/session
curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_register_user.yml -X POST https://cas:8081/session
## get session
curl -k -s --cert conf/client.crt --key conf/client-key.key https://$SCONE_CAS_ADDR:8081/session/mongodb
curl -k -s --cert conf/client.crt --key conf/client-key.key https://$SCONE_CAS_ADDR:8081/session/mongodb-user-setup


SCONE_CONFIG_ID=mongodb-setup/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongo /home/payload.json
vault-dynamic-secret

SCONE_CONFIG_ID=vault-dynamic-secret/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault write database/config/admin \
    plugin_name=mongodb-database-plugin \
    allowed_roles="demo-client" \
    connection_url="mongodb://{{username}}:{{password}}@mongodb:27017/admin" \
    username="myUserAdmin" \
    password="$$SCONE::mongodb_password$$"

SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongo
mongodb
SCONE_CONFIG_ID=mongodb-user-setup/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongo /home/payload.json
SCONE_CONFIG_ID=mongodb/init SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongod
SCONE_CONFIG_ID=mongodb/auth SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongod
SCONE_CONFIG_ID=env-print/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/env-print