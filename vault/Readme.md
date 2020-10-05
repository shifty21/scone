## Vault Codebase

1. Server command 
    * Checks dev flags
        * true - predefined config merges with config provided
        * false - takes configuration from flag
    * Configures grpc logfaker - 
    * memory profile
    * configure telemetry - prometheus, metricssink, and inmemmetrics(metricsutil to initialize inmem and prometheus)
    * initialize physical backend and logger for it - with Specific backend function provided by PhysicalBackends map
    * initialize service discovery if provided - registration of vault as a service with their current state(active, initialized and standby)in consul or k8
    * barrierSeal= auto or shamir and unwrapSeal
    * CoreConfig struct
    * Initialize separate HA storage backend (separate from storage)
    * Vault envirnoment adresses - redirect, advertise and api 
    * Initialize the core
    * configure listeners with address, max req size, duration 
    * Attept unsealing in background. Needed when vault with multiple servers is configured with auto-unseal but uninitialized. Once one server initializes the storage backend, this goroutine will pick up the unseal keys and unseal this instance.
    if !core.IsInSealMigration() block
    * UnsealWithStoredKeys is used to unseal with stored keys, this is where c.raftInfo.joinInProgress is checked get sealtouse(seal configration used(only 1 implementation of seal is present seal_autoseal)) and check till "keyringFound && haveMasterKey" is true
    * if storage type is raft - InitiateRetryJoin
    * run dev if dev flag
    * initialize http handlers
    * wait for shutdown
    * Unsealing is done in core.go's NewCore function
2. OperatorUnsealCommand also calls the http api endpoint 