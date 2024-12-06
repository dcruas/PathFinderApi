# Use the official Go image as the base image
FROM golang:1.23.3

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./

# Download all dependencies (creates a module cache)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
