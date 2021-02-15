FROM golang

ADD . /Grabber
WORKDIR /Grabber

RUN go test ./... && go build

ENTRYPOINT ./Grabber
