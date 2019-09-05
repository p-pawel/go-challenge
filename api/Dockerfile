FROM golang:1.8

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v
RUN go build

CMD ./go-challenge

EXPOSE 3000
