# 1. Aşama: Build (Derleme)
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Bağımlılıkları kopyala ve indir
COPY go.mod go.sum ./
RUN go mod download

# Kaynak kodları kopyala ve derle
COPY . .
RUN go build -o main .

# 2. Aşama: Run (Çalıştırma - Çok küçük bir imaj kullanıyoruz)
FROM alpine:latest
WORKDIR /app

# Build aşamasından çıkan çalışabilir dosyayı al
COPY --from=builder /app/main .
# Config dosyasını veya .env varsa onu da kopyalamak gerekebilir
# COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]