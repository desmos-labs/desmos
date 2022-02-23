package keeper

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

// WithWasmKeeper decorates profiles keeper with the cosmwasm keeper
func (k Keeper) WithWasmKeeper(wasmKeeper wasmkeeper.Keeper) Keeper {
	k.wasmKeeper = wasmKeeper
	return k
}
