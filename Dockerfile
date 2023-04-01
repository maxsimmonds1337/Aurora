# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Compile the Go binary
RUN go build -o main .

# Set the entry point for the container
ENTRYPOINT ["/app/main"]
