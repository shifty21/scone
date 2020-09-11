#!/bin/bash
#register vault
cd /resources/ && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_vault.yml -X POST https://$SCONE_CAS_ADDR:8081/session
#register vault_initializer
cd /resources/ && curl -k -s --cert conf/client_back.crt --key conf/client-key_back.key --data-binary @session_vault_init.yml -X POST https://$SCONE_CAS_ADDR:8081/session