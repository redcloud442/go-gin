
# Use official Go base image
FROM golang:1.22 as builder

# Set working directory inside container
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go app (target: cmd/server)
RUN go build -o main ./cmd/server

# Use a lightweight base image for runtime
FROM debian:bullseye-slim

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Set environment variables (optional)
ARG PORT
ARG NEON_CONNECTION_STRING
ENV GIN_MODE=release
ENV PORT=$PORT
ENV NEON_CONNECTION_STRING=$NEON_CONNECTION_STRING


# Expose the default Gin port
EXPOSE $PORT

# Run the binary
CMD ["./main"]
