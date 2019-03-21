FROM golang
MAINTAINER tenling
RUN go get github.com/tenling/tenling-bot
WORKDIR /go/src/github.com/tenling/tenling-bot
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only
RUN go install -v ./...
RUN go build -o main .

CMD ["/usr/local/go/bin/go", "run", "main.go"]