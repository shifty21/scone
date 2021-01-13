
### Build consul-template with and without gccgo
go build -o $GOPATH/bin/consul-template -v
go build -compiler gccgo -o /root/go/bin/consul-template -v

### Consul-template template re-rending with vault behaviour
LifetimeWatcher is a process for watching lifetime of a secret.
Please note that Vault does not support blocking queries. As a result, Consul Template will not immediately reload in the event a secret is changed as it does with Consul's key-value store. Consul Template will renew the secret with Vault's Renewer API. The Renew API tries to use most of the time the secret is good, renewing at around 90% of the lease time (as set by Vault).

To specify ttl for a specifiy secret 
vault kv put kv/my-secret ttl=30m my-value=s3cr3t
To update ttl of secret engine
vault secrets tune -default-lease-ttl=2m secret

### rc service
rc-status
rc-status --list
rc-update add {service-name} {run-level-name}
rc-service {service-name} start/stop/restart

rc-update add demo-client default