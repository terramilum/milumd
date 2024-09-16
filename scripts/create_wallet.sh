#!/bin/bash

# This is secret. Please change it when installing a full node.

# Generate a strong random 32-byte password if none is provided
if [ -z "$1" ] && [ -z "$PASSWORD" ]; then
    echo "No password provided. Generating a secure random password..."
    PASSWORD=$(tr -dc 'a-zA-Z0-9' < /dev/urandom | head -c 32)
    echo "Generated password: $PASSWORD"
else
    # Assign password from argument or environment variable if provided
    PASSWORD=${1:-$PASSWORD}
fi

echo "Wallet creation is starting..."

# Check if the validator wallet already exists
if ! mirumd keys show validator > /dev/null 2>&1; then
    # Create a new validator wallet
    if (echo "$PASSWORD"; echo "$PASSWORD") | mirumd keys add validator; then
        echo "Validator wallet successfully created."
    else
        echo "Error: Failed to create validator wallet."
        exit 1
    fi
else
    echo "Validator wallet already exists."
fi

# Instructions for the user
echo "!!!!!!!! Store your mnemonic words and wallet password. !!!!!!!!"
