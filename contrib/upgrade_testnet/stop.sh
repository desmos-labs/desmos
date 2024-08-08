#!/usr/bin/env bash

TESTNETDIR=$(pwd)/contrib/upgrade_testnet
docker compose -f $TESTNETDIR/docker-compose.yml down