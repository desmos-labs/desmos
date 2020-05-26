module github.com/desmos-labs/desmos

go 1.13

require (
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/desmos-labs/Go-Emoji-Utils v1.1.1-0.20200515063516-9c493b11de3e
	github.com/gorilla/mux v1.7.3
	github.com/otiai10/copy v1.0.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.3
	github.com/tendermint/tm-db v0.5.0
)

replace github.com/cosmos/cosmos-sdk => github.com/RiccardoM/cosmos-sdk v0.34.4-0.20200514074150-9bee4bc6ad6a
