FROM golang:1.19.5

WORKDIR /app

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN golangci-lint run --timeout=2m ./...
CMD [ "go", "test", "./..." ]