HOMEP=${HOMEP:-~/.mirumd}
# give any key name for full node name. This will be visible at block explorer as a validator name. 
# should be changed. If no changes, you can change it at config.toml file with moniker name.
MONIKER=${MONIKER:-nodeks.com}
# chain id to replace genesis file with existing one.
# must be correct chain id equal to gitup folder name under networks repository.
CHAIN_ID=${CHAIN_ID:-mirum-1}
# configuration file names. no need to change
FILENAME=${FILENAME:-"$HOMEP"/config/genesis.json}
CONFIG=${CONFIG:-"$HOMEP"/config/config.toml}  
APPTOML=${APPTOML:-"$HOMEP"/config/app.toml}
CLIENTTOML=${CLIENTTOML:-"$HOMEP"/config/client.toml} 
IS_PROD=${IS_PROD:-true}

rm -rf "$HOMEP"

mirumd init --chain-id "$CHAIN_ID" "$MONIKER"

rm -rf "$HOME"/networks

git clone https://github.com/terramirum/networks.git "$HOME"/networks

SOURCE_GENESIS="$HOME"/networks/"$CHAIN_ID"/genesis.json

result=$(stat $SOURCE_GENESIS)
if [ $? -ne 0 ]; then
  echo "Error: genesis file not found"
  exit 1
fi 

cp -rf $SOURCE_GENESIS $FILENAME 

for file in "$CONFIG" "$APPTOML" "$CLIENTTOML"; do
    sed -i 's/localhost/0.0.0.0/' "$file"
    sed -i 's/127.0.0.1/0.0.0.0/' "$file"
done

if [ "$IS_PROD" = true ]; then
    sed -i 's/log_level = "info"/log_level = "main:info,state:info,*:error"/' $CONFIG 
fi

PERSISTENT_PEERS_PATH="$HOME"/networks/"$CHAIN_ID"/persistent_peers

result=$(stat $PERSISTENT_PEERS_PATH)
if [ $? -ne 0 ]; then
  echo "Error: genesis file not found"
  exit 1
fi   

PERSISTENT_PEERS=$(cat $PERSISTENT_PEERS_PATH)

# making 1 sec block time.
sed -i 's/timeout_commit = "5s"/timeout_commit = "3s"/' $CONFIG
sed -i "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" $CONFIG

echo "NOTICE !!!!!"

echo "Chain is starting. Node will be synchronized with name " $MONIKER
echo "Follow belowing step."
echo "1. Execute create_wallet.sh to create wallet. Please backup your mnemonic words"
echo "2. Execute create_validator.sh to become validator node."

echo "After chain start and being syched, then stop the program."
echo 'Use "nohup mirumd start &" to run process at background. '
echo "Enjoy !!!!!!!!!"

echo "Use this command to start chain. If you use different home folder then add --home flag"

echo "mirumd start"

sleep 5

# mirumd start --home "$HOME"