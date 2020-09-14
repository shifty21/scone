#!/bin/bash
#vault
#vault kv put  secret/hello hashicorp="101 2nd St"
#register vault
cd /resources/ && curl -k -s --cert conf/vault.crt --key conf/vault-key.key --data-binary @session_vault.yml -X POST https://$SCONE_CAS_ADDR:8081/session
#register vault_initializer
cd /resources/ && curl -k -s --cert conf/vault-initializer.crt --key conf/vault-initializer-key.key --data-binary @session_vault_init.yml -X POST https://$SCONE_CAS_ADDR:8081/session
#vault 
#cd /resources/ && vault server -config=config.hcl
#SCONE_CONFIG_ID=vault_demo/vault vault server -config=config.hcl




#register consul-template
cd /resources/consul-template && curl -k -s --cert conf/consul-template.crt --key conf/consul-template-key.key --data-binary @session_consul_template.yml -X POST https://$SCONE_CAS_ADDR:8081/session
#register vault_initializer
cd /resources/consul-template && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_client.yml -X POST https://$SCONE_CAS_ADDR:8081/session

#consul-template
#SCONE_CONFIG_ID=consul_template/consul_template consul-template -auth -config config_back.hcl -once