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

# Start from an official lightweight base image
FROM debian:bullseye-slim

# Install necessary packages, including Google Chrome
RUN apt-get update && apt-get install -y \
    wget \
    ca-certificates \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libatspi2.0-0 \
    libcups2 \
    libdbus-glib-1-2 \
    libxcomposite1 \
    libxrandr2 \
    libgbm-dev \
    && wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list' \
    && apt-get update && apt-get install -y google-chrome-stable \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=builder /app/main .

# Expose the port used by your application
EXPOSE 8080

# Command to run the application
CMD ["./main"]
