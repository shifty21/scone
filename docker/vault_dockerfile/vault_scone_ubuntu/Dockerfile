#tested with 20.04 adjust make cores according to system
FROM sconecuratedimages/research:scone-run-ubuntu-18.04-scone4.2.1-v0.3
RUN apt-get update
ENV DEBIAN_FRONTEND="noninteractive"
ENV TZ=Europe/Berlin
#gccgo related libraries
RUN apt-get install -qq -y apt-utils gcc git subversion make g++ flex libgmp-dev libmpfr-dev libmpc-dev curl
RUN apt-get install -qq -y dialog apt-utils wget vim
RUN apt-get install -qq -y bison binutils-dev libipt-dev libtool texinfo 
#git hash e109f6e438b72ef3e403162971068d28d09b82f5 gccgo (GCC) 11.0.0
RUN git clone --progress --branch devel/gccgo git://gcc.gnu.org/git/gcc.git gccgo
WORKDIR /gccgo/
RUN ./contrib/download_prerequisites
#build gccgo
RUN mkdir objdir
RUN cd objdir && ../configure --prefix=/opt/gccgo --enable-languages=c,c++,go \
    --disable-libquadmath \
    --disable-libquadmath-support \
    --disable-werror \
    --disable-multilib
RUN cd objdir && make -j4
RUN echo "gccgo make status $?"
RUN cd objdir && make install -j4
RUN echo "gccgo make install status $?"
#clone consul-template and vault and install golang 1.15
WORKDIR /root/
RUN mkdir -p /root/go/src/github.com/hashicorp
RUN cd /root/go/src/github.com/hashicorp && git clone https://github.com/hashicorp/vault.git
RUN wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
RUN echo export PATH=$PATH:/usr/local/go/bin:/opt/gccgo/bin >> /root/.profile

#gccgo lib location  add to LD as well
RUN echo export LIBDIR=/opt/gccgo/lib/../lib64 >> /root/.profile
RUN echo export LD_LIBRARY_PATH=/usr/local/lib64:/opt/gccgo/lib/../lib64 >> /root/.profile
RUN cd ~/go/src/github.com/hashicorp/vault && git checkout tags/v1.5.3 -b dev
RUN cd /root/go/src/github.com/hashicorp/vault && /usr/local/go/bin/go mod tidy
RUN cd /root/go/pkg/mod/github.com/modern-go/reflect2@v1.0.1 && printf '// +build !gccgo \n \n \n' | cat - type_map.go > /tmp/out && mv /tmp/out type_map.go
#RUN cd /root/go/src/github.com/hashicorp/vault && /usr/local/go/bin/go build -compiler gccgo -o /root/go/bin/vault -v
#copy resource
ADD resources /resources