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
    pwd: "abc123",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb --bind_ip=0.0.0.0
mongo --port 27017  --authenticationDatabase "admin" -u "myUserAdmin" -p "abc123"

commands
mkdir -p /usr/local/var/mongodb

rc-update reload 

/etc/init.d/demo_client
openrc add demo_client default
rc-service demo-client reload