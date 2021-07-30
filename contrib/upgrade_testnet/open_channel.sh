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

TX_FLAGS="--keyring-backend test --home ./build/node0/desmos --chain-id $CHAIN_ID --yes --fees 100udaric --broadcast-mode block"

# Create a new Alice key
ALICE_KEY="alice"
ALICE_ADDRESS="desmos1yvp004qpvxnus27mp3qsdw64fcm3rhgt2zgdrc"
ALICE_MNEMONIC="volume tape print syrup exchange tuna dumb call width place flat glide update mesh puzzle flash tilt rifle identify crane entry enact mansion grant"
echo $ALICE_MNEMONIC | desmos keys add $ALICE_KEY --recover --home ./build/node0/desmos --keyring-backend test >/dev/null 2>&1

# Send Alice 300udaric
echo "Sending amount to Alice"
desmos tx bank send node0 $ALICE_ADDRESS 300udaric $TX_FLAGS

# Perform a failing transfer transaction from Alice sending 50udaric (amount < balance)
echo "Sending transfer transaction with 50udaric"
desmos tx ibc-transfer transfer transfer channel-0 desmos1zfuyr8jd65d5lhhnulg5ze3jwgx4slrfx57esp 50udaric --from $ALICE_KEY $TX_FLAGS

# Try creating a channel
echo "Creating a channel"
desmos tx ibc channel open-init transfer transfer connection-1 --ordered=false --from node0 $TX_FLAGS

#echo "Open channel completed"
#desmos tx ibc-transfer transfer transfer channel-0 desmos1zfuyr8jd65d5lhhnulg5ze3jwgx4slrfx57esp 100udaric $TX_FLAGS
#echo "Transfer transaction successful"
