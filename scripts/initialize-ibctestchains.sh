#!/usr/bin/env bash

HOMEDIR=$1
GENFILE=$2

# Ensure desmos is installed
if ! [ -x "$(which desmos)" ]; then
  echo "Error: desmos is not installed. Try running 'make install'" >&2
  exit 1
fi


if [ -d $HOMEDIR ]; then
    rm -r $HOMEDIR/*
fi

# generate ibc0 chain config
bash $GENFILE 0 1 $HOMEDIR --y

# generate ibc1 chain config
bash $GENFILE 1 1 $HOMEDIR --y
