#!/usr/bin/env bash

ORACLE_TYPES_FOLDER=./x/oracle/types/
mkdir -p $ORACLE_TYPES_FOLDER
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/id.go" > $ORACLE_TYPES_FOLDER/id.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/codec.go" > $ORACLE_TYPES_FOLDER/codec.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/error.go" > $ORACLE_TYPES_FOLDER/error.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/keys.go" > $ORACLE_TYPES_FOLDER/keys.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/packets.go" > $ORACLE_TYPES_FOLDER/packets.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/params.go" > $ORACLE_TYPES_FOLDER/params.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/x/oracle/types/oracle.pb.go" > $ORACLE_TYPES_FOLDER/oracle.pb.go

OBI_FOLDER=./pkg/obi
mkdir -p $OBI_FOLDER
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/pkg/obi/schema.go" > $OBI_FOLDER/schema.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/pkg/obi/encode.go" > $OBI_FOLDER/encode.go
curl "https://raw.githubusercontent.com/bandprotocol/chain/v2.3.1/pkg/obi/decode.go" > $OBI_FOLDER/decode.go
