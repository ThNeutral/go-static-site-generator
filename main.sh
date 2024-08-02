#!/bin/bash

# Create a temporary directory
TEMP_DIR=$(mktemp -d)

# Check if mktemp failed
if [ ! -d "$TEMP_DIR" ]; then
  echo "Failed to create temp directory"
  exit 1
fi

# Define the name of the temporary executable
EXECUTABLE="$TEMP_DIR/temp_go_executable"

# Check if go source file is passed as argument
# if [ -z "$0/main.go" ]; then
#   echo "Usage: $0 <path_to_go_file>"
#   exit 1
# fi

GO_FILE="./main.go"

# Build the Go file into the temporary executable
go build -o "$EXECUTABLE" "$GO_FILE"

# Check if build was successful
if [ $? -ne 0 ]; then
  echo "Failed to build $GO_FILE"
  exit 1
fi

# Run the executable
"$EXECUTABLE"

# Clean up: remove the temporary directory
rm -rf "$TEMP_DIR"