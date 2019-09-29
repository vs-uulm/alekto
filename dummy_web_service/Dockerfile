FROM golang:1.10.4

WORKDIR /go/src/github.com/ma-zero-trust-prototype/dummy_web_service
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["dummy_web_service"]