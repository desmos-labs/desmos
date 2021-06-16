#!/bin/bash
shopt -s expand_aliases

echo "===> Deleting previous data"
rm -r -f $HOME/.desmos

echo "===> Initializing chain"
desmos init testchain --chain-id testchain

echo "===> Editing genesis.json"
sed -i .bak 's/stake/udaric/g' $HOME/.desmos/config/genesis.json
sed -i .bak 's/"voting_period": "172800s"/"voting_period": "120s"/g' ~/.desmos/config/genesis.json
sed -i .bak 's/max_subscriptions_per_client = 5/max_subscriptions_per_client = 20/g' ~/.desmos/config/config.toml

sed -i .bak 's/pruning = "default"/pruning = "custom"/g' ~/.desmos/config/app.toml
sed -i .bak 's/pruning-keep-recent = "0"/pruning-keep-recent = "362880"/g' ~/.desmos/config/app.toml
sed -i .bak 's/pruning-interval = "0"/pruning-interval = "100"/g' ~/.desmos/config/app.toml

echo "===> Creating genesis accounts"
desmos add-genesis-account jack 100000000000000000000udaric
desmos add-genesis-account desmos16f9wz7yg44pjfhxyn22kycs0qjy778ng877usl 10000000udaric
desmos add-genesis-account desmos1ev4zxv2jpwsnw5ffkl0cfkag53aug52snr7uwy 100000000000udaric

echo "===> Collecting genesis trasanctions"
desmos gentx jack 1000000000udaric --chain-id testchain
desmos collect-gentxs

echo "===> Starting chain"
desmos start
