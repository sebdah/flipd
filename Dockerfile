# vim: ft=dockerfile :
FROM golang:1.7-alpine

RUN mkdir -p /go/src/github.com/sebdah/flipd/ && \
	apk -U add git bash && \
	go get -v github.com/alecthomas/gometalinter && \
	gometalinter --install

ADD . /go/src/github.com/sebdah/flipd/
WORKDIR /go/src/github.com/sebdah/flipd/

RUN go build -o /go/bin/flipd-thrift ./source/cmd/thrift
RUN go build -o /go/bin/flipd-http ./source/cmd/http

EXPOSE 9090

CMD ["true"]
