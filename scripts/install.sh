#!/bin/bash

# Default values
PLATFORM="x64"
OS="linux"

# Parse optional command-line arguments
while [[ "$#" -gt 0 ]]; do
  case $1 in
    -p|--platform)
      PLATFORM="$2"
      shift
      ;;
    -o|--os)
      OS="$2"
      shift
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
  shift
done

# Build Tailwind CSS download URL
URL="https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-${OS}-${PLATFORM}"

# Download Tailwind CSS
curl -sLO "$URL"
chmod +x "tailwindcss-${OS}-${PLATFORM}"
mv "tailwindcss-${OS}-${PLATFORM}" tailwindcss
