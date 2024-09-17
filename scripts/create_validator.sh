#!/bin/bash
set -e

PASSWORD=$1

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
  echo "Error: You must provide a validator moniker name for your wallet."
  exit 1
fi

# Ensure a wallet password is provided
if [ -z "$PASSWORD" ]; then
  echo "Error: You must provide a wallet password as the first argument."
  exit 1
fi

# Execute the create-validator command with password input and auto-signing
(echo "$PASSWORD"; echo "y") | mirumd tx staking create-validator \
  --amount=25000000000000MIRUM \
  --pubkey=$(mirumd tendermint show-validator) \
  --moniker="$NAME" \
  --chain-id=mirum-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="25000000000000" \
  --gas="2500000" \
  --gas-prices="0.1MIRUM" \
  --from=$NAME \
  --yes
