# Use the official Golang image as a builder
FROM golang:1.20 AS builder

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    gcc \
    libvips-dev \
    git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN go build -o server-app .

# Start a new stage with the necessary dependencies installed
FROM ubuntu:latest

# Install libvips
RUN apt-get update && apt-get install -y \
    libvips

# Copy the built binary from the builder stage
COPY --from=builder /app/server-app /usr/local/bin/server-app
RUN chmod +x /usr/local/bin/server-app

# Set the entry point to the binary
ENTRYPOINT ["/usr/local/bin/server-app"]

# Expose the port that the server listens on
EXPOSE 8080
