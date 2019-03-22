FROM golang
MAINTAINER tenling
WORKDIR /go/src/github.com/tenling/tenling-bot
ADD . .
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only
RUN go install -v ./...
RUN go build -o main .

CMD ["/usr/local/go/bin/go", "run", "main.go"]