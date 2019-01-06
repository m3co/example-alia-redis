FROM golang:1.11
RUN go get github.com/rianby64/example-alia-redis-server
WORKDIR /go/src/github.com/rianby64/example-alia-redis-server
RUN go build -o server .
CMD [ "/go/src/github.com/rianby64/example-alia-redis-server/server", "-m", "disk" ]
