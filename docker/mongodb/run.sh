#!/bin/bash
StartMongoWO_Auth="mongod --port 27017"
$StartMongoWO_Auth &
MONGO_PID=$!
printf "Ran mongodb %s\n" "$MONGO_PID"
sleep 5
mongo --port 27017  --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'myUserAdmin', pwd: 'abc123', roles: [{ role: 'userAdminAnyDatabase', db: 'admin' }] });"
createUser=$!
printf "Create user command %s \n" "$createUser"
kill "$MONGO_PID"
killed_mongodb=$!
printf "Killed non auth monodb exit code %s\n" "$killed_mongodb"
printf "Starting mongodb with auth enabled\n"
StartMongoW_Auth="mongod --auth --port 27017"
$StartMongoW_Auth
