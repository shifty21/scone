ulimit -n 64000
source /root/.profile
#build vault
#add // +build !gccgo to the files giving reflect.
#change type_map.go file in reflect2 package, which will pe shown by build error
/usr/local/go/bin/go build -compiler gccgo -o vault_binary -v 
#add gccgo compiler flag and remove CGO_ENABLED=0 in consul-template Makefile's dev function and run make dev
# clone scone repo https://github.com/shifty21/scone to build vault initilizer, run make -i in root dir
