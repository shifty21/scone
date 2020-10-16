#!/bin/bash
SESSION_LIST="vault vault-init consul-template demo-client"
 
set_plugin() {
    if [ "$2" == "vault" ]; then
        RUN_SESSION="vault"
    elif [ "$2" == "vault-init" ]; then
        RUN_SESSION="vault-init"
    elif [ "$2" == "consul-template" ]; then
        RUN_SESSION="consul-template"
    elif [ "$2" == "demo-client" ]; then
        RUN_SESSION="demo-client"
    else 
        RUN_SESSION=${SESSION_LIST}
    fi
}

get_mrenclave() {
    if [ "$1" == "vault" ]; then
        COMMAND="SCONE_HEAP=8G SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault --version"
    elif [ "$1" == "vault-init" ]; then
        COMMAND="SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/vault-init --version"
    elif [ "$1" == "consul-template" ]; then
        COMMAND="SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template --version"
    fi
    echo "$COMMAND"
}

get_hash() {
    OUTPUT=$(eval "$1" 2>&1)
    HASH=$(echo "$OUTPUT" | grep "Enclave hash:" | awk -F ' ' '{print $3}')
    echo "$HASH"
}

register(){
    
    for plugin in ${RUN_SESSION}; do
        printf "\n=============================================================\n"
        local hash
        local cmd
        cmd=$(get_mrenclave "${plugin}")
        printf "Calculating hash for [%s] with command=[%s]\n" "${plugin}" "$cmd"
        hash=$(get_hash "$cmd")
        printf "Updating hash for [%s] with value=[%s]\n "  "${plugin}" "$hash"
        printf "\n=============================================================\n"
        if [ "$hash" == "" ]; then
            hash="56d37ffa65f3aa0e83aad166aa2b339b2bdda20254617c4b3b8e3fa794c41991"
        fi
        replace_hash="sed -i 's/\[[A-Za-z0-9]*\]\/*/\[$hash\]/' session.yml"
        
        cd /root/go/bin/resources/"${plugin}" && eval "$replace_hash"
        cat /root/go/bin/resources/"${plugin}"/session.yml
        printf '\nRegistring CAS session for %s\n=============================================================\n' "${plugin}"
        cd /root/go/bin/resources/"${plugin}" && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
        printf "\n============================================================= \n"
    done
}

check(){
    for plugin in ${RUN_SESSION}; do
        printf '\nChecking CAS session for %s\n=============================================================\n' "${plugin}"
        cd /root/go/bin/resources/"${plugin}" && curl -k -s --cert conf/client.crt --key conf/client-key.key https://"$SCONE_CAS_ADDR":8081/session/"${plugin}"
        printf "\n============================================================= \n"
    done
}

help (){
        printf '\nThis scripts register/checks session to CAS.
The first parameter can be - register or check.
The second parameter take value among %s or empty for all. \n' "${SESSION_LIST}"       
}

if [ "$1" == "register" ]; then
    set_plugin "$@"
    register "$@"
elif [ "$1" == "check" ]; then
    set_plugin "$@"
    check
elif [ "$1" == "help" ]; then
    help
else 
    help
fi