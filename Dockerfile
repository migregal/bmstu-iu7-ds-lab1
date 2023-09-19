FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

RUN apk add --no-cache make

ADD . .

RUN cd src && make build && mv bin/apiserver /bin


FROM alpine

RUN apk add --no-cache bash ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/apiserver .

ENTRYPOINT ["/bin/apiserver"]
