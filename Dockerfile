FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

ADD . .

RUN cd src && go build -v -o /bin/app ./cmd/apiserver


FROM alpine

RUN apk add --no-cache bash ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .

ENTRYPOINT ["/bin/app"]
