FROM registry.scontain.com:5050/sconecuratedimages/www2019:vault-1.5.0-alpine

RUN echo "http://dl-cdn.alpinelinux.org/alpine/v3.12/main" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/v3.12/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update
RUN apk add make vim git curl git musl-utils busybox-extras openssh-client
RUN apk add libc-dev gcc-go
#for service management
RUN apk add openrc
ADD resources /root/go/bin/resources

ADD gitconfig/id_rsa /root/.ssh/id_rsa
ADD gitconfig/config /root/.ssh/config
RUN chmod 0600 /root/.ssh/id_rsa
RUN chmod 0600 /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN cat ~/.ssh/config
RUN eval "$(ssh-agent -s)" && ssh-add /root/.ssh/id_rsa

ADD . / /root/go/src/github.com/shifty21/scone/
RUN cd /root/go/src/github.com/shifty21/scone/demo-client && go build -compiler gccgo -o /root/go/bin/demo-client -v
RUN go get github.com/mitchellh/gox
RUN cd /root/go/src/github.com/shifty21 && GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone git@github.com:shifty21/consul-template.git
RUN cd /root/go/src/github.com/shifty21/consul-template && go build -compiler gccgo -o /root/go/bin/consul-template -v
RUN cd /root/go/src/github.com/shifty21/scone/ && go build -compiler gccgo -o /root/go/bin/vault-init -v
ENV PATH="${PATH}:/root/go/bin"

ADD demo-client /etc/init.d/
RUN chmod +x /etc/init.d/demo-client
RUN rc-update add demo-client
