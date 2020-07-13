FROM golang:1.14.4-alpine3.12

RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh
WORKDIR /app
COPY ./ /app
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

