#adjust make cores according to system
FROM alpine

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/main > /etc/apk/repositories
RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories
RUN apk update
#gccgo related libraries
RUN apk add gcc git subversion make g++ flex curl dialog wget vim
# libipt-dev apk not present  libgmp-dev libmpfr-dev libmpc-dev
RUN apk add bison binutils-dev libtool texinfo curl-dev
#git hash e109f6e438b72ef3e403162971068d28d09b82f5 gccgo (GCC) 11.0.0
RUN git clone --progress --branch devel/gccgo git://gcc.gnu.org/git/gcc.git gccgo
#ADD gccgo /gccgo/
WORKDIR /gccgo/
RUN ./contrib/download_prerequisites

#build gccgo
RUN apk add mpc1-dev gmp-dev mpfr-dev gcc-gnat gnats libc-dev linux-headers libc-utils
RUN mkdir objdir
RUN cd objdir && ../configure --prefix=/opt/gccgo --enable-languages=go \
    --disable-libquadmath \
    --disable-libquadmath-support \
    --disable-werror \
    --disable-multilib
RUN cd objdir && make -j2
RUN echo "gccgo make status $?"
RUN cd objdir && make install -j2
RUN echo "gccgo make install status $?"
#RUN echo export LIBDIR=/opt/gccgo/lib/../lib64 >> /root/.profile
#RUN echo export LD_LIBRARY_PATH=/usr/local/lib64:/opt/gccgo/lib/../lib64 >> /root/.profile

