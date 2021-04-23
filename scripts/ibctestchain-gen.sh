#!/usr/bin/env bash

id=$1
accountNum=$2
HOMEDIR=$3

# initial a test chain genesis and keybase config
 desmos init testnet --chain-id "ibc$id"  --home "$HOMEDIR/ibc$id"

for (( i = 0; i < $accountNum; i++ ))
do
    desmos keys add "test$id-$i" --home "$HOMEDIR/ibc$id" --keyring-backend "test"
    desmos add-genesis-account "test$id-$i" "100000000000desmos,100000000000stake" --home "$HOMEDIR/ibc$id" --keyring-backend "test"
done

desmos gentx "test$id-0" "500000000stake" --chain-id "ibc$id" --home "$HOMEDIR/ibc$id" --keyring-backend "test"

desmos collect-gentxs --home "$HOMEDIR/ibc$id"