package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	wasmdesmos "github.com/desmos-labs/desmos/v2/cosmwasm"
	profileskeeper "github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	profileswasm "github.com/desmos-labs/desmos/v2/x/profiles/wasm"
)

const (
	// DefaultDesmosInstanceCost is how much SDK gas we charge each time we load a WASM instance
	DefaultDesmosInstanceCost uint64 = 60_000
	// DefaultDesmosCompileCost is how much SDK gas is charged *per byte* for compiling WASM code
	DefaultDesmosCompileCost uint64 = 2
)

// DesmosWasmGasRegister is defaults plus a custom compile amount
func DesmosWasmGasRegister() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultDesmosInstanceCost
	gasConfig.CompileCost = DefaultDesmosCompileCost

	return gasConfig
}

func NewDesmosWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(DesmosWasmGasRegister())
}

// NewDesmosCustomQueryPlugin initialize the custom queries to desmos app for contracts
func NewDesmosCustomQueryPlugin(cdc codec.Codec, profilesKeeper profileskeeper.Keeper) wasm.QueryPlugins {
	queriers := map[string]wasmdesmos.Querier{
		wasmdesmos.QueryRouteProfiles: profileswasm.NewProfilesWasmQuerier(profilesKeeper, cdc),
	}

	querier := wasmdesmos.NewQuerier(queriers)

	return wasm.QueryPlugins{
		Custom: querier.QueryCustom,
	}
}

// NewDesmosCustomMessageEncoder initialize the custom message encoder to desmos app for contracts
func NewDesmosCustomMessageEncoder(cdc codec.Codec) wasm.MessageEncoders {
	// Initialization of custom Desmos messages for contracts
	parserRouter := wasmdesmos.NewParserRouter()
	parsers := map[string]wasmdesmos.MsgParserInterface{
		wasmdesmos.WasmMsgParserRouteProfiles: profileswasm.NewWasmMsgParser(cdc),
		// add other modules parsers here
	}

	parserRouter.Parsers = parsers
	return wasm.MessageEncoders{
		Custom: parserRouter.ParseCustom,
	}
}
