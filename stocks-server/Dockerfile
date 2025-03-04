# Use Go as the base image for building the application
FROM golang:1.23 AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker's build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o stock-restapi-server stock_server.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set up a working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/stock-restapi-server .

# Expose the application port
EXPOSE 8081

# Run the compiled binary
CMD ["./stock-restapi-server"]