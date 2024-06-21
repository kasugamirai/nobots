# Use the official Golang image, which includes all the necessary build tools
# Make sure to choose the correct tag for the base image if you need a specific version of Go or base OS.
FROM arm64v8/golang:latest

# Set the environment variables for the Go application
ENV GOOS=linux \
    GOARCH=arm64 \
    CGO_ENABLED=1

# Set the working directory inside the container
WORKDIR /usr/src/myapp

# Copy the local package files to the container's workspace.
COPY . .

# Build the Go app
# Make sure to replace 'cmd/swiss/swiss.go' with your actual file path and 'myapp' with your desired binary name.
RUN go build -v  cmd/swiss/swiss.go

