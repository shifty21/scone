FROM registry.scontain.com:5050/sconecuratedimages/apps:mongodb-alpine-scone4.2.1

RUN apk add vim curl
RUN mkdir -p /usr/local/var/mongodb
COPY docker/mongodb/run.sh /home/run.sh
RUN chmod +x /home/run.sh
ADD resources/mongodb /home/mongodb
ENTRYPOINT [ "/bin/sh" ]