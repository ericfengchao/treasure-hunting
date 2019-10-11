FROM golang:1.12

RUN apt-get update || exit 0
RUN apt-get upgrade -y
RUN apt-get install vim sudo curl -y

RUN mkdir -p /go/src/github.com/ericfengchao/treasure-hunting

WORKDIR /go/src/github.com/ericfengchao/treasure-hunting
COPY . /go/src/github.com/ericfengchao/treasure-hunting
VOLUME /go/src/github.com/ericfengchao/treasure-hunting

RUN go build -o $GOPATH/bin/tracker_server github.com/ericfengchao/treasure-hunting/cmd/tracker/server

ENTRYPOINT ["tracker_server", "50050", "5", "10"]

EXPOSE 50050

#CMD ["tracker_server", "50050", "5", "10"]
