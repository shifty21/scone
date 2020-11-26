#Setup initialization of services for a demo instance
#1. vault with config without auto mode 2. vault-init shamir 3. Start the benchmarking 
/root/go/bin/vault server -config /root/go/bin/resources/vault/config_vanilla.hcl

# add flag to disable gpg
/root/go/bin/vault-init scone 

#add option to save token