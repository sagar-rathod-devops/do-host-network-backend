# 1. Use official Golang image as base
FROM golang:1.21-alpine

# 2. Set the Current Working Directory inside the container
WORKDIR /app

# 3. Copy go.mod and go.sum files
COPY go.mod go.sum ./

# 4. Download all Go modules
RUN go mod download

# 5. Copy the source code
COPY . .

# 6. Build the Go app
RUN go build -o main .

# 7. Expose port (change this if needed)
EXPOSE 8000

# 8. Command to run the executable
CMD ["./main"]
