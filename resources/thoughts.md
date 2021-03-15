system page size (225) is smaller than minimum page size (4096)
fatal error: bad system page size
runtime: panic before malloc heap initialized

runtime stack:
runtime.dopanic_m
	../../../src_gcc/libgo/go/runtime/panic.go:1211
runtime.fatalthrow
	../../../src_gcc/libgo/go/runtime/panic.go:1071
runtime.throw
	../../../src_gcc/libgo/go/runtime/panic.go:1042
runtime.mallocinit
	../../../src_gcc/libgo/go/runtime/malloc.go:455
runtime.schedinit
	../../../src_gcc/libgo/go/runtime/proc.go:550
main
	/home/buildozer/aports/main/gcc/src/gcc-10.2.0/libgo/runtime/go-main.c:56

	:0
[SCONE|INFO] src/syscall/exit.c:26:syscall_SYS_exit_group(): Persisting fspf data before terminating enclave due to exit call from the application with status 2
[SCONE|DEBUG] src/shielding/hierarchy.c:428:fs_hierarchy_update_persist_fspfs(): Persisting of fspf not necessary: no fspf registered


While building consul on alpine
14 with gcc-go
go1: internal compiler error: in insert, at go/gofrontend/gogo.cc
https://codereview.appspot.com/237960043/diff/40001/go/expressions.cc

ncurses for tput in alpine
go 1.15 - build success  - make tools and make dev

Use CAS to define policy for providing the encryption key to Vault in such a way
that three entities must give access. The objective is that three (or more) entities need to agree to start the secret service. In addition, this task needs to ensure no human even knows part of the keys.
Possible solution - 
1. Multiple cas sessions but the cas session should provide their certificates and session details so vault_initializer can initialize vault.
2. If cas sessions could provide pgp keys then vault could encrypt the init response. Which can further be decrypted by private key from sessions. 

Try these things
1. Generate gpg keys and try using them to auto-unseal, this will encrypt the response of init. These can also be used in normal shamir based initialization. -  done
2. Create one more package in vault-initializer that will initialize auto-unseal case with pgp keys and ask CAS sessions from stakeholders. Is this really possible ? since this would require the key and cert of the sessions. - done with one key which can be provide by CAS - done need to check the CAS session secret import
3. Check the case where vault is initialized and sealed if keys are there in memory unseal else exit. Or try to store in a location? - memory check is inserted with keys in CAS
4. Remove global variable for initresponse - done
5. Decrypt pgp then decode to string - done
6. Combine pgp and other init process and use options to intialize vault interface. Try having one configuration that can have options. 
7. Checked the injection of secret by using secret from other cas sessions - done
   

1. Test case 
2.   Demo - done for dynamic database secret (decrets are pushed by consul-template to the demo-client session)
   1.  vault and vault-init cas registration
   2.  run vault-init grpc and initialize vault - this will export the token to consul-template or any other sevice specified in config
   3.  Add vault commands to be executed for demo in script - done
   4.  register consul-template and demo-client in cas
   5.  Create script for mongodb intialization script - done
   6.  run consul-template and demo-client, demo-client will watch configuration
3.  Setup thesis document
4.  analyze benchmarking result and form graphs
5.  re-run benchmarking with modified script
6.  Add dynamic secret script for benchmarking
7.  Get the related work data
8.  Get vault, consul-template document needed for thesis



### Related work
AWS KMS is probably the closest to what we implemented using Vault, but the cost of KMS Keys and the API rate limits made it a non-starter.

Confidant by Lyft was very interesting, but ultimately the complexity of implementing authentication, the lack of per-customer revocation, and our KMS API rate limits concerns steered us away.

Cerberus by Nike was very interesting, but the stack wasn't something we were comfortable with running in production.


[SCONE|WARN] src/syscall/anon.c:171:mmap_anon(): Protected heap memory exhausted! Set SCONE_HEAP environment variable to increase it.
unable to allocate additional stack space: errno 12
SIGABRT: abort
PC=0x100284456a m=62 sigcode=2

