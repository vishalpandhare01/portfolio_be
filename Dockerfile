# Use Golang base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 3000

# Run the binary when the container starts
CMD ["./main"]
