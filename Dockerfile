FROM golang:1.11
RUN go get github.com/rianby64/example-alia-redis
CMD [ "go run /go/src/github.com/rianby64/example-alia-redis/main.go" ]
