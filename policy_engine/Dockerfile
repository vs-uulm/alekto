FROM ma-shared-lib:latest

RUN mkdir -p /go/src/github.com/ma-zero-trust-prototype/policy_engine

WORKDIR /go/src/github.com/ma-zero-trust-prototype/policy_engine

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["policy_engine"]

