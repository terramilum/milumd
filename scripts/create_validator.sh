#!/bin/bash

# Check if MONIKER is provided as the first argument
if [ -z "$1" ]; then
    echo "Error: No moniker provided. Please provide a moniker as the first argument."
    exit 1
fi

MONIKER=$1

# Check if the moniker is set to the placeholder "moniker"
if [ "$MONIKER" == "moniker" ]; then
    echo 'Error: Moniker name cannot be "moniker". Please choose a unique moniker.'
    exit 1
fi

# Execute the create-validator command
mirumd tx staking create-validator \
  --amount=25000000000000MIRUM \
  --pubkey=$(mirumd tendermint show-validator) \
  --moniker="$MONIKER" \
  --chain-id=mirum-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="25000000000000" \
  --gas="2500000" \
  --gas-prices="0.01MIRUM" \
  --from=validator
