FROM alpine
MAINTAINER Tom Denham <tom@projectcalico.org>
ADD libnetwork-plugin /libnetwork-plugin
ENTRYPOINT ["/libnetwork-plugin"]

