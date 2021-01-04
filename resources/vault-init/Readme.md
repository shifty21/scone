### Scone's Vault-Init Codebase

There are multiple service provided by this binary including
1. gPRC service that acts as secure initalization service
2. Multiple initialization types
   1. Vanilla that initialize w/o gpg or any other security(this is purely for testing purpose)
   2. Shamir based initialiation that uses gpg keys provided by other CAS sessions
   3. Auto-initialization which does auto-initialization using the gRPC serivce provided by scone
   Both vault-init and Shamir intialization can export token to external sessions
3. Register service that can be used to register cas session with provided yaml
4. Benchmarking automation python module - that leverages wrk and vault provided benchmaring script to autmate the benchmark genration
### Scone notes
1. Image
   1. For Injection of file image is used in session description. It can have injection_file or volume
      1. volume - system region that can use fspf, and be shared by exporting 
      2. injection_files - These are updated by CAS session if they have $$SCONE$$ in them, if the content is not provided
        contents of the file at the specified path are stored in-memory which are serverd by scone