goroutine 1834 [running]:

goroutine 1 [chan receive, 5 minutes]:
main.main
	/root/go/src/github.com/shifty21/scone/demo-client/main.go:125

goroutine 33 [syscall, 5 minutes]:
	goroutine in C code; stack unavailable
created by signal.Notify..func1
	../../../src_gcc/libgo/go/os/signal/signal.go:127 +0x50

goroutine 34 [select, 5 minutes, locked to thread]:

goroutine 35 [sleep]:
time.Sleep
	../../../src_gcc/libgo/go/runtime/time.go:187
main.loadConfig..func1
	/root/go/src/github.com/shifty21/scone/demo-client/main.go:34
created by main.main
	/root/go/src/github.com/shifty21/scone/demo-client/main.go:68 +0x298


### Comparisons
1. Conjur by cyberark


Important section
design
implt
evaluation

Errors 
for Binary secret by CAS 
Caused by: SCONE variable substitution failed - the service configuration needs to be fixed!
Caused by: Environment variable 'MONGO_PASSWORD' contains malformed UTF-8 after SCONE value replacement
for ascii secret value = w(`tF2Q'$xda
error by vault
Failed to parse K=V data: invalid key/value pair "\\": format must be key=value

1. mongodb
   1. Generate tls certificate with CAS and export them to vault-dynamic (wont work, try with vault and export to mongodb)
   2. User this to start mongodb
   3. vault-dynamic registers dynamic secret and then consul-template can watch that config
2. nginx
	1. Configure vault PKI and setup root CA 
	2. Generate Certifcate and key via consul-template
	3. one more module in scone to setup PKI 
	4. consul-template will watch the time
3. consul-template command 
   1. on running command directly to run demo-client "child: fork/exec /root/go/bin/demo-client: function not implemented"
   2. on using rc-service or sh, the binaries are not visible to consul-template
4. export of x509 certificate from CAS doesnt seem to me working
   
5. Separate registration and generation so nginx and mongodb certificate can be generated 
6. generate mongodb certifcate and export to mongodb and vault-init for dynamic secret
7. generate nginx secret to nginx session and start it
8. user mongodb secret to start mongodb and start demo-client
9. Test consul-template if session update is getting reflected without writing to file system

### pki 
https://www.vaultproject.io/api-docs/secret/pki
Api integration needed
1. Enable pki done 
2. Generate Root done
<!-- 3. intermediate
4. https://www.vaultproject.io/api-docs/secret/pki#sign-intermediate -->
5. https://www.vaultproject.io/api-docs/secret/pki#create-update-role done 
6. https://www.vaultproject.io/api-docs/secret/pki#generate-certificate done
<!-- 7. https://www.vaultproject.io/api-docs/secret/pki#set-signed-intermediate -->
   vault write pki/config/urls issuing_certificates="http://127.0.0.1:8200/v1/pki/ca" crl_distribution_points="http://127.0.0.1:8200/v1/pki/crl" done

Vault-init-auto - export vault token to below
- vault-init-pki (setup import for vault token)
- consul-template each service will have
  - export TLS certificate 
    - nginx and start nginx
    - mongodb and start mongodb
    - vault-init-dynamicsecret
- vault-init-dynamicsecret (setup and import tls certificate and vault token)
- consul-template
  

Data benchmarking runs
- basic version
  - time
  - threads
  - connections
  - no of requests/latency(50,90,99,99.99 percentile, mean and stdev)/bytes of data/duration/
- feature version
  - time
  - threads
  - connections
  - no of requests/latency(50,90,99,99.99 percentile, mean and stdev)/bytes of data/duration/

Components involved
Vault - Vault will provide the PKI engine and database engine. It comes with HTTP and command line interface which can be directly used. Both of which require supplying token for communication.

Vault startup provides 2 kinds of un-sealing. First one requires you to supply key-shares via the rest endpoint which gives back set of key-shares that one use to unseal the vault. Second one is auto-unseal which requires trusting a 3rd party KMS for encryption and decryption.

gRPC KMS - For supporting the KMS for encryption and decryption we will create a gRPC based service that will perform the encryption and decryption.

Vault-initializer - To securely distribute the vault token and make sure only authorized application will initialize we will create a Go based vault-initializer that will create a wrapper over the existing vault REST API's. We will support both shamir-key based startup and auto-unseal startup. Once vault is unsealed it results in a root token. For this thesis we will use this token, but in production this should not be used. 

Shamir-key based startup module - Starting vault with shamir-key based technique involves by first calling the intialization api that takes in secret shares and secret thresholds. To enable encryption of keys vault provides option to set PGP keys and these keys are to be provided by the stakeholders. While starting up vault with SCONE backed by CAS theses key would be provided by multple CAS sessions via export secret option. Each key share is encrypted with the public key provided in the order of supplied keys. The output is an encrypted key share response. Each of these encrypted keys will be decrypted by the Private key counter-part of the public key which is provided by the CAS session.
{
  "secret_shares": 10,
  "secret_threshold": 5
}
Once vault is initialized it requires to be unsealed which is done by supplying the key-shares one by one till the threshold is reached. Upon reaching the threshold vault is unsealed and outputs a root token.

Auto-Unseal based startup module - Unlike Shamir based Unseal mechanism this one requires external KMS engines like AliCloud KMS, Amazon KMS, Azure Key Vault, and Google Cloud KMS. The external KMS engines provide API endpoints to encrypt and decrypt vault keys. This can be easily provided by scone by providing these two endpoints to encrypt and decrypt data. Startup with auto-unseal only requires us to call one endpoint which will output recovery key shares and a root token. The recovery key shares can be used to perform admin operation in case of loss of root token.

To support this create a encryption service that provides these two endpoints. This can be done via REST endpoints or gRPC service. Since gRPC is the defacto standard for inter-service interaction we will explore that option.

CAS - CAS will be used to store sessions and updated sessions of all the applications. A secure way to supply the tokens to these application can be done via CAS sessions.

CAS module - CAS module will perform 2 roles. First to support the supply of tokens other service session. This CAS module will interact with CAS to get the latest CAS session of a service and update it by adding a CAS secret called VAULT_TOKEN and exporting it to services that need it. In the thesis 3 services will need the token to perform their opeartions including pki_engine module, database_secret engine module and consul_template.

Second role of CAS module will be to inject the generated secrets from vault into CAS session of the application. This will be done in consul_template.

go-kms-wrapping - Auto-Unseal in vault requires an external KMS service for encryption and decryption of master key. Vault calls the KMS services via "go-kms-wrapping"\cite{go-kms-wrapping} library which acts as wrapper for calling the KMS stores. We will need to include our gRPC services client endpoints in "go-kms-wrapper". There will also be a small modification required on vault to include the configuration and calling of gRPC service.

gRPC-KMS - We will create a gRPC service which will expose 2 interfaces over TLS to support encryption and decryption needs of vault.
consul_template - The secrets supported by consul_template are rendered on disk. We will replace this by CAS module which will render the secrets as injection files in CAS sessions. It requires one to supply "cas" stanza in consul-template config.

cas {
  get_session_api = "https://cas:8081/session"
  cert = "/root/go/bin/resources/vault-init/conf/client.crt"
  key = "/root/go/bin/resources/vault-init/conf/client-key.key"
  session_name = "vault-init-dynamicsecret"
  session_file = "/root/go/bin/resources/vault-init/session_dynamic.yml"
  predecessor_hash_file = "/root/go/bin/resources/vault-init/dynamicsecret_predecessor_hash.yaml"
}

In a production environment consul-template is deployed with the application that requires the secrets.


We will cover 2 Application Scenarios in the thesis. First, using the PKI secret engine and second Database secret engine. Both engines are described in detail in background section. We have taken Nginx as a sample application to enable scone support for vault's PKI engine and MongoDB for the database secret engine. Both of these are used with their respective sconcurated images.

Steps involved in these scenarios.
1. Initialize and unseal vault. These steps are performed by user and the root token provided by Vault is used by user to configure PKI and database secret engine along with consul_template.
   1. Configure the PKI secret engine with the token from step 1
   2. Configure the Database secret engine with the token from step 1
2. Supply consul-template the token received in step 1. Configure consul-template to read certificates from vault for Nginx and MongoDB and database credentials for MongoDB.
3. Secret rendering by consul-template
   1. Render nginx certificates and start nginx. 
   2. Render MongoDB certificates and start MongoDB
   3. Render database credentials for MongoDB and start the demo-client to test connection
4. Consul-template keep monitoring for TTL of the secrets for any change and re-renders the secrets on expiry.

Using vault with SCONE will involve following steps.
1. Initialize and unseal vault with vault_initializer. The root token provided by Vault after unsealing is exported to the CAS sessions of service that will require this for operation. This will involve pki_module, database secret enigne module and consul-template.
   1. Configure the PKI secret engine with the token import from CAS secret by vault_intializer.
   2. Configure the Database secret engine with the token import from CAS secret by vault_intializer.
2. Use the imported token from vault_initializer to startup up consul-template. Configure consul-template to read certificates from vault for Nginx and MongoDB and database credentials for MongoDB.
3. Secret rendering by consul-template will be done in CAS sessions of respective service. Each service will have its own consul-template session.
   1. Render nginx certificates in the nignx CAS session.
   2. Render MongoDB certificates in MongoDB's CAS session.
   3. Render database credentials for MongoDB in demo-client's CAS session and test test connection
4. Consul-template keep monitoring for TTL of the secrets for any change and updates secrets into CAS sessions upon expiry.


The docker-Compose services will involve 7 services
1. CAS - CAS is the container for Configuration and Attestation Service
2. LAS - LAS is the container required for local attestation and creation of quotes
3. MongoDB - Container to run SCONE based MongoDB
4. Demo-client - Container to run consul-template and demo-client that will test connection to MongoDB
5. Nginx - Container to run consul-template and SCONE based Nginx with secrets generated by Vault.
6. Benchmarking - Container to perform benchmarking
7. Vault - Container for SCONE based Vault and vault_intializer
8. Consul - Container for Vault's backend DB




1 consul - complete with 1 extra list run
1 mem - without list run 



need 
1 list run of mem - 2
1 complete run of mem - 8
1 consul run without list - 6
1 run contins - 4 native 4 scone - 8 days

Vault start
SCONE - 36.75
Native - 0.56


vault start
scone - 67.56sec, 37.41sec, 58.22sec, 37.23sec, 38.20sec = 52.278sec
native - 1.16sec,2.47sec, 0.93sec, 1.90sec, 0.94sec = 1.48sec
Init
SCONE - 56.42s, 18.72s, 48.38s, 48.22s, 20.23s = 38.394 sec
Native - 0.95*2, 0.93*3 = 0.938sec

PKI setup
native - 1.56sec
scone - 28.05s, 24.68s = 26.36sec

Dynamic - 43.84s
native - 20.13s



nginx

native - 0.51sec
SCONE - 18.82sec

Consul-template for certificate generation
native - 17.27s
SCONE - 29.33s

mongodb
native - 5.26
scone - 7.95sec

credentials by consul-template
native - 0.52s
SCONE - 16.89s

demo -
native - 0.85s
scone - 52.34sec  


Error while running command from consul-template via SCONE runtime
2021/03/12 23:58:12.370950 [INFO] (child) spawning: /bin/sh /root/go/bin/run.sh
[SCONE|INFO] src/enclave/enclave.c:389:_enclave_terminate(): Persisting fspf data before terminating enclave due to unhandled signal 6
[SCONE|DEBUG] src/shielding/hierarchy.c:432:fs_hierarchy_update_persist_fspfs(): Persisting of fspf not necessary: no fspf registered
[SCONE|ERROR] tools/starter/signal.c:55:die_by_signal(): Enclave terminated due to signal: Aborted