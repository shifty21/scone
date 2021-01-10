FROM sconecuratedimages/www2019:mongodb-alpine-scone4.2.1

RUN apk add vim curl
COPY docker/mongodb/run.sh /home/run.sh
RUN chmod +x /home/run.sh
ADD resources/mongodb /home/mongodb
# ENTRYPOINT ["/home/run.sh"]