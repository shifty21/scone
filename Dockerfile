FROM sconecuratedimages/www2019:vault-0.10.0-alpine

ADD resources resources

ADD . / /root/go/src/github.com/shifty21/scone/

RUN apk add make vim git curl git musl-utils busybox-extras
RUN chown root /usr/sbin/vault
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

RUN mkdir /root/go/src/github.com/hashicorp
RUN mv consul-template/ /root/go/src/github.com/hashicorp/