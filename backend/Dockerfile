# Start from the official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install git and other necessary tools
RUN apk update && apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 3000

# Run the executable
CMD ["./main"]
