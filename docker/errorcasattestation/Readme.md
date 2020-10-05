
## Reproduce consul-template attestation error
Just run the docker-compose file it set cas and las environment flags

Consul-template cas files directory
```
>ls /resources/consul-template/
conf                         config.hcl                   config_back.hcl              find_address.tpl             session_client.yml           session_consul_template.yml

Register session -
Hash generation - SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template to update the hash 
cd /resources/consul-template/ && curl -k -s --cert conf/consul-template.crt --key conf/consul-template-key.key --data-binary @session_consul_template.yml -X POST https://$SCONE_CAS_ADDR:8081/session

RUN - SCONE_CONFIG_ID=consul_template1/consul1 SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template

```