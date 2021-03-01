start
export MONGO_INITDB_ROOT_USERNAME=root
export MONGO_INITDB_ROOT_PASSWORD=rootPassXXX
mongod --auth --port 27017 --dbpath /usr/local/var/mongodb
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
    db.getSiblingDB("\$external").runCommand(
    {
        createUser:"CN=mongodb",
        roles: [{"role" : "userAdminAnyDatabase","db" : "admin"}]
    });

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

4.2
mongod --auth --tlsMode requireTLS --tlsCertificateKeyFile  mongodb.pem --tlsCAFile mongodb-ca.crt --tlsAllowConnectionsWithoutCertificates --dbpath /usr/local/var/mongodb --bind_ip 127.0.0.1 --setParameter authenticationMechanisms=MONGODB-X509,SCRAM-SHA-1
3.6
mongod --auth  --sslMode requireSSL --sslPEMKeyFile  mongodb.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 --sslCAFile mongodb-ca.crt  --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509 --sslAllowInvalidCertificates


vault write database/config/admin \
    plugin_name=mongodb-database-plugin allowed_roles="demo-client" connection_url="mongodb://@localhost:27017/admin?ssl=true" tls_certificate_key=@mongodb-client.pem tls_ca=@mongodb-ca.crt



consul-template -config /root/go/bin/resources/consul-template/config_mongodb.hcl  -once -render-to-disk

mongod -auth --sslMode requireSSL --sslAllowConnectionsWithoutCertificates --sslAllowInvalidCertificates --sslPEMKeyFile  mongo.pem --dbpath /usr/local/var/mongodb  --bind_ip 127.0.0.1 --sslCAFile issue_ca.crt  --setParameter authenticationMechanisms=SCRAM-SHA-1,MONGODB-X509

To test the connection via client 
mongo --port 27017 --ssl  --sslAllowInvalidCertificates --authenticationDatabase "admin" -u  "" -p

mongo --port 27017 --ssl  --sslAllowInvalidCertificates --authenticationDatabase "admin" -u  "" -p

mongo --eval 'db.getSiblingDB("\$external").runCommand({createUser:"CN=mongodb",roles: [{"role" : "userAdminAnyDatabase","db" : "admin"}]});'