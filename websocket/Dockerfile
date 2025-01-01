FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o websocket websocket.go

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y libc6 && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/websocket .

EXPOSE 8081

CMD ["./websocket"]
