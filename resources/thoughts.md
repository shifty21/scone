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

