#!/bin/bash

$(go clean -testcache)

# Function to run tests in a given directory
run_tests() {
  local dir=$1
  echo "Running tests in ${dir}..."
  cd "${dir}"
  result=$(go test ./...)
  if [[ $? -eq 0 ]]; then
    echo "All tests passed in ${dir}!"
  else
    echo "Some tests failed in ${dir}. Check the output below for details."
  fi
  echo "$result"
  cd - > /dev/null
}

# Define the directories containing the packages
textnode_dir="./internals/textnode"
htmlnode_dir="./internals/htmlnode"
splitter_dir="./internals/splitter"

# Run tests for each package
run_tests "$textnode_dir"
run_tests "$htmlnode_dir"
run_tests "$splitter_dir"