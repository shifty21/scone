
## Reproduce consul-template attestation error
Just run the docker-compose file it set cas and las environment flags

```
cd /resources/consul-template

Hash generation - SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template --version

curl -k -s --cert conf/cleint.crt --key conf/client-key.key --data-binary @session.yml -X POST https://$SCONE_CAS_ADDR:8081/session

RUN - SCONE_CONFIG_ID=consul-template/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template -auth -config /resources/consul-template/config.hcl -once

```