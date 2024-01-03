#!/bin/sh
#set -o errexit -o nounset -o pipefail

# DEFAULT_HOMEP=${HOMEP:-~/.mirumd}
# HOMEP=${HOMEP:-/mnt/volume_fra1_02/terramirum}
HOMEP=${HOMEP:-~/.mirumd}
PASSWORD=${PASSWORD:-12345678}
STAKE=${STAKE_TOKEN:-TRM}
FEE=${FEE_TOKEN:-uTRM}
CHAIN_ID=${CHAIN_ID:-terramirum-1}
MONIKER=${MONIKER:-main}
GENESIS=${GENESIS:-"$HOMEP"/config/genesis.json}
APPTOML=${APPTOML:-"$HOMEP"/config/app.toml}
CLIENTTOML=${CLIENTTOML:-"$HOMEP"/config/client.toml}
CONFIG=${CONFIG:-"$HOMEP"/config/config.toml}
IS_PROD=${IS_PROD:-true}


rm -rf "$HOMEP"

mirumd init --chain-id "$CHAIN_ID" "$MONIKER" --home "$HOMEP"
# staking/governance token is hardcoded in config, change this
sed -i "s/\"stake\"/\"$STAKE\"/" $GENESIS
# this is essential for sub-1s block times (or header times go crazy)
if grep -F "time_iota_ms" $GENESIS
then 
    sed -i 's/"time_iota_ms": "1000"/"time_iota_ms": "500"/' $GENESIS
fi

apt update
apt install -y jq

# to enable the api server
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' $APPTOML
# to change the voting_period
jq '.app_state.gov.voting_params.voting_period = "600s"' $GENESIS > temp.json && mv temp.json $GENESIS

# to change the inflation
jq '.app_state.mint.minter.inflation = "0.010000000000000000"' $GENESIS > temp.json && mv temp.json $GENESIS
jq '.app_state.mint.params.inflation_rate_change = "0.010000000000000000"' $GENESIS > temp.json && mv temp.json $GENESIS
jq '.app_state.mint.params.inflation_max = "0.020000000000000000"' $GENESIS > temp.json && mv temp.json $GENESIS
jq '.app_state.mint.params.inflation_min = "0.001000000000000000"' $GENESIS > temp.json && mv temp.json $GENESIS

# making 1 sec block time.
sed -i 's/timeout_commit = "5s"/timeout_commit = "3s"/' $CONFIG

for file in "$CONFIG" "$APPTOML" "$CLIENTTOML"; do
    sed -i 's/localhost/0.0.0.0/' "$file"
    sed -i 's/127.0.0.1/0.0.0.0/' "$file"
done

if [ "$IS_PROD" = true ]; then
    sed -i 's/log_level = "info"/log_level = "main:info,state:info,*:error"/' $CONFIG 
fi

if ! mirumd keys show validator --home "$HOMEP"; then
   (echo "$PASSWORD"; echo "$PASSWORD") | mirumd keys add validator --home "$HOMEP"
fi
# hardcode the validator account for this instance
echo "$PASSWORD" | mirumd genesis add-genesis-account validator "100000000000000000$STAKE" --home "$HOMEP"

# submit a genesis validator tx
## Workraround for https://github.com/cosmos/cosmos-sdk/issues/8251
(echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | mirumd genesis gentx validator "50000000000000$STAKE" --chain-id="$CHAIN_ID" --amount="50000000000000$STAKE" --home "$HOMEP"
## should be:
# (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | mirumd gentx validator "100000000000$STAKE" --chain-id="$CHAIN_ID"
mirumd genesis collect-gentxs --home "$HOMEP"
