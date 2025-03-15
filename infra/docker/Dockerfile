FROM golang:1.23-alpine

ARG CONFIG_FILE

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/screen-time-limiter/main.go

EXPOSE 8081

CMD ["./server"]
