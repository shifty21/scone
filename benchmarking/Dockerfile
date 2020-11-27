FROM ubuntu

RUN apt-get update
RUN apt-get install build-essential libssl-dev git libluajit-5.1-dev luajit vim -y 
RUN cd /home && git clone https://github.com/wg/wrk.git wrk
ADD patchfile /home
RUN ls /home
RUN cd /home && patch wrk/Makefile  patchfile 
RUN cd /home/wrk && make WITH_LUAJIT=/usr WITH_OPENSSL=/usr
RUN cp /home/wrk/wrk /usr/local/bin
RUN cd /home && git clone https://github.com/hashicorp/vault-guides.git
COPY benchmark.py /home/vault-guides/operations/benchmarking/wrk-core-vault-operations/
COPY resource.yaml /home/vault-guides/operations/benchmarking/wrk-core-vault-operations/
RUN apt-get install python3 python3-pip -y
RUN pip3 install pyyaml