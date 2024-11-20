# Start from the official Golang image as a build stage
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source code into the working directory
COPY . .

# Build the Go application
RUN go build -o main cmd/server/main.go

# Start from a minimal image for the final stage
FROM alpine:latest

# Install SSL certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=builder /app/main .

# Expose the port used by your application
EXPOSE 7080

# Command to run the application
CMD ["./main"]
