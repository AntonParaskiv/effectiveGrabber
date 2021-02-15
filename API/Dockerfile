FROM golang

ADD . /API
WORKDIR /API

RUN go test ./... && go build

EXPOSE 4000

ENTRYPOINT ./API
