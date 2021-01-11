mongod --port 27017 --dbpath /usr/local/var/mongodb
/usr/local/etc/mongod.conf
log location on mac /usr/local/var/log/mongodb/mongo.log
brew services list
brew services start
create user

show dbs
use admin
db.createUser(
  {
    user: "myUserAdmin",
    pwd: "w(`tF2Q'$xda",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)

mongo --port 27017  --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'myUserAdmin', pwd: 'abc123', roles: [{ role: 'userAdminAnyDatabase', db: 'admin' }] });"

mongo --port 27017  --authenticationDatabase "admin" -u "myUserAdmin" -p "abc123"
SCONE_CONFIG_ID=mongodb-setup/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongo /home/payload.json


./vault write database/config/admin plugin_name=mongodb-database-plugin allowed_roles="demo-client" connection_url="mongodb://{{username}}:{{password}}@mongodb:27017/admin" username="myUserAdmin" password="w(`tF2Q'$xda"