FROM ma-shared-lib:latest

RUN mkdir -p /go/src/github.com/ma-zero-trust-prototype/sso_auth

WORKDIR /go/src/github.com/ma-zero-trust-prototype/sso_auth

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["sso_auth"]

