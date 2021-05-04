#!/usr/bin/env bash

# Wait chains start
sleep 10

# Create LCD
ibc-setup connect

# Set channels
ibc-setup channel --src=ibc0 --dest=ibc1 --src-port=transfer --dest-port=transfer --version=ics20-1
ibc-setup channel --src=ibc0 --dest=ibc1 --src-port=links --dest-port=links --version=links-1

ibc-relayer start --log-level=debug --poll 5