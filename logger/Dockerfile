FROM ma-shared-lib:latest

RUN mkdir -p /go/src/github.com/ma-zero-trust-prototype/logger

WORKDIR /go/src/github.com/ma-zero-trust-prototype/logger

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["logger"]

