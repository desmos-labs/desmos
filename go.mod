module github.com/desmos-labs/desmos

go 1.15

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/cosmos/cosmos-sdk v0.40.0-rc5
	github.com/desmos-labs/Go-Emoji-Utils v1.1.1-0.20200515063516-9c493b11de3e
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0
	github.com/tendermint/tm-db v0.6.3
	google.golang.org/genproto v0.0.0-20201111145450-ac7456db90a6
	google.golang.org/grpc v1.33.2
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
