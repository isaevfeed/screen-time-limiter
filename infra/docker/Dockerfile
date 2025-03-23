FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/screen-time-limiter/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/migrations ./migrations

EXPOSE server

CMD ["./build"]