#!/bin/bash
SCONE_CONFIG_ID=consul-template/dev SCONE_VERSION=1 /opt/scone/lib/ld-scone-x86_64.so.1 /root/go/bin/consul-template -auth -config /root/go/bin/resources/consul-template/config.hcl -once