
### Install wrk
https://github.com/wg/wrk/wiki/Installing-Wrk-on-Linux

apt-get install build-essential libssl-dev git -y 
git clone https://github.com/wg/wrk.git wrk

cd wrk
make WITH_LUAJIT=/usr WITH_OPENSSL=/usr
# move the executable to somewhere in your PATH, ex:
cp wrk /usr/local/bin


libluajit-5.1-dev   luajit
change make file according to https://github.com/gruebel/wrk/commit/e44b77d5369a6b25b201f14d3e52be9036b8ae87

nohup wrk -t6 -c16 -d20s -H "X-Vault-Token: $VAULT_TOKEN" -s write-random-secrets.lua $VAULT_ADDR -- 10000 > prod-test-write-1000-random-secrets-t6-c16-20sec.log &


# Update write-random-secrets.lua to the path "secret2"
sed -i -e 's+/v1/secret/+/v1/secret2/+g' write-random-secrets.lua
vault secrets enable -path secret2 -version 1 kv
nohup wrk -t1 -c1 -d20s -H "X-Vault-Token: $VAULT_TOKEN" -s write-random-secrets.lua $VAULT_ADDR -- 10000 > prod-test-write-1000-random-secrets-t6-c16-20sec.log &

Luascript
https://github.com/wg/wrk/blob/master/SCRIPTING