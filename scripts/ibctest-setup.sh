#!/usr/bin/env bash

WORKDIR="build/ibc"
GENFILE="scripts/ibctestchain-gen.sh"

if [ -d $WORKDIR ]; then
    rm -r $WORKDIR
fi

bash $GENFILE --y