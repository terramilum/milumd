#!/bin/bash

# This is secret. Please change it when installing a full node.

# Check if password is provided via environment variable or as a script argument
if [ -z "$1" ] && [ -z "$PASSWORD" ]; then
    echo "Error: No password provided. Please provide a password as an argument or set the PASSWORD environment variable."
    exit 1
fi

# Assign password from argument if provided
PASSWORD=${1:-$PASSWORD}

# Check if the password is the placeholder "password"
if [ "$PASSWORD" == "password" ]; then
    echo 'Error: Password cannot be "password". Please choose a secure and unique password.'
    exit 1
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
echo "!!!!!!!! Store your mnemonic words to backup or import it into any Keplr wallet. !!!!!!!!"
