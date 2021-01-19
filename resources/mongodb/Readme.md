start
export MONGO_INITDB_ROOT_USERNAME=root
export MONGO_INITDB_ROOT_PASSWORD=rootPassXXX
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb
login

password:A1a-7R7jwF7NcxNKyjzH username:v-root-demo-client-ohPdgESyLYtTrsmo7yc4-1610733467
mongo --port 27017  --authenticationDatabase "admin" -u "v-root-demo-client-ohPdgESyLYtTrsmo7yc4-1610733467" -p

    db.getSiblingDB("$external").auth(
    {
        mechanism: "MONGODB-X509",
        user: "CN=jack"
    }
    );

     db.getSiblingDB("admin").auth({
        mechanism: "MONGODB-X509",
        user: "v-root-demo-client-hfy1PfEFO0nJ8Fdd27QS-1610729501",
        password: "A1a-ZHEEAPRnuUpZK9ca"
    });



    db.getSiblingDB("$external");
    db.createUser({user: "CN=127.0.0.1",
               roles: [{"role" : "userAdminAnyDatabase","db" : "admin"}]
    });
    db.getSiblingDB("\$external").runCommand(
    {
        createUser:"CN=jack",
        roles: [{"role" : "userAdminAnyDatabase","db" : "admin"}]
    });
    db.getSiblingDB("\$external").runCommand(
    {
        createUser:"CN=127.0.0.1",
        roles: [{"role" : "userAdminAnyDatabase","db" : "admin"}]
    });
    mechanisms: ["MONGODB-X509"],
);
mongo --port 27017  --eval "db = db.getSiblingDB('admin'); db.createUser({ user: 'myUserAdmin', pwd: 'jackhammer', roles: [{ role: 'userAdminAnyDatabase', db: 'admin' }] });"
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
SCONE_CONFIG_ID=mongodb-user-setup/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongo /home/payload.js
SCONE_CONFIG_ID=mongodb/init SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongod --dbpath /usr/local/var/mongodb
SCONE_CONFIG_ID=mongodb/auth SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /usr/bin/mongod -auth --bind_ip=0.0.0.0 --dbpath /usr/local/var/mongodb
SCONE_CONFIG_ID=env-print/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/env-print

3.6
mongo --ssl --host localhost --sslCAFile mongodb-ca.crt --sslPEMKeyFile mongodb-client.pem
4.2
mongo --ssl --host localhost --tlsCAFile mongodb-ca.crt --tlsCertificateKeyFile mongodb-client.pem
both?
    db.getSiblingDB("$external");
    db.createUser({user: "CN=127.0.0.1",
               roles: [{"role" : "userAdminAnyDatabase","db" : "admin"}]
    });
mongo --port 27017 --ssl --sslAllowInvalidCertificates  --authenticationDatabase "admin" -u "v-root-demo-client-If6DXtGtDyW99R9xpgS9-1610739799" -p

4.2
mongod --auth --tlsMode requireTLS --tlsCertificateKeyFile  mongodb.pem --tlsCAFile mongodb-ca.crt --tlsAllowConnectionsWithoutCertificates --dbpath /usr/local/var/mongodb --bind_ip 127.0.0.1 --setParameter authenticationMechanisms=MONGODB-X509,SCRAM-SHA-1
3.6
mongod --auth  --sslMode requireSSL --sslPEMKeyFile  mongodb.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 --sslCAFile mongodb-ca.crt  --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509 --sslAllowInvalidCertificates


vault write database/config/admin \
    plugin_name=mongodb-database-plugin allowed_roles="demo-client" connection_url="mongodb://@localhost:27017/admin?ssl=true" tls_certificate_key=@mongodb-client.pem tls_ca=@mongodb-ca.crt


    mongo --port 27017 --ssl  --sslAllowInvalidCertificates --authenticationDatabase "admin" -u  "v-root-demo-client-As7igycgBpWm4YvaWJqV-1611019219" -p

    mongod -auth  --sslMode requireSSL --sslAllowConnectionsWithoutCertificates --sslAllowInvalidCertificates --sslPEMKeyFile  mongo.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 --sslCAFile issue_ca.crt  --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509  --sslDisabledProtocols 'none'