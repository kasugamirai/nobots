RUN_NAME=nobot

# Create output directories
mkdir -p output/bin output/conf output/pkl

# Copy configuration files
cp -r conf/* output/conf

cp -r pkl/* output/pkl

cp -r scripts/* output/

# Copy and set permissions for bootstrap script
cp scripts/* output 2>/dev/null
chmod +x output/*

# Build the Go application
go build -o output/bin/${RUN_NAME} cmd/swiss/swiss.go