## Vault 

#### Vault vanilla initialization and unseal via terminal

```console
 /root/go/bin/vault operator init
 /root/go/bin/vault status
 /root/go/bin/vault operator unseal
 /root/go/bin/vault status
 /root/go/bin/vault secrets enable -path=kv kv
 /root/go/bin/vault secrets enable kv
 /root/go/bin/vault secrets list
 /root/go/bin/vault kv put kv/hello hashicorp="101 2nd St"
 /root/go/bin/vault kv get kv/hello
```
#### [Initialize vault with http](vault_init/README.md)
#### [Vault codebase working](codebase/Readme.md)