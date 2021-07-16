#!/usr/bin/env bash

NODES=$1
UPGRADE_NAME=$2

# Copy the Desmos binary into the proper folder
UPGRADE_FOLDER="build/node0/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin"
if [ ! -d "$UPGRADE_FOLDER" ]; then
  echo "===> Setting up upgrade binary"

  for ((i = 0; i < $NODES; i++)); do
    echo "====> Node $i"
    mkdir -p "build/node$i/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin"
    cp $(which desmos) "build/node$i/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin/desmos"
  done
fi
