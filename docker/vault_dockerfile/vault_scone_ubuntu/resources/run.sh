#!/bin/bash
#vault
#vault kv put  secret/hello hashicorp="101 2nd St"
#register vault
cd /resources/ && curl -k -s --cert conf/vault.crt --key conf/vault-key.key --data-binary @session_vault.yml -X POST https://$SCONE_CAS_ADDR:8081/session
# Get vault session - curl -k -s --cert conf/vault.crt --key conf/vault-key.key https://$SCONE_CAS_ADDR:8081/session/vault_demo
#register vault_initializer
cd /resources/ && curl -k -s --cert conf/vault-initializer.crt --key conf/vault-initializer-key.key --data-binary @session_vault_init.yml -X POST https://$SCONE_CAS_ADDR:8081/session
# Get vault_init session - curl -k -s --cert conf/vault-initializer.crt --key conf/vault-initializer-key.key https://$SCONE_CAS_ADDR:8081/session/vault_init_demo

#register consul-template
cd /resources/consul-template && curl -k -s --cert conf/consul-template.crt --key conf/consul-template-key.key --data-binary @session_consul_template.yml -X POST https://$SCONE_CAS_ADDR:8081/session
# Get session client -  curl -k -s --cert conf/consul-template.crt --key conf/consul-template-key.key https://$SCONE_CAS_ADDR:8081/session/consul_template
#register client
cd /resources/consul-template && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_client.yml -X POST https://$SCONE_CAS_ADDR:8081/session
# Get session client -  curl -k -s --cert conf/client.crt --key conf/client-key.key https://$SCONE_CAS_ADDR:8081/session/client
#consul-template
#SCONE_CONFIG_ID=consul_template/consul_template consul-template -auth -config config_back.hcl -once


#Running vault with config file and scone
#cd /resources/ && vault server -config=config.hcl
#SCONE_CONFIG_ID=vault_demo/vault vault server -config=config.hcl
#SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 ./vault server -dev 

#empty consul of vault config 
#consul kv delete -recurse vault/

consul-template
SCONE_CONFIG_ID=consul_template/consul SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template
