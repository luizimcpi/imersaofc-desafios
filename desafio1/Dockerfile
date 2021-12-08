FROM golang:1.17

WORKDIR /go/src

RUN go run cmd/main.go

CMD ["tail", "-f", "/dev/null"]