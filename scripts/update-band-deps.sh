#!/usr/bin/env bash

ORACLE_TYPES_FOLDER=./x/oracle/types/
mkdir -p $ORACLE_TYPES_FOLDER
curl "https://raw.githubusercontent.com/bandprotocol/bandchain-packet/master/packet/packet.go" > $ORACLE_TYPES_FOLDER/packet.go
curl "https://raw.githubusercontent.com/bandprotocol/bandchain-packet/master/packet/packet.pb.go" > $ORACLE_TYPES_FOLDER/packet.pb.go

OBI_FOLDER=./pkg/obi
mkdir -p $OBI_FOLDER
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.0.0-rc2/pkg/obi/schema.go" > $OBI_FOLDER/schema.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.0.0-rc2/pkg/obi/encode.go" > $OBI_FOLDER/encode.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.0.0-rc2/pkg/obi/decode.go" > $OBI_FOLDER/decode.go
