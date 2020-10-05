
## Vault Dockerfile
This Docker file  sets up environment for vault for cross-compilation. Buidling vault with gccgo gives error in "modern-go/reflect2" package. Just add a comment in the type_map.go file. " // +build !gccgo" which will skip this particular file for compilation by gccgo.