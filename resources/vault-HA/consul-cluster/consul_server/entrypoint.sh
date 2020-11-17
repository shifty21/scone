#!/bin/bash
advertise_addr=`awk 'END{print $1}' /etc/hosts`
printf "environment advertise_addr==[%s]\n" ${advertise_addr}
sed -i "s/ADVERTISE_ADDRESS/${advertise_addr}/g" /usr/local/etc/consul/consul_server.json
envsubst < /usr/local/etc/consul/consul_server.json | tee /usr/local/etc/consul/consul_server.json
cat /usr/local/etc/consul/consul_server.json