#!/bin/bash
set -e

# Moniker (node name). Visible on block explorers as the validator name.
MONIKER=$1

# Check if the MONIKER parameter is provided. If not, exit with an error.
if [ -z "$MONIKER" ]; then
  echo "Error: You must provide a moniker name for your node. Use the following syntax:"
  echo "bash <(curl -s 'https://raw.githubusercontent.com/terramirum/mirumd/main/scripts/install_full_node.sh') <moniker-name>"
  exit 1
fi

# Check if the moniker is set to the placeholder "moniker"
if [ "$MONIKER" == "moniker" ]; then
    echo 'Error: Moniker name cannot be "moniker". Please choose a unique moniker.'
    exit 1
fi

# Default home path for mirumd
HOMEP=${HOMEP:-~/.mirumd}

# Chain ID for the network. Must match the chain ID in the networks repo.
CHAIN_ID=${CHAIN_ID:-mirum-1}

# Configuration file paths
FILENAME=${FILENAME:-"$HOMEP"/config/genesis.json}
CONFIG=${CONFIG:-"$HOMEP"/config/config.toml}
APPTOML=${APPTOML:-"$HOMEP"/config/app.toml}
CLIENTTOML=${CLIENTTOML:-"$HOMEP"/config/client.toml}

# Set this to true for production, false for development/testing.
IS_PROD=${IS_PROD:-true}

# Remove any existing installation in the home path
rm -rf "$HOMEP"

# Initialize the node with the specified moniker and chain ID
mirumd init --chain-id "$CHAIN_ID" "$MONIKER"

# Clean up old networks folder and clone the latest networks repository
rm -rf "$HOME"/networks
git clone https://github.com/terramirum/networks.git "$HOME"/networks

# Path to the genesis file in the networks repository
SOURCE_GENESIS="$HOME"/networks/"$CHAIN_ID"/genesis.json

# Check if the genesis file exists, otherwise exit with an error
if [ ! -f "$SOURCE_GENESIS" ]; then
  echo "Error: Genesis file not found at $SOURCE_GENESIS"
  exit 1
fi

# Copy the genesis file to the config folder
cp -rf "$SOURCE_GENESIS" "$FILENAME"

# Update configurations to allow external connections (0.0.0.0 instead of localhost)
for file in "$CONFIG" "$APPTOML" "$CLIENTTOML"; do
  sed -i 's/localhost/0.0.0.0/' "$file"
  sed -i 's/127.0.0.1/0.0.0.0/' "$file"
done

# If this is a production environment, adjust the logging level
if [ "$IS_PROD" = true ]; then
  sed -i 's/log_level = "info"/log_level = "main:info,state:info,*:error"/' "$CONFIG"
fi

# Path to the persistent peers file in the networks repository
PERSISTENT_PEERS_PATH="$HOME"/networks/"$CHAIN_ID"/persistent_peers

# Check if the persistent peers file exists, otherwise exit with an error
if [ ! -f "$PERSISTENT_PEERS_PATH" ]; then
  echo "Error: Persistent peers file not found at $PERSISTENT_PEERS_PATH"
  exit 1
fi

# Get the list of persistent peers and update the config
PERSISTENT_PEERS=$(cat "$PERSISTENT_PEERS_PATH")
sed -i 's/timeout_commit = "5s"/timeout_commit = "3s"/' "$CONFIG"
sed -i "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" "$CONFIG"

# Display helpful messages to the user
echo "NOTICE !!!!!"
echo "Chain is starting. Node will be synchronized with name $MONIKER."
echo "Follow these steps:"
echo "1. Execute create_wallet.sh to create a wallet. Make sure to back up your mnemonic words."
echo "2. Execute create_validator.sh to become a validator node."
echo ""
echo "After the chain is started and synchronized, stop the program."
echo 'Use "nohup mirumd start &" to run the process in the background.'
echo "Enjoy !!!!!!!!"
echo ""
echo "Use the following command to start the chain. If you use a different home folder, add the --home flag:"
echo "mirumd start"

# Pause for a few seconds before exiting
sleep 5

# Start the mirumd process with the specified home path (optional)
# mirumd start --home "$HOME"
