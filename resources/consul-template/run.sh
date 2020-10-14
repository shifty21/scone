#!/bin/bash
SCONE_CONFIG_ID=consul-template/dev /root/go/bin/consul-template -auth -config config_back.hcl -once