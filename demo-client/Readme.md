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
    user: "admin",
    pwd: "password",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)
mongo --port 27017  --authenticationDatabase "admin" -u "admin" -p "password"