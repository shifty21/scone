FROM sconecuratedimages/www2019:mongodb-alpine-scone4.2.1

COPY docker/mongodb/run.sh /home/run.sh
RUN chmod +x /homr/run.sh
ENTRYPOINT ["/home/run.sh"]