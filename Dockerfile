# ------------ Build Stage ------------
    FROM golang:1.23.5 as builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    RUN go build -o main ./cmd/server
    
    # ------------ Runtime Stage ------------
    FROM golang:1.23.5 as runner
    
    WORKDIR /app
    
    COPY --from=builder /app/main .
    
    # Set environment variables
    ARG PORT
    ARG NEON_CONNECTION_STRING
    
    ENV GIN_MODE=release
    ENV PORT=${PORT}
    ENV NEON_CONNECTION_STRING=${NEON_CONNECTION_STRING}
    
    EXPOSE ${PORT}
    
    CMD ["./main"]
    