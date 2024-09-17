#!/bin/bash
set -e

NAME=$1

#!/bin/bash

# Path to the TOML file
CONFIG_FILE="$HOME/.mirumd/config/config.toml"
# Check if config file exists
if [ -f "$CONFIG_FILE" ]; then
    # Extract the moniker value
    moniker=$(grep '^moniker' "$CONFIG_FILE" | cut -d'=' -f2 | sed 's/[ "]//g')
    if [ -n "$moniker" ]; then
        echo "Moniker found: $moniker"
        NAME="$moniker"
    fi
fi

if [ -z "$NAME" ]; then
  echo "Error: You must provide a wallet name for your wallet."
  exit 1
fi

PASSWORD=$(tr -dc 'a-zA-Z0-9' < /dev/urandom | head -c 32)
# Set password manually
# PASSWORD=${1:-$PASSWORD}

HOSTIP=$(ip addr show eth0 | grep 'inet ' | awk '{print $2}' | cut -d'/' -f1)

echo "Host Ip: $HOSTIP"
echo "Wallet password: $PASSWORD"

# Check if the validator wallet already exists
if ! mirumd keys show $NAME > /dev/null 2>&1; then
    # Create a new validator wallet
    if (echo "$PASSWORD"; echo "$PASSWORD") | mirumd keys add $NAME; then
        echo "Validator wallet successfully created."
        # Instructions for the user
        echo "!!!!!!!! Store your mnemonic words and wallet password. !!!!!!!!"
    else
        echo "Error: Failed to create validator wallet."
        exit 1
    fi
else
    echo "Validator wallet already exists."
    mirumd keys show $NAME
fi

