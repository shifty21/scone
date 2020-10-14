FROM ubuntu

RUN apt-get update && apt-get -y install curl gnupg2 lsb-release software-properties-common apt-utils && apt-get clean all
RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add -
RUN apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
RUN apt-get update && apt-get -y install consul