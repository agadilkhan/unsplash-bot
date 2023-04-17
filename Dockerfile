FROM golang:1.19.0

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go
ENTRYPOINT ["/app/main"]