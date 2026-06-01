# Stage 1: Build the Go binary
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy dependency files and download
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Stage 2: Create a minimal production image
FROM alpine:latest

# Install CA certificates for HTTPS (needed for LiveKit/Postgres/S3)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/api .

# Expose the API port
EXPOSE 8080

# Command to run
CMD ["./api"]
