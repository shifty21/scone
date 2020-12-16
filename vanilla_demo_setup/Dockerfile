FROM sconecuratedimages/www2019:vault-1.5.0-alpine

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
# RUN cat /etc/apk/repositories
RUN apk update
RUN apk add make vim git curl
RUN apk add gcc-go

RUN rm /root/go/bin/*

#packages
#vault_initializer
ADD resources /root/go/bin/resources

ADD gitconfig/id_rsa /root/.ssh/id_rsa
ADD gitconfig/config /root/.ssh/config
RUN chmod 0600 /root/.ssh/id_rsa
RUN chmod 0600 /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN cat ~/.ssh/config
RUN eval "$(ssh-agent -s)" && ssh-add /root/.ssh/id_rsa
ADD . / /root/go/src/github.com/shifty21/scone/

#vault and consul-template
RUN go get github.com/mitchellh/gox
RUN cd /root/go/src/github.com/shifty21/scone/ && go build  -o /root/go/bin/vault-init -v
RUN cd /root/go/src/github.com/shifty21 && GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone git@github.com:shifty21/go-kms-wrapping.git
RUN cd /root/go/src/github.com/shifty21 && GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone git@github.com:shifty21/vault.git
# RUN cd /root/go/src/github.com/shifty21 && GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone git@github.com:shifty21/consul-template.git
# RUN cd /root/go/src/github.com/shifty21/consul-template && go build -o /root/go/bin/consul-template -v
RUN cd /root/go/src/github.com/shifty21/vault && go build -o /root/go/bin/vault -v
RUN export PATH=$PATH:/root/go/bin
