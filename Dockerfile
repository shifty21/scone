FROM sconecuratedimages/www2019:vault-1.5.0-alpine

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk updates
RUN apk add make vim git curl git musl-utils busybox-extras openssh-client
RUN apk add libc-dev gcc-go

RUN rm /root/go/bin/*

#packages
#vault_initializer
ADD resources /root/go/bin/resources
ADD . / /root/go/src/github.com/shifty21/scone/
RUN cd /root/go/src/github.com/shifty21/scone/ && go build -compiler gccgo -o /root/go/bin/vault_init -v

ADD gitconfig/id_rsa /root/.ssh/scone
ADD gitconfig/config /root/.ssh/config
RUN chmod 0600 /root/.ssh/scone
RUN chmod 0600 /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN cat ~/.ssh/config
RUN eval "$(ssh-agent -s)" && ssh-add /root/.ssh/scone
#vault and consul-template
RUN go get github.com/mitchellh/gox
RUN mkdir /root/go/src/github.com/hashicorp
RUN cd /root/go/src/github.com/shifty21 && git@github.com:shifty21/go-kms-wrapping.git
RUN cd /root/go/src/github.com/shifty21 && git clone git@github.com:shifty21/consul-template.git
RUN cd /root/go/src/github.com/shifty21/consul-template && go mod tidy
RUN cd /root/go/src/github.com/shifty21/consul-template && go build -compiler gccgo -o /root/go/bin/consul-template -v

RUN cd /root/go/src/github.com/hashicorp && git clone git@github.com:shifty21/vault.git
RUN cd /root/go/src/github.com/hashicorp/vault && git checkout tags/v1.5.3 -b dev
RUN cd /root/go/src/github.com/hashicorp/vault && go mod tidy
RUN cd /root/go/pkg/mod/github.com/modern-go/reflect2@v1.0.1 && printf '// +build !gccgo \n \n \n' | cat - type_map.go > /tmp/out && mv /tmp/out type_map.go
RUN cd /root/go/src/github.com/hashicorp/vault && go build -compiler gccgo -o /root/go/bin/vault -v


RUN export PATH=$PATH:/root/go/bin
