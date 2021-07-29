#!/usr/bin/env bash

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

desmos tx ibc channel open-init transfer transfer connection-1 $TX_FLAGS

echo "Open channel completed"

desmos tx ibc-transfer transfer transfer channel-0 desmos1zfuyr8jd65d5lhhnulg5ze3jwgx4slrfx57esp 100udaric $TX_FLAGS

echo "Transfer transaction successful"
