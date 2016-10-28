FROM alpine
MAINTAINER Tom Denham <tom@projectcalico.org>
RUN apk --no-cache add iproute2
ADD libnetwork-plugin /calico

