# Start with a base image containing Go
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app
# Copy the go.mod and go.sum files first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Set the command to run the binary
CMD ["./main"]
