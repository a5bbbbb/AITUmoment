# Build stage
FROM golang:1.23.4 as builder

WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

# Add ca-certificates for HTTPS and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy necessary application files
COPY view/ view/
COPY db/migrations/ db/migrations/

# Create directory for logs
RUN mkdir -p /app/logs

EXPOSE 8080

CMD ["./main"]
