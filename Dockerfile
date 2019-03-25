FROM golang:1.12.1-stretch

WORKDIR /go/src/github.com/fwchen

COPY . .

ENV GO111MODULE=on

RUN go get

RUN go build main.go

CMD ["/go/src/github.com/fwchen/main"]
