FROM sconecuratedimages/www2019:vault-1.5.0-alpine

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update
RUN apk add make vim git curl git musl-utils busybox-extras openssh-client
RUN apk add libc-dev gcc-go
#packages
#vault
RUN go get github.com/mitchellh/gox
RUN mkdir /root/go/src/github.com/hashicorp
RUN cd /root/go/src/github.com/hashicorp && git clone https://github.com/hashicorp/vault.git
RUN cd /root/go/src/github.com/hashicorp/vault && git checkout tags/v1.5.3 -b dev
RUN cd /root/go/src/github.com/hashicorp/vault && go mod tidy
RUN cd /root/go/pkg/mod/github.com/modern-go/reflect2@v1.0.1 && printf '// +build !gccgo \n \n \n' | cat - type_map.go > /tmp/out && mv /tmp/out type_map.go
RUN cd /root/go/src/github.com/hashicorp/vault && go build -compiler gccgo -o /root/go/bin/vault -v
RUN export PATH=$PATH:/root/go/bin
ADD resources /resources
