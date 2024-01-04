#!/bin/bash

# Set the Go binary name
binary_name="bot_fcoder"

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
