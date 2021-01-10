start
export MONGO_INITDB_ROOT_USERNAME=root
export MONGO_INITDB_ROOT_PASSWORD=rootPassXXX
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb
login
mongo --port 27017  --authenticationDatabase "jack" -u "hammer" -p

mongo --port 27017  --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'root', roles: [{ role: 'root', db: 'admin' }] });"

["/usr/bin/mongo --port 27017","--eval \"db = db.getSiblingDB('admin'); db.createUser({ user: 'myUserAdmin', pwd: '$$SCONE::db_password$$', roles: [{ role: 'userAdminAnyDatabase', db: 'admin' }] });\"]

curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://cas:8081/session
SCONE_CONFIG_ID=mongodb-setup/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongo /home/payload.json
vault-dynamic-secret

SCONE_CONFIG_ID=vault-dynamic-secret/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault write database/config/admin \
    plugin_name=mongodb-database-plugin \
    allowed_roles="demo-client" \
    connection_url="mongodb://{{username}}:{{password}}@mongodb:27017/admin" \
    username="myUserAdmin" \
    password="$$SCONE::mongodb_password$$"