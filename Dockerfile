FROM golang:1.16
# add source code
RUN go get github.com/muratmirgun/api
# create a working directory
WORKDIR /go/src/github.com/muratmirgun/api


# run main.go
CMD ["go", "run", "main.go"]