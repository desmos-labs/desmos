#!/usr/bin/env bash

HOMEDIR=$1
GENFILE=$2
ACCOUNTNUM=10

# Ensure desmos is installed
if ! [ -x "$(which desmos)" ]; then
  echo "Error: desmos is not installed. Try running 'make install'" >&2
  exit 1
fi


if [ -d $HOMEDIR ]; then
    rm -r $HOMEDIR/*
fi

echo "Generate ibc0 testnet config"

# generate ibc0 chain config
bash $GENFILE 0 $ACCOUNTNUM $HOMEDIR --y &> ibc0-account.out

echo "Generate ibc1 testnet config"

# generate ibc1 chain config
bash $GENFILE 1 $ACCOUNTNUM $HOMEDIR --y &> ibc1-account.out
