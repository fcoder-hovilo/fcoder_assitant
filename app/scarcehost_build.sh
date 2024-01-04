#!/bin/bash

# Set the Go binary name
binary_name="bot_fcoder"

# Check if Go is installed
if ! command -v go &> /dev/null; then
  echo "Go is not installed. Installing Go..."

  # Install Go using apt-get
  apt-get update
  apt-get install -y golang
  if [ $? -ne 0 ]; then
    echo "Failed to install Go. Please check your system and try installing Go manually."
    exit 1
  fi
fi

# Build the Go program
go build -o $binary_name main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful. Running the program..."
  # Run the compiled executable
  ./$binary_name
else
  echo "Build failed. Please check your Go code for errors."
fi
