# Use official Go 1.23 image
FROM golang:1.23.0-alpine

# Set working directory
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose app port
EXPOSE 8000

# Run the built binary
CMD ["./main"]
