#!/usr/bin/env bash

id=$1
ACCOUNTNUM=$2
HOMEDIR=$3

# initial a test chain genesis and keybase config
desmos init testnet --chain-id "ibc$id"  --home "$HOMEDIR/ibc$id"

# Add test node key
desmos keys add "node$id" --home "$HOMEDIR/ibc$id" --keyring-backend "test" --no-backup
desmos add-genesis-account "node$id" "100000000000desmos,100000000000stake" --home "$HOMEDIR/ibc$id" --keyring-backend "test"

# Add test account keys
for (( i = 0; i < $ACCOUNTNUM; i++ ))
do
    desmos keys add "test$id-$i" --home "$HOMEDIR/ibc$id" --keyring-backend "test" --no-backup
    desmos add-genesis-account "test$id-$i" "100000000000desmos,100000000000stake" --home "$HOMEDIR/ibc$id" --keyring-backend "test"
done

# Add relayer account defiened in relayer-config to genesis
desmos add-genesis-account "desmos1punhxfyxvnwup70mc6cz9cmuxu74ars7x6hgjm" "100000000000desmos,100000000000stake" --home "$HOMEDIR/ibc$id"

# Create validator set to genesis
desmos gentx "node$id" "500000000stake" --chain-id "ibc$id" --home "$HOMEDIR/ibc$id" --keyring-backend "test" --node-id "node$id"
desmos collect-gentxs --home "$HOMEDIR/ibc$id"