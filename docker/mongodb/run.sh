#!/bin/bash
# StartMongoWO_Auth="mongod --port 27017"
# $StartMongoWO_Auth &
# MONGO_PID=$!
# printf "Ran mongodb %s\n" "$MONGO_PID"
# sleep 5
# cd /root/go/bin/resources/mongodb && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
SCONE_CONFIG_ID=mongodb/setup SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongod --dbpath /usr/local/var/mongodb
SCONE_CONFIG_ID=mongodb/setup SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongod --dbpath /usr/local/var/mongodb --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'myUserAdmin', pwd: '$$SCONE::db_password$$', roles: [{ role: 'userAdminAnyDatabase', db: 'admin' }] });"
createUser=$!
printf "Create user command %s \n" "$createUser"
kill "$MONGO_PID"
killed_mongodb=$!
printf "Killed non auth monodb exit code %s\n" "$killed_mongodb"
printf "Starting mongodb with auth enabled\n"
StartMongoW_Auth="mongod --auth --port 27017"
$StartMongoW_Auth
