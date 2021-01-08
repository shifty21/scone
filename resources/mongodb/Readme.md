start
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb
login
mongo --port 27017  --authenticationDatabase "jack" -u "hammer" -p

mongo --port 27017  --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'root', roles: [{ role: 'root', db: 'admin' }] });"