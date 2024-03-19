# base image
FROM golang:latest

# set root path /app 
WORKDIR /app

# Copy go.mod and go.sum 
COPY go.mod go.sum ./

# Download all dependencies 
RUN go mod download

# Copy source code from current directory to working directory in container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# run binary file from build command
CMD ["./main"]