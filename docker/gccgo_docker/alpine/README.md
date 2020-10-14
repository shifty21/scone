## GCCGO Alpine Dockerfile

The Dockerfile tries to build gccgo from source but few errors were encountered during the process.

1. For devel/gccgo of gccgo branch
```
sigtab.go:38:2: error: duplicate value for index 29
   38 |  _SIGPOLL: {_SigNotify, "SIGPOLL: pollable event occurred"},
      |  ^
../../../libgo/go/runtime/signal_unix.go:786:3: error: range clause must have array, slice, string, map, or channel type
  786 |   for i := range sigtable {
      |   ^
../../../libgo/go/runtime/signal_unix.go:1016:2: error: range clause must have array, slice, string, map, or channel type
 1016 |  for i := range sigtable {
      |  ^
../../../libgo/go/runtime/signal_unix.go:786:7: error: invalid type for range clause
  786 |   for i := range sigtable {
      |       ^
../../../libgo/go/runtime/signal_unix.go:1016:6: error: invalid type for range clause
 1016 |  for i := range sigtable {
      |      ^
libtool: compile:  /gccgo/objdir/./gcc/gccgo -B/gccgo/objdir/./gcc/ -B/opt/gccgo/x86_64-pc-linux-musl/bin/ -B/opt/gccgo/x86_64-pc-linux-musl/lib/ -isystem /opt/gccgo/x86_64-pc-linux-musl/include -isystem /opt/gccgo/x86_64-pc-linux-musl/sys-include -minline-all-stringops -O2 -g -I . -c -fgo-pkgpath=image/color ../../../libgo/go/image/color/color.go ../../../libgo/go/image/color/ycbcr.go -o image/color.o >/dev/null 2>&1
f="image/color.o"; if test ! -f $f; then f="image/.libs/color.o"; fi; objcopy -j .go_export $f image/color.s-gox.tmp; /bin/sh ../../../libgo/mvifdiff.sh image/color.s-gox.tmp `echo image/color.s-gox | sed -e 's/s-gox/gox/'`
echo timestamp > image/color.s-gox
make[4]: *** [Makefile:2869: runtime.lo] Error 1

```
This links talks about the same problem https://github.com/richfelker/musl-cross-make/issues/44. I posted a query and someone pointed out to the new gcc-go in alpine https://pkgs.alpinelinux.org/package/edge/main/x86_64/gcc-go. 

https://groups.google.com/g/gofrontend-dev/c/-DfK2VkZ9zE talks about the patch for the above problem.



On commenting the duplicate mentioned in above error, you will get the below error. This is a pretty old bug from 2012 https://gcc.gnu.org/bugzilla/show_bug.cgi?id=52218. This was supposed to be fixed in https://patchwork.ozlabs.org/project/gcc/patch/yddzkcq240c.fsf@manam.CeBiTec.Uni-Bielefeld.DE/

``` 
../../../libgo/runtime/proc.c:172:4: error: #error unknown case for SETCONTEXT_CLOBBERS_TLS
  172 | #  error unknown case for SETCONTEXT_CLOBBERS_TLS
      |    ^~~~~
depbase=`echo runtime/yield.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CC   --mode=compile /gccgo/objdir/./gcc/xgcc -B/gccgo/objdir/./gcc/ -B/opt/gccgo/x86_64-pc-linux-musl/bin/ -B/opt/gccgo/x86_64-pc-linux-musl/lib/ -isystem /opt/gccgo/x86_64-pc-linux-musl/include -isystem /opt/gccgo/x86_64-pc-linux-musl/sys-include    -DHAVE_CONFIG_H -I. -I../../../libgo  -I ../../../libgo/runtime -I../../../libgo/../libffi/include -I../libffi/include -pthread -L../libatomic/.libs  -fexceptions -fnon-call-exceptions -fsplit-stack -Wall -Wextra -Wwrite-strings -Wcast-qual  -minline-all-stringops  -D_GNU_SOURCE -D_LARGEFILE_SOURCE -D_FILE_OFFSET_BITS=64 -I ../../../libgo/../libgcc -I ../../../libgo/../libbacktrace -I ../../gcc/include -g -O2 -MT runtime/yield.lo -MD -MP -MF $depbase.Tpo -c -o runtime/yield.lo ../../../libgo/runtime/yield.c &&\
mv -f $depbase.Tpo $depbase.Plo
../../../libgo/runtime/proc.c: In function ‘runtime_gogo’:
../../../libgo/runtime/proc.c:290:2: warning: implicit declaration of function ‘fixcontext’; did you mean ‘setcontext’? [-Wimplicit-function-declaration]
  290 |  fixcontext(ucontext_arg(&newg->context[0]));
      |  ^~~~~~~~~~
      |  setcontext
../../../libgo/runtime/proc.c: In function ‘runtime_mstart’:
../../../libgo/runtime/proc.c:555:2: warning: implicit declaration of function ‘initcontext’; did you mean ‘setcontext’? [-Wimplicit-function-declaration]
  555 |  initcontext();
      |  ^~~~~~~~~~~
```


2. For gcc10 branch this is the error which will come when we do make. Alpine doesnt seems to have fstab. More searching needs to be done for this.


```
../../../../libsanitizer/sanitizer_common/sanitizer_platform_limits_posix.cpp:61:10: fatal error: fstab.h: No such file or directory
   61 | #include <fstab.h>
      |          ^~~~~~~~~
compilation terminated.
```
