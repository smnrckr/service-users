# 1. Aşama: Build aşaması
FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Statik derleme ile Linux için binary üretelim
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

# 2. Aşama: Çalıştırma aşaması (Alpine yerine Debian)
FROM debian:stable-slim

WORKDIR /root/

COPY --from=builder /app/server .

# Çalıştırma komutu
CMD ["./server"]