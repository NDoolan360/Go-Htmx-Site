#!/bin/bash

# Relative to parent dir
cd "`dirname "$0"`"/..

# Detecting OS
if [[ "$(uname -s)" == "Darwin" ]]; then
  OS="macos"
elif [[ "$(expr substr $(uname -s) 1 5)" == "Linux" ]]; then
  OS="linux"
elif [[ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]]; then
  OS="windows"
elif [[ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]]; then
  OS="windows"
else
  echo "Unknown OS: $(uname -s)"
  exit 1
fi

# Detecting Platform
if [[ "$(uname -m)" == "arm64" ]]; then
  PLATFORM="arm64"
elif [[ "$(uname -m)" == "x86_64" ]]; then
  PLATFORM="x64"
elif [[ "$(uname -m)" == "armv7" ]]; then
  PLATFORM="armv7"
else
  echo "Unknown platform: $(uname -m)"
  exit 1
fi

mkdir resources
cd resources

# Download Tailwind CSS
TAILWIND_URL="https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-${OS}-${PLATFORM}"
curl -sLO "$TAILWIND_URL"
chmod +x "tailwindcss-${OS}-${PLATFORM}"
mv "tailwindcss-${OS}-${PLATFORM}" tailwindcss
