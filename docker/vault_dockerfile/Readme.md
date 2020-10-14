
## Vault Dockerfile
This Docker file  sets up environment for vault for cross-compilation. Buidling vault with gccgo gives error in "modern-go/reflect2" package. This is handeled in Dockerfile by editing type_map.go file to add " // +build !gccgo", this make sure gccgo skips this file for building.

[Follow](../../resources/Readme.md) for registring and running vault