#!/usr/bin/env bash

# initial a test chain genesis and keybase config
desmos keys add "test0" --home "build/ibc/ibc0" --keyring-backend "test"
desmos init testnet --chain-id "ibc0"  --home "build/ibc/ibc0"
desmos add-genesis-account "test0" "100000000000stake" --home "build/ibc/ibc0" --keyring-backend "test"
desmos gentx "test0" "1000000000stake" --chain-id "ibc0" --home "build/ibc/ibc0" --keyring-backend "test"
desmos collect-gentxs --home "build/ibc/ibc0"

#desmos start --home "build/ibc/ibc0"