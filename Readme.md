# Vault Initializer
This code base intializes vault instance with shamir based initialization. The intialization is done in 2 ways as of now. 
1. Authorize vault by CAS
2. Authorize vault by CAS along with keyshare encryption. The decryption keys are provided by CAS if the initializer is authenticated by CAS.

For Demo all the files are under resources directory with a total of 4 session that needs to be registered to CAS
1. vault - session_vault.yml vault.crt and valt-key.key
2. vault_initializer - session_vault_init.yml vault-initializer.crt vault=initializer-key.key
3. consul-template session_consul_template.yml consul-template.crt consul-template-key.key
4. demo_client session_client.yml client.crt client-key.key
5. run.sh - registers all the session to local CAS instance
6. resources/config.hcl - vault config file, it requires consul as backend for now. Will try default backend.
6. resources/consul-template/config.hcl - configration file for consul-template containing the template and output file path

## Demo process

docker-compose contains all the required services.

1. Run docker-compose.yaml - this will bring up vault, consul, cas and las instances.
2. login to vault container and run run.sh to register all the binaries.
3. Start vault with - SCONE_CONFIG_ID=vault_demo/vault vault server -config=config.hcl
4. Start vault_initializer - this will intialize vault based on the type suggeted. For the 2nd senario mentioned above the private key for decryption woult be inserted by CAS at /home/mykey.pem. This ensures that only authenticated application is able to intialize vault.
5. Once vault have successfully initialied add key to vault - vault kv put secret/hello hashicorp="101 2nd St"
6. Update demo client's CAS session details to environment variables or configuration file(/resources/consul-template/config.hcl) for consul_template to pick up
7. Run consul_template to render the demo template (/resources/consul-template/find_address.tpl) -> SCONE_CONFIG_ID=consul_template/consul_template consul-template -auth -config config_back.hcl -once
8. The template should be rendered on if demo_client session is verifed by CAS /resources/consul-template/hashicorp_address.txt
