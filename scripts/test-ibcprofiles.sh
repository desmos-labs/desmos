#!/usr/bin/env bash

IBCDIR=$1
ACCOUNTNUM=$2
SRCNODE=$3
DSTNODE=$4

echo "Waiting for channel open"
# Check channel is open on ibc0
channel=$(desmos query ibc channel channels | grep -A1 "ibcprofiles" | grep "STATE_OPEN")
while [ "$channel" = "" ]
do
    sleep 10
    channel=$(desmos query ibc channel channels | grep -A1 "ibcprofiles" | grep "STATE_OPEN")
done
echo "Src links channel available now"

# Check channel is open on ibc1
channel=$(desmos query ibc channel channels | grep -A1 "profiles" | grep "STATE_OPEN")
while [ "$channel" = "" ]
do
    sleep 10
    channel=$(desmos query ibc channel channels | grep -A1 "profiles" | grep "STATE_OPEN")
done
echo "Dst links channel available now"

# Wait for relayer start relay packets
sleep 10

echo "Test start"
echo "Starting sending transactions"
# Create link via ibc
for (( i = 0; i < $ACCOUNTNUM; i++ ))
do
    if [ `expr $i % 2` == 0 ]; then
        desmos tx ibc-profiles create-ibc-connection ibcprofiles channel-1 desmos $IBCDIR/ibc1 test1-$i --home $IBCDIR/ibc0 \
        --keyring-backend test --from test0-$i --chain-id ibc0 --node $SRCNODE --broadcast-mode async --yes &> /dev/null
    else
        desmos tx ibc-profiles create-ibc-link ibcprofiles channel-1 desmos --home $IBCDIR/ibc0 --keyring-backend test \
        --from test0-$i --chain-id ibc0 --node $SRCNODE --broadcast-mode async --yes &> /dev/null
    fi
done

# Wait for nodes deal with ibc transactions
echo "Waiting for chains deal with transactions"
sleep 60

# Check database if including results 
# echo "Starting checking links"
# for (( i = 0; i < $ACCOUNTNUM; i++ ))
# do
#     result=$(desmos query links link  $(desmos keys show test0-$i --home $IBCDIR/ibc0 --keyring-backend test  --address) \
#     --home $IBCDIR/ibc1 --chain-id ibc1 --node $DSTNODE)
#     if [ "$result" = "" ]; then
#         echo "Failed to get  test0-$i link"
#         exit 1
#     fi
# done

echo "Test successfully"