FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o poker-server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/poker-server .

CMD ["./poker-server"]
