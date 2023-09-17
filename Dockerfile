FROM alpine
LABEL maintainer="Gregory @migregal Mironov"

RUN apk add --no-cache bash ca-certificates

COPY src/bin/apiserver /bin/apiserver

RUN chmod 755 /bin/apiserver

CMD exec /bin/apiserver
