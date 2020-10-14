## GCCGO Ubuntu Dockerfile 

This Docker file builds gccgo in ubuntu base image along with copying vault and consul-template for cross-compilation. Buidling vault with gccgo gives error in "modern-go/reflect2" package. This is handeled in Dockerfile by editing type_map.go file to add " // +build !gccgo", this make sure gccgo skips this file for building.

Cross-compilation is tested on vault version1.5.3, hence at the end of dockerfile you will find checkout command in vault directory.

