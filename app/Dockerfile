FROM golang:1.14.4-alpine3.12

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/pskclub/tendermint-domain-manament/app
ADD . /go/src/github.com/pskclub/tendermint-domain-manament/app
RUN go build
CMD ./app
