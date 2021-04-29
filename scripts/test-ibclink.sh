#!/usr/bin/env bash

IBCDIR=$1
ACCOUNTNUM=$2
NODE=$3

for (( i = 0; i < $ACCOUNTNUM; i++ ))
do
    if [ `expr $i % 2` == 0 ]; then
        desmos tx links create-ibc-connection links channel-1 desmos $IBCDIR/ibc1 test1-$i --home $IBCDIR/ibc0 --keyring-backend test --from test0-$i --chain-id ibc0 --node $NODE --yes &> ibclink-test.out
    else
        desmos tx links create-ibc-link links channel-1 desmos --home $IBCDIR/ibc0 --keyring-backend test --from test0-$i --chain-id ibc0 --node $NODE --yes &> ibclink-test.out
    fi
done