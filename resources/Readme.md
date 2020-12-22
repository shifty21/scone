# Vault in SCONE Demo 
Connection ssh -v yateenderk@141.76.44.85

For Demo all the files are under resources directory with a total of 4 components that needs to be registered to CAS
1. [vault](vault/Readme.md)
2. [vault-init](vault-init/Readme.md)  
3. [consul-template](consul-template/Readme.md) 
4. [demo-client](demo-client/Readme.md)
   
Other files 
1. run.sh and update.sh -  [Follow](resources/../Readme.md)

## Process
1. Run docker-compose.yaml - this will bring up vault, consul, cas and las instances.
2. login to vault container and run [run.sh](#runsh) to register all the binaries.
   ```
   sh run.sh register all
   ```
3. Start vault first only in case of shamir based initialization, in grpc based auto-initialization perform step 4 first.
   ```
   SCONE_CONFIG_ID=vault/dev SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault server -config /root/go/bin/resources/vault/config.hcl
   SCONE_CONFIG_ID=vault1/dev1 SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault server -config /root/go/bin/resources/vault/config_vanilla.hcl
   ```
4. Start vault_initializer - this will intialize vault based on the type suggeted. For the 2nd senario mentioned above the private key for decryption would be inserted by CAS at /home/mykey.pem. This ensures that only authenticated application is able to intialize vault.
   ```
   SCONE_CONFIG_ID=vault-init/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault-init grpc
   export VAULT_ADDR=http://127.0.0.1:8200
   /root/go/bin/vault operator init -recovery-shares=1 -recovery-threshold=1
   /root/go/bin/vault login s.HEE6qF2teSZa0aNNgQYaHqMY
   SCONE_CONFIG_ID=vault-init1/dev1 SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault-init scone
   SCONE_CONFIG_ID=vault-init1/dev1 SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault-init auto
   ```
5. Once vault have successfully initialied add key to vault
   ```
   /root/go/bin/vault secrets enable -path=secret/ kv
   /root/go/bin/vault kv put secret/hello hashicorp="101 2nd St"
   ```
6. Update demo-client's CAS session details to environment variables or configuration file(/root/go/bin/resources/consul-template/config.hcl) for consul_template to pick up
7. Run consul-template to render the demo template (/root/go/bin/resources/consul-template/find_address.tpl) -> 
   ```
   update /root/go/bin/resources/consul-template/config.hcl  with root token from vault
   
   SCONE_CONFIG_ID=consul-template/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template -auth -config /root/go/bin/resources/consul-template/config.hcl -once
   ```
8.  The template should be rendered in(/root/go/bin/resources/consul-template/hashicorp_address.txt) if demo-client session is verifed by CAS.

### Additional commands
```
consul agent -dev
consul kv put hashicorp/street_address "101 2nd St"
consul kv delete -recurse vault/ 
```

## Run.sh
This script automates CAS session registration for consul-template, demo-client, vault and vault-init.
It calculates mrenclave for everyone and update the session files
To use run.sh
```
run.sh $1 $2
Where $1 have register or check value only
$2 can be among 4 values - vault vault-init consul-template demo-client
```
Using update.sh
```

```
### Encryption keys generation
```
openssl genrsa -out mykey.pem 1024
openssl rsa -in mykey.pem -pubout > mykey.pub
openssl rsa -in mykey.pem -pubout -out pubkey.pem
```

### Optimizations
#### Vault optimization
https://www.vaultproject.io/docs/configuration/storage/consul#consistency_mode
Vault supports using 2 of the 3 Consul Consistency Modes.

#### consul optimization
Consul is write limited by disk I/O and read limited by CPU.
https://www.consul.io/docs/install/performance
For a read-heavy workload, configure all Consul server agents with the allow_stale DNS option, or query the API with the stale consistency mode. By default, all queries made to the server are RPC forwarded to and serviced by the leader. By enabling stale reads, any server will respond to any query, thereby reducing overhead on the leader. Typically, the stale response is 100ms or less from consistent mode but it drastically improves performance and reduces latency under high load.

#### Enterprise
Read-heavy clusters may take advantage of the enhanced reading feature (Enterprise) for better scalability. This feature allows additional servers to be introduced as non-voters. Being a non-voter, the server will still participate in data replication, but it will not block the leader from committing log entries.

### Transit backend
The primary use case for transit is to encrypt data from applications while still storing that encrypted data in some primary data store. This relieves the burden of proper encryption/decryption from application developers and pushes the burden onto the operators of Vault.
https://www.vaultproject.io/docs/secrets/transit
vault secrets enable transit
create enc key
vault write -f transit/keys/my-key
write secret 
vault write transit/encrypt/my-key plaintext=$(base64 <<< "my secret data")
decrypt
vault write transit/decrypt/my-key ciphertext=vault:v1:8SDd3WHDOjf7mq69CyCqYjBXAiQQAVZRkFM13ok481zoCmHnSeDX9vyf7w==
base64 --decode <<< "bXkgc2VjcmV0IGRhdGEK"

### Rotate Key 
vault write -f transit/keys/my-key/rotate
future enc uses new key, older ones are used to decrypt old data
rewrap with new key
vault write transit/rewrap/my-key ciphertext=vault:v1:8SDd3WHDOjf7mq69CyCqYjBXAiQQAVZRkFM13ok481zoCmHnSeDX9vyf7w==

creating tokens via api - initialize vault and generate this token which will be used by others
https://www.vaultproject.io/api-docs/auth/token

### Dynamic Secrets
For creating dynamic secrets you need to consider 3 things secrets (the dynamically created secret), Authentication (A token other than root token to access the secret), Policy(Ties together secret and authenticated users capability)
https://www.vaultproject.io/docs/secrets/databases
Policies types
* Roles - Controls the tables to which user has access  - database/roles/<role name>
   Paramters 
   * db_name
   * creation_statements - defines templated role statement that will be filled upon user request with credentials
   * revocation_statements - revokes access after ttl
   * default_ttl
   * max_ttl
   usage - vault write database/roles/<dbrole name>
 * Connection - databse/config/<wizard> - manages root access to database
   * plugin_name
   * allowed_roles
   * connection_url - ex -"{{username}}:{{password}}@tcp(127.0.0.1:3306)/"
   * username -  initial credential 
   * password
   Usage - vault write database/config/wizard <rest of paratmers>
   vault read database/creds/db-app
* Read policy
  * database/creds/db-app
* Password policy - sys/policies/password/example - https://learn.hashicorp.com/tutorials/vault/database-secrets
  * Set of rules to which password should adher to 


#### Extras

https://www.vaultproject.io/docs/configuration/listener/tcp

telemetry configuration
https://www.vaultproject.io/docs/configuration/telemetry.html

#### statis roles
   Creates static roles where only the password is rotated

#### Mongodb dynamic role
   https://www.vaultproject.io/docs/secrets/databases/mongodb
   https://www.vaultproject.io/api-docs/secret/databases/mongodb#revocation_statements
   Integration - https://learn.hashicorp.com/tutorials/vault/application-integration

#### Create token for access
https://www.vaultproject.io/api/auth/token
Token is a auth method in vault which is enabled by default
curl \
    --header "X-Vault-Token: ..." \
    --request POST \
    --data @payload.json \
    http://127.0.0.1:8200/v1/auth/token/create
    


 docker-compose rm -f -s -v cas
 1. vault
    1. register vault and vault-init, update hash of vault-init and predecessor hash
    2. update session_auto.yaml and register vault-init auto and demo-parent1
   
cd /root/go/bin/resources/demo-client && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session1.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
cd /root/go/bin/resources/vault && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
cd /root/go/bin/resources/vault-init && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_auto.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
cd /root/go/bin/resources/vault-init && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session

cd /root/go/bin/resources/demo-client && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session

cd /root/go/bin/resources/consul-template && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session


cd /root/go/bin && SCONE_CONFIG_ID=vault-init/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault-init grpc
cd /root/go/bin && SCONE_CONFIG_ID=vault/dev SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault server -config /root/go/bin/resources/vault/config.hcl
cd /root/go/bin && SCONE_CONFIG_ID=vault-init-auto/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault-init auto
 1. demo-client
    1. 