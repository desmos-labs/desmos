#!/usr/bin/env bash

WORKDIR="build/ibc"
GENFILE="scripts/ibctestchain-gen.sh"

# Ensure desmos is installed
if ! [ -x "$(which desmos)" ]; then
  echo "Error: desmos is not installed. Try running 'make install'" >&2
  exit 1
fi

# Ensure relayer is installed
if ! [ -x "$(which rly)" ]; then
  echo "Error: relayer is not installed. Try running 'make install-relayer'" >&2
  exit 1
fi

if [ -d $WORKDIR ]; then
    rm -r $WORKDIR
fi

# generate ibc0 chain config
bash $GENFILE 0 1 $WORKDIR --y

# generate ibc1 chain config
bash $GENFILE 1 1 $WORKDIR --y
