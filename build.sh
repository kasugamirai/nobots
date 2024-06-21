RUN_NAME=nobot

# Create output directories
mkdir -p output/bin output/conf output/pkl

# Copy configuration files
cp -r conf/* output/conf

cp -r scripts/* output/

cp -r pkl/* output/pkl


# Copy and set permissions for bootstrap script
cp scripts/* output 2>/dev/null
chmod +x output/*

Set GOOS and GOARCH for Linux arm64
export GOOS=linux
export GOARCH=amd64

# Build the Go application
go build -o output/bin/${RUN_NAME} cmd/swiss/swiss.go