FROM golang:1.15
RUN go get github.com/muratmirgun/api
WORKDIR /go/src/api
RUN go run main.go
