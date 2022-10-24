FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/fizzbuzz-server ./cmd/server.go

CMD ["fizzbuzz-server", "0.0.0.0:80"]
