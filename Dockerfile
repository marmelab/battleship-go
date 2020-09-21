FROM golang:1.15.2

WORKDIR $GOPATH

# Download dependencies
RUN go get -u -v golang.org/x/lint/golint
