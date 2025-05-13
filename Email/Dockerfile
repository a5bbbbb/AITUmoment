# Build stage
FROM golang:1.24.2 as builder
WORKDIR /app
# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download
# Copy the rest of the application
COPY . .
COPY .env .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o email-service ./cmd/main.go

# Final stage
FROM alpine:latest
# Add ca-certificates for HTTPS and timezone data
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
# Copy the binary from builder
COPY --from=builder /app/email-service .
# Create directory for logs (will be mounted as volume)
RUN mkdir -p /app/logs && chmod 777 /app/logs
CMD ["./email-service"]



