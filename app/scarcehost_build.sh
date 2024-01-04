#!/bin/bash

# Set the Go binary name
binary_name="CHANGE_ME"


apt install libc6

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Running the program..."
  # Run the compiled executable
  ./$binary_name
else
  echo "Build failed. Please check your Go code for errors."
fi
