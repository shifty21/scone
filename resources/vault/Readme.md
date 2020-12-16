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

### Vault get root token with unseal keys
https://learn.hashicorp.com/tutorials/vault/generate-root
vault operator unseal and provide unseal keys
Recover keys can be used to unseal only for rekeying and generating root token again
https://www.vaultproject.io/api-docs/system/init#recovery_shares

```
export VAULT_ADDR=http://127.0.0.1:8200
/root/go/bin/vault operator generate-root -init -> this will give OTP for decoding root token
/root/go/bin/vault operator generate-root -> enter unseal key here one at a time
/root/go/bin/vault operator generate-root -decode=RUoKWDwJPj54PgETCiYiJSwlAR8+Pw0yUSM -otp=6dMjrsmv7YQK8DqvnjGKnUNf0a
/root/go/bin/vault login s.G2NzSHOgPX2bSSBOFTPjCTaB
```