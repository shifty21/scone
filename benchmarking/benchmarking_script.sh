
wrk -t6 -c16 -d20s -H "X-Vault-Token: $VAULT_TOKEN" -s write-random-secrets.lua $VAULT_ADDR -- 10000 > prod-test-write-1000-random-secrets-t6-c16-20sec.log 