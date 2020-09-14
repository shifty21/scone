FROM sconecuratedimages/www2019:vault-0.10.0-alpine


RUN echo -n "" > /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN apk update
ADD resources resources
ADD . / /root/go/src/github.com/shifty21/scone/

RUN apk add make vim git curl git musl-utils busybox-extras go openssh-client
RUN chown root /usr/sbin/vault
#RUN rm -rf /usr/sbin/vault
#RUN hash -r
RUN mv /usr/local/bin/scone-gccgo /usr/local/bin/gccgo

#consul-template
ADD gitconfig/scone /root/.ssh/scone
ADD gitconfig/config /root/.ssh/config
RUN chmod 0600 /root/.ssh/scone
RUN chmod 0600 /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN cat ~/.ssh/config
RUN eval "$(ssh-agent -s)" && ssh-add /root/.ssh/scone

RUN git clone git@github.com:shifty21/consul-template.git
RUN git clone https://github.com/hashicorp/vault.git

RUN mkdir /root/go/src/github.com/hashicorp
RUN mv consul-template/ /root/go/src/github.com/hashicorp/
RUN mv vault/ /root/go/src/github.com/hashicorp/
RUN cd /root/go/src/github.com/hashicorp/consul-template && go mod tidy
RUN cd /root/go/src/github.com/hashicorp/vault && go mod tidy

RUN go get github.com/mitchellh/gox
RUN export PATH=$PATH:/root/go/bin

