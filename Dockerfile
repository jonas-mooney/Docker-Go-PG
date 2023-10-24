# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o main .

FROM alpine:3.14

# WORKDIR /app

# Expose the port on which the server will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
