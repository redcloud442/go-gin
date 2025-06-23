# ------------ Build Stage ------------

# Use Go 1.23.5 as the builder image
FROM golang:1.23.5 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application (adjust path if needed)
RUN go build -o main ./cmd/server

# ------------ Runtime Stage ------------

# Use a lightweight Debian image for production
FROM debian:bullseye-slim

# Set working directory
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/main .

# Set environment variables
ARG PORT
ARG NEON_CONNECTION_STRING

ENV GIN_MODE=release
ENV PORT=${PORT}
ENV NEON_CONNECTION_STRING=${NEON_CONNECTION_STRING}

# Expose the specified port
EXPOSE ${PORT}

# Command to run the binary
CMD ["./main"]
