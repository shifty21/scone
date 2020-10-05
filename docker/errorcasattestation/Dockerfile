FROM sconecuratedimages/www2019:vault-0.10.0-alpine

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update
RUN apk add make vim git curl git musl-utils busybox-extras openssh-client
RUN apk add libc-dev gcc-go
#packages
#vault_initializer
ADD resources resources
RUN go get github.com/mitchellh/gox
RUN mkdir /root/go/src/github.com/hashicorp
RUN cd /root/go/src/github.com/hashicorp && git clone https://github.com/hashicorp/consul-template.git
RUN cd /root/go/src/github.com/hashicorp/consul-template && go mod tidy
RUN cd /root/go/src/github.com/hashicorp/consul-template && go build -compiler gccgo -o /root/go/bin/consul-template -v
RUN export PATH=$PATH:/root/go/bin
