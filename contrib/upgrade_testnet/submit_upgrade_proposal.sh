#!/usr/bin/env bash

NODES=$1
UPGRADE_NAME=$2
UPGRADE_HEIGHT=$3
NODE="http://localhost:26657"

# Make sure the chain is running
CNT=0
ITER=5
SLEEP=30

echo "===> Checking chain status"
while [ ${CNT} -lt $ITER ]; do
  echo "====> Test n. $CNT"
  curr_block=$(curl -s $NODE/status | jq -r '.result.sync_info.latest_block_height')

  if [ ! -z ${curr_block} ] ; then
    break
  fi

  echo "====> Chain is still offline. Sleeping..."
  let CNT=CNT+1
  sleep $SLEEP
done

echo "====> Chain is online. Ready to submit proposal"

CHAIN_ID=$(curl -s $NODE/status | jq -r '.result.node_info.network')
if [ -z "$CHAIN_ID" ]; then
  echo "Missing chain id"
  exit 1
fi

TX_FLAGS="--from node0 --keyring-backend test --home ./build/node0/desmos --chain-id $CHAIN_ID --yes --fees 100udaric"

# Import all nodes keys into node0
echo ""
echo "===> Importing keys into Node 0 keystore"
for ((i = 1; i < $NODES; i++)); do
  echo "====> Node $i"
  NODE_SECRET=$(cat "build/node$i/desmos/key_seed.json" | jq .secret -r)
  echo $NODE_SECRET | desmos keys add "node$i" --recover --home ./build/node0/desmos --keyring-backend test >/dev/null 2>&1
done

sleep $SLEEP

echo ""
echo "===> Submitting upgrade proposal"
RESULT=$(desmos tx gov submit-proposal \
  software-upgrade $UPGRADE_NAME \
  --title Upgrade \
  --description Description \
  --upgrade-height $UPGRADE_HEIGHT \
  $TX_FLAGS 2>&1)
TX_HASH=$(echo "$RESULT" | grep txhash | sed -e 's/txhash: //')
if [ -z "$TX_HASH" ]; then
  echo "Error while submitting transaction: $RESULT"
  exit 1
fi

sleep 6s

echo ""
echo "===> Getting proposal id"
PROPOSAL_ID=$(desmos q tx $TX_HASH --output json 2>&1 | jq .logs[0].events[2].attributes[0].value -r)
echo "Proposal ID: $PROPOSAL_ID"

echo ""
echo "===> Depositing proposal"
desmos tx gov deposit $PROPOSAL_ID 10000000udaric $TX_FLAGS >/dev/null 2>&1

sleep 6s

echo ""
echo "===> Voting proposal"
for ((i = 0; i < $NODES; i++)); do
  echo "====> Node $i"
  desmos tx gov vote $PROPOSAL_ID yes $TX_FLAGS --from "node$i" >/dev/null 2>&1
done

sleep 6s

echo ""
echo "===> Waiting for upgrade height ($UPGRADE_HEIGHT)"
while true; do
  curr_block=$(curl -s $NODE/status | jq -r '.result.sync_info.latest_block_height')
  docker-compose -f $(pwd)/contrib/upgrade_testnet/docker-compose.yml logs --tail=1

  if [ ! -z ${curr_block} ] ; then
    echo "Current block: ${curr_block}"
  fi

  if [ ! -z ${curr_block} ] && [ ${curr_block} -gt ${UPGRADE_HEIGHT} ]; then
    echo "Upgrade height passed. Upgrade was successful!"
    exit 0
  fi

  sleep $SLEEP
done