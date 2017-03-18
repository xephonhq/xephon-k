FROM java:8-alpine

MAINTAINER at15 at15@dongyue.io

# Thanks to @zbintliff in https://github.com/kairosdb/kairosdb/issues/288
RUN apk upgrade libssl1.0 --update-cache && \
    apk add curl ca-certificates bash

RUN mkdir /opt; \
  cd /opt; \
  curl -L https://github.com/kairosdb/kairosdb/releases/download/v1.1.2/kairosdb-1.1.2-1.tar.gz | \
  tar zxvfp -

# Use H2
COPY kairosdb.h2.properties /opt/kairosdb/conf/kairosdb.properties

EXPOSE 4242 8080 2003 2004
ENTRYPOINT [ "/opt/kairosdb/bin/kairosdb.sh" ]
CMD [ "run" ]
