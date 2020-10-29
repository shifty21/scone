SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault --version

cd /root/go/bin/resources/demo-client && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session1.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
#update hash
cd /root/go/bin/resources/vault-init && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_auto.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
#update hash
cd /root/go/bin/resources/vault-init && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session_scone.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
