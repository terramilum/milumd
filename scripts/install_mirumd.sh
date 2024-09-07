#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# application version like 0.1.3
APP_VERSION=$1
# application wasmvm version in go.mod file wasmvm tag like 1.2.3
WASMVM_VERSION=$2

if [[ -z "$APP_VERSION" || -z "$WASMVM_VERSION" ]]; then
  echo "Usage: $0 <APP_VERSION> <WASMVM_VERSION>"
  exit 1
fi

echo "Installing mirumd version $APP_VERSION with wasmvm version $WASMVM_VERSION"

# Update and upgrade the system
sudo apt update -y
sudo apt upgrade -y

# Install necessary packages
sudo apt install -y curl git jq lz4 build-essential unzip

# Define the path where mirumd will be installed
PATH_BIN=$HOME/code/bin

# Add the binary path to .bash_profile if not already included
if ! grep -q "$PATH_BIN" "$HOME/.bash_profile"; then
  echo "export PATH=\$PATH:$PATH_BIN" >> "$HOME/.bash_profile"
  source "$HOME/.bash_profile"
fi

source "$HOME/.bash_profile"

echo "Binary path: $PATH_BIN"

# Create the binary directory if it doesn't exist
mkdir -p "$PATH_BIN"

# Navigate to the binary directory
cd "$PATH_BIN"

# Download and extract the mirumd executable
wget https://github.com/terramirum/mirumd/releases/download/v$APP_VERSION/mirumd-$APP_VERSION-linux-amd64.tar.gz
tar -xvf mirumd-$APP_VERSION-linux-amd64.tar.gz

# Remove the downloaded tar.gz file
rm -f mirumd-$APP_VERSION-linux-amd64.tar.gz

# Ensure the mirumd file is executable
chmod +x mirumd

# Download libwasmvm and place it in /usr/lib
sudo wget -P /usr/lib https://github.com/CosmWasm/wasmvm/releases/download/v$WASMVM_VERSION/libwasmvm.x86_64.so

# Ensure the library is linked correctly
sudo ldconfig

# Verify that mirumd is available by checking its version
mirumd version
