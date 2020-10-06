#!/bin/bash
SESSION_LIST="vault vault-init consul-template demo-client-session"
 
set_plugin() {
    if [ "$2" == "vault" ]; then
        RUN_SESSION="vault"
    elif [ "$2" == "vault-init" ]; then
        RUN_SESSION="vault-ini"
    elif [ "$2" == "consul-template" ]; then
        RUN_SESSION="consul-template"
    elif [ "$2" == "demo-client-session" ]; then
        RUN_SESSION="demo-client-session"
    else 
        RUN_SESSION=${SESSION_LIST}
    fi
}

register(){
    for plugin in ${RUN_SESSION}; do
        printf '\nRegistring CAS session for %s\n=============================================================\n' "${plugin}"
        cd /resources/"${plugin}" && curl -k -s --cert conf/client.crt --key conf/client-key.key --data-binary @session.yml -X POST https://"$SCONE_CAS_ADDR":8081/session
        printf "\n============================================================= \n"
    done
}

check(){

    for plugin in ${RUN_SESSION}; do
        printf '\nChecking CAS session for %s\n=============================================================\n' "${plugin}"
        cd /resources/"${plugin}" && curl -k -s --cert conf/client.crt --key conf/client-key.key https://"$SCONE_CAS_ADDR":8081/session/"${plugin}"
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
    register
elif [ "$1" == "check" ]; then
    set_plugin "$@"
    check
elif [ "$1" == "help" ]; then
    help
else 
    help
fi
