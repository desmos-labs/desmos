package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	wasmdesmos "github.com/desmos-labs/desmos/v3/cosmwasm"
	postskeeper "github.com/desmos-labs/desmos/v3/x/posts/keeper"
	postswasm "github.com/desmos-labs/desmos/v3/x/posts/wasm"
	profileskeeper "github.com/desmos-labs/desmos/v3/x/profiles/keeper"
	profileswasm "github.com/desmos-labs/desmos/v3/x/profiles/wasm"
	relationshipskeeper "github.com/desmos-labs/desmos/v3/x/relationships/keeper"
	relationshipswasm "github.com/desmos-labs/desmos/v3/x/relationships/wasm"
	reportsKeeper "github.com/desmos-labs/desmos/v3/x/reports/keeper"
	reportswasm "github.com/desmos-labs/desmos/v3/x/reports/wasm"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	subspaceswasm "github.com/desmos-labs/desmos/v3/x/subspaces/wasm"
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

// NewDesmosCustomQueryPlugin initialize the custom querier to handle desmos queries for contracts
func NewDesmosCustomQueryPlugin(
	cdc codec.Codec,
	profilesKeeper profileskeeper.Keeper,
	subspacesKeeper subspaceskeeper.Keeper,
	relationshipsKeeper relationshipskeeper.Keeper,
	postsKeeper postskeeper.Keeper,
	reportsKeeper reportsKeeper.Keeper,
) wasm.QueryPlugins {
	queriers := map[string]wasmdesmos.Querier{
		wasmdesmos.QueryRouteProfiles:      profileswasm.NewProfilesWasmQuerier(profilesKeeper, cdc),
		wasmdesmos.QueryRouteSubspaces:     subspaceswasm.NewSubspacesWasmQuerier(subspacesKeeper, cdc),
		wasmdesmos.QueryRouteRelationships: relationshipswasm.NewRelationshipsWasmQuerier(relationshipsKeeper, cdc),
		wasmdesmos.QueryRoutePosts:         postswasm.NewPostsWasmQuerier(postsKeeper, cdc),
		wasmdesmos.QueryRouteReports:       reportswasm.NewReportsWasmQuerier(reportsKeeper, cdc),
		// add other modules querier hereo
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
		wasmdesmos.WasmMsgParserRouteProfiles:      profileswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteSubspaces:     subspaceswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteRelationships: relationshipswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRoutePosts:         postswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteReports:       reportswasm.NewWasmMsgParser(cdc),
		// add other modules parsers here
	}

	parserRouter.Parsers = parsers
	return wasm.MessageEncoders{
		Custom: parserRouter.ParseCustom,
	}
}
