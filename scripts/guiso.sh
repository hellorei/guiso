#!/usr/bin/env bash

set -e

# Get the directory of this script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Change to the project directory
cd "$DIR"

# Function to check dependencies
check_dependencies() {
  command -v make >/dev/null 2>&1 || { echo >&2 "make is required but not installed. Aborting."; exit 1; }
  # TODO Add more dependency checks here
}

# Function to print help
print_help() {
  make help
}

# Check dependencies
check_dependencies

# Handle commands
if [ $# -eq 0 ]; then
  print_help
elif [ "$1" = "help" ]; then
  print_help
else
  # Run make with all arguments passed to this script
  make "$@"
fi