## Resources
This directory contains all the configuration and keys for consul-template, demo-client, vault and vault-init to register sesssions to CAS.

To use run.sh
```
run.sh $1 $2

Where $1 can be register or check
$2 can be among 4 values - vault vault-init consul-template demo-client-session

```

To insert value in vault 
```
vault kv put  secret/hello hashicorp="101 2nd St"
vault kv put secret/hello hashicorp/street_address="101 2nd St"
consul kv put hashicorp/street_address "101 2nd St"

```
To render consul-template 
```
SCONE_CONFIG_ID=consul_template/consul_template consul-template -auth -config config_back.hcl -once

```


Running vault with config file and scone
```
vault server -config=config.hcl

With CAS
SCONE_CONFIG_ID=vault_demo/vault vault server -config=config.hcl

Dynamic linking
SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 ./vault server -dev 
```

### Empty consul of vault config after initialization

Make sure you restart vault once this is done
```
consul kv delete -recurse vault/
```
### Encryption keys generation

To generate pub-pvt key pair

```
openssl genrsa -out mykey.pem 1024
openssl rsa -in mykey.pem -pubout > mykey.pub
openssl rsa -in mykey.pem -pubout -out pubkey.pem
```

