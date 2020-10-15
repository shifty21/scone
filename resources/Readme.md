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
3. Start vault with 
   ```
   SCONE_CONFIG_ID=vault/dev SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault server -config /resources/vault/config.hcl
   ```
4. Start vault_initializer - this will intialize vault based on the type suggeted. For the 2nd senario mentioned above the private key for decryption would be inserted by CAS at /home/mykey.pem. This ensures that only authenticated application is able to intialize vault.
   ```
   SCONE_CONFIG_ID=vault-init/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault_init
   /root/go/bin/vault operator init -recovery-shares=1 -recovery-threshold=1
   ```
5. Once vault have successfully initialied add key to vault
   ```
   vault kv put secret/hello hashicorp="101 2nd St"
   vault kv put secret/hello hashicorp/street_address="101 2nd St"
   ```
6. Update demo-client's CAS session details to environment variables or configuration file(/resources/consul-template/config.hcl) for consul_template to pick up
7. Run consul-template to render the demo template (/resources/consul-template/find_address.tpl) -> 
   ```
   SCONE_CONFIG_ID=consul-template/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template -auth -config /resources/consul-template/config.hcl -once
   ```
8.  The template should be rendered in(/resources/consul-template/hashicorp_address.txt) if demo-client session is verifed by CAS.

### Additional commands
```
consul agent -dev
vault kv put secret/hello hashicorp="101 2nd St"
vault kv put secret/hello hashicorp/street_address="101 2nd St"
consul kv put hashicorp/street_address "101 2nd St"
consul kv delete -recurse vault/ 
```

## Run.sh
This script automates CAS session registration for consul-template, demo-client, vault and vault-init.

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

