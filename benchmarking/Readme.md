
### Install wrk
https://github.com/wg/wrk/wiki/Installing-Wrk-on-Linux

apt-get install build-essential libssl-dev git -y 
git clone https://github.com/wg/wrk.git wrk

cd wrk
make WITH_LUAJIT=/usr WITH_OPENSSL=/usr
# move the executable to somewhere in your PATH, ex:
cp wrk /usr/local/bin


libluajit-5.1-dev luajit
change make file according to https://github.com/gruebel/wrk/commit/e44b77d5369a6b25b201f14d3e52be9036b8ae87

nohup wrk -t6 -c16 -d20s -H "X-Vault-Token: $VAULT_TOKEN" -s write-random-secrets.lua $VAULT_ADDR -- 10000 > prod-test-write-1000-random-secrets-t6-c16-20sec.log &


# Update write-random-secrets.lua to the path "secret2"
sed -i -e 's+/v1/secret/+/v1/secret2/+g' write-random-secrets.lua
e84cf16c5d71257dd11bf1e0a2fced08311a5a9e63d3d625b65e44dccd30014a55
2020/12/11 21:03:15 UnsealPerKey|Unsealing for key c19f00eda5ca577c58d6d7e71e9cbf8b9d2b8c2079303c22f43c42df3866259767
2020/12/11 21:03:15 UnsealPerKey|Unsealing for key af1e831d82f9f6cbe12b649ff9b638ec574be9bf4d58a1424361e58bd68f1659f5

export VAULT_ADDR=http://vault1:8200
export VAULT_TOKEN=s.0Gu7cfsRTCTlJqtjjb129wB8
./vault login s.0Gu7cfsRTCTlJqtjjb129wB8
./vault secrets enable -path secret2 -version 1 kv
./vault auth enable userpass
./vault write auth/userpass/users/loadtester password=benchmark policies=default

nohup wrk -t1 -c1 -d20s -H "X-Vault-Token: $VAULT_TOKEN" -s write-random-secrets.lua $VAULT_ADDR -- 10000 > prod-test-write-1000-random-secrets-t6-c16-20sec.log &

Luascript
https://github.com/wg/wrk/blob/master/SCRIPTING

curl --header "X-Vault-Token: $VAULT_TOKEN" --request POST --data @payload.json http://127.0.0.1:8200/v1/secret/hello

curl --header "X-Vault-Token: $VAULT_TOKEN" http://vault:8200/v1/secret/hello


 &{[ec9f2fb75d63e78e2106c30ebe4e8c6dc9d45aaecd2a92f6f568bd218a151a8839 d9e87c734b9548d38d1d3a117d6505352e575e7de191f574120c1be67ed6bc4f99 899b7322d4196a1a192081a1b493b639fe17e50810961a88db3d754a25f2673220 d0a7f409fcc85ad5657f3f765a3cd6037c2e46d58fb806105c1ad11757c5fa4243 3b66d087c4085d3cf629be8fda1ce982c54dffc4ed5f7484041fee531694c4aee5] [7J8vt11j544hBsMOvk6MbcnUWq7NKpL29Wi9IYoVGog5 2eh8c0uVSNONHToRfWUFNS5XXn3hkfV0Egwb5n7WvE+Z iZtzItQZahoZIIGhtJO2Of4X5QgQlhqI2z11SiXyZzIg 0Kf0CfzIWtVlfz92WjzWA3wuRtWPuAYQXBrRF1fF+kJD O2bQh8QIXTz2Kb6P2hzpgsVN/8TtX3SEBB/uUxaUxK7l] [] [] s.8G00yeiBG8x4FOXdInvvDueR}

 Unseal|Auto-Initialization Enabled, no need to call unseal apis. InitResponse: [&{[] [] [] [34acec262f9b27b36b336a726c9d37627708ac18a50eb2b9dff7a7d7261d5fb8] s.fKVhmaDVj4FZJCS4gQFPfG5g}]