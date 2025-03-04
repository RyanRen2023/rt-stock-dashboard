# Use the official Go image as the base image
FROM golang:1.23 AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o stock-server server.go

# Use a minimal base image to reduce final image size
FROM alpine:latest

# Set up a working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/stock-server .

# Expose the application port
EXPOSE 8080

# Run the compiled binary
CMD ["./stock-server"]