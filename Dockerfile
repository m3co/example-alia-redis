FROM golang:1.11
RUN go get github.com/rianby64/example-alia-redis
WORKDIR /go/src/github.com/rianby64/example-alia-redis
RUN go build -o server .
CMD [ "/go/src/github.com/rianby64/example-alia-redis/server", "-m", "disk" ]