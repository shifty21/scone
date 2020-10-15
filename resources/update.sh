#!/bin/bash
VAULT_SESSION=vault/session.yml
get_hash(){
OUTPUT=$(SCONE_HEAP=$2G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault --version 2>&1)
}

print() {
    HASH=$(echo "$OUTPUT" | grep "Enclave hash:" | awk -F ' ' '{print $3}')
    printf "Updating Hash of vault in %s %s for $2G \n" "$VAULT_SESSION" "$HASH"
}

update() {
    echo "version: \"0.3\"
name: vault$2
    
services:
    - name: dev$2
      image_name: sconecuratedimages/www2019:vault-1.5.0-alpine
      mrenclaves: [$HASH]
      command: /root/go/bin/vault server -config /resources/vault/config.hcl
      pwd: /
      environment:
        SCONE_MODE: hw
        SCONE_HEAP: $2G

security:
  attestation:
    tolerate: [debug-mode, outdated-tcb]
    ignore_advisories: "*"
        " > $VAULT_SESSION
}

if [ "$1" == "print" ]; then
    get_hash "$@"
    print "$@"
elif [ "$1" == "update" ]; then
    get_hash "$@"
    print "$@"
    update "$@"
else 
    get_hash "$@"
    print "$@"
    update "$@"
    sh run.sh register vault
fi