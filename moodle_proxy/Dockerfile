FROM ma-shared-lib:latest

RUN mkdir -p /go/src/github.com/ma-zero-trust-prototype/moodle_proxy

WORKDIR /go/src/github.com/ma-zero-trust-prototype/moodle_proxy

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["moodle_proxy"]

