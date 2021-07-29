#!/usr/bin/env bash

CHAIN_ID=$(curl -s $NODE/status | jq -r '.result.node_info.network')
TX_FLAGS="--from node0 --keyring-backend test --home ./build/node0/desmos --chain-id $CHAIN_ID --yes --fees 100udaric"

desmos tx ibc channel open-init transfer transfer connection-1 $TX_FLAGS

echo "Open channel completed"

desmos tx ibc-transfer transfer transfer channel-0 desmos1zfuyr8jd65d5lhhnulg5ze3jwgx4slrfx57esp 100udaric $TX_FLAGS

echo "Transfer transaction successful"
