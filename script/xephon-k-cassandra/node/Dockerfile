#FROM alpine:3.5
FROM ubuntu:16.04

MAINTAINER at15 at15@dongyue.io

# Thanks to @zbintliff in https://github.com/kairosdb/kairosdb/issues/288
#RUN apk upgrade libssl1.0 --update-cache && \
#    apk add curl ca-certificates bash

# KairosDB must start after Cassandra is running and thrift protocol is enabled
COPY wait-for-it.sh /usr/bin/wait-for-it
COPY xkdc.sh /usr/bin/xkdc
COPY xkd /usr/bin/xkd

EXPOSE 23333
ENTRYPOINT []
CMD [ "/usr/bin/xkdc" ]
