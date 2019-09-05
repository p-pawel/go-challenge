FROM golang:1.12

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app
ENV GO111MODULE="on"

RUN go get -v
RUN go build

CMD ./go-challenge

EXPOSE 3000
