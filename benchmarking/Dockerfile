FROM ubuntu

RUN apt-get update
RUN apt-get install build-essential libssl-dev git libluajit-5.1-dev luajit vim -y 
RUN cd /home && git clone https://github.com/wg/wrk.git wrk
ADD patchfile /home
RUN ls /home
RUN cd /home && patch wrk/Makefile  patchfile 
RUN cd /home/wrk && make WITH_LUAJIT=/usr WITH_OPENSSL=/usr
RUN cp /home/wrk/wrk /usr/local/bin