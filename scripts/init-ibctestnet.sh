#!/usr/bin/env bash

IBCDIR=$1
ACCOUNTNUM=$2
ibc0=$3
ibc1=$4

# Init account tx in ibc
for (( i = 0; i < $ACCOUNTNUM; i++ ))
do
    # Init ibc0 accounts
    desmos tx bank send test0-$i desmos1punhxfyxvnwup70mc6cz9cmuxu74ars7x6hgjm 1desmos --home "$IBCDIR/ibc0" \
    --keyring-backend "test" --node $ibc0 --chain-id ibc0 --broadcast-mode async --yes &>/dev/null

    # Init ibc1 account by creating profile
    desmos tx profiles save test1_$i --home "$IBCDIR/ibc1" --from test1-$i \
    --keyring-backend "test" --node $ibc1 --chain-id ibc1 --broadcast-mode async --yes &>/dev/null

done

# Init ibc0 accounts which will send ibc account link packet to ibc1 on ibc1 chain
for (( i = 0; i < $ACCOUNTNUM; i++ ))
do 
    if [ `expr $i % 2` == 1 ]; then
        desmos tx bank send test1-$i $(desmos keys show test0-$i --home $IBCDIR/ibc0 --keyring-backend "test" --address) 1000desmos --home "$IBCDIR/ibc1" \
        --keyring-backend "test" --node $ibc1 --chain-id ibc1 --broadcast-mode async --yes &>/dev/null
    fi
done

sleep 15

# Create ibc0 profile on ibc1 chain
for (( i = 0; i < $ACCOUNTNUM; i++ ))
do 
    if [ `expr $i % 2` == 1 ]; then
        desmos tx profiles save test0_$i --home "$IBCDIR/ibc0" --from test0-$i \
        --keyring-backend "test" --node $ibc1 --chain-id ibc1 --broadcast-mode async --yes &>/dev/null
    fi
done