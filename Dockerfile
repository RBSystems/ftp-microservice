FROM golang:1.6

RUN mkdir -p /go/src/github.com/byuoitav
ADD . /go/src/github.com/byuoitav/ftp-microservice

WORKDIR /go/src/github.com/byuoitav/ftp-microservice
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/ftp-microservice"]

EXPOSE 8002
