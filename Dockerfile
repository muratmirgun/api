FROM golang:1.15
RUN go get github.com/muratmirgun/api
# create a working directory
WORKDIR /go/src/github.com/muratmirgun/api
# add source code

# run main.go
CMD ["go", "run", "main.go"]