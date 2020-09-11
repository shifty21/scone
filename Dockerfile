FROM sconecuratedimages/www2019:vault-0.10.0-alpine

ADD resources resources

ADD . / /root/go/src/github.com/shifty21/scone/

RUN apk add make vim git curl go musl-utils
RUN chown root /usr/sbin/vault
RUN mv /usr/local/bin/scone-gccgo /usr/local/bin/gccgo