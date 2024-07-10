# Use the official Golang image as a build stage
FROM golang:1.23rc1-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the files to the working directory
COPY . .

# Compile the Go program
RUN go build -o server server.go

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary and the .env file from the builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/.env .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the Go application
CMD ["./server"]