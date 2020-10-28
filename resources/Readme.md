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

* If you dont have a root token but the keys are available 
   https://learn.hashicorp.com/tutorials/vault/generate-root
   vault operator unseal and provide unseal keys
Recover keys can be used to unseal only for rekeying and generating root token again
https://www.vaultproject.io/api-docs/system/init#recovery_shares

Rekeying process