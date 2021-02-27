FROM registry.scontain.com:5050/sconecuratedimages/apps:mongodb-alpine-scone4.2.1
RUN apk add git
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update
RUN apk  --force-overwrite  add make vim musl-utils busybox-extras openssh-client
RUN apk add libc-dev gcc-go

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
RUN cd /root/go/src/github.com/shifty21 && GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone git@github.com:shifty21/consul-template.git
RUN cd /root/go/src/github.com/shifty21/consul-template && go build -compiler gccgo -o /root/go/bin/consul-template -v
RUN cd /root/go/src/github.com/shifty21/scone/ && go build -compiler gccgo -o /root/go/bin/vault-init -v
ENV PATH="${PATH}:/root/go/bin"

RUN mkdir -p /usr/local/var/mongodb
COPY docker/mongodb/run.sh /home/run.sh
RUN chmod +x /home/run.sh
ADD resources/mongodb /home/mongodb
ENTRYPOINT [ "/bin/sh" ]