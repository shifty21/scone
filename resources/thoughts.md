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
1. Generate gpg keys and try using them to auto-unseal, this will encrypt the response of init. These can also be used in normal shamir based initialization.
2. Create one more package in vault-initializer that will initialize auto-unseal case with pgp keys and ask CAS sessions from stakeholders. Is this really possible ? since this would require the key and cert of the sessions.
3. Check the case where vault is initialized and sealed if keys are there in memory unseal else exit. Or try to store in a location?
4. Remove global variable for initresponse
5. Decrypt pgp then decode to string