## GCCGO Ubuntu Dockerfile 
This Docker file builds gccgo in ubuntu base image along with copying vault and consul-template for cross-compilation. Buidling vault with gccgo gives error in "modern-go/reflect2" package. Just add a comment in the type_map.go file. " // +build !gccgo" which will skip this particular file for compilation by gccgo.

cross-compilation is tested on vault version1.5.3, hence at the end of dockerfile you will find checkout command in vault directory.

