# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o main .

# Stage 2: Run (minimal runtime image)
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
# Copy config/.env if needed
# COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]