# Stage 1: Build the application
FROM golang:1.23.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Install necessary dependencies (e.g., SSL certificates, etc.)
RUN apk --no-cache add ca-certificates

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./myapp .


# Stage 2: Create a minimal production image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the application port (if your app listens on port 8080)
EXPOSE 8080

# Run the application
CMD ["./myapp"]
