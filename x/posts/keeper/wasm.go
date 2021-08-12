package keeper

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	subspacestypes "github.com/desmos-labs/desmos/x/subspaces/types"
)

// WithWasmKeeper decorate the posts keeper with Wasm keeper
func (k Keeper) WithWasmKeeper(wasmKeeper wasm.Keeper) Keeper {
	k.wk = wasmKeeper
	return k
}

// ExecuteTokenomics perform the tokenomics for each subspace in the given context
func (k Keeper) ExecuteTokenomics(ctx sdk.Context) {
	k.IterateSubspacesTokenomics(ctx, func(index int64, tokenomics subspacestypes.Tokenomics) (stop bool) {
		contractAddr, _ := sdk.AccAddressFromBech32(tokenomics.ContractAddress)

		_, err := k.wk.Sudo(ctx, contractAddr, tokenomics.Message)

		k.Logger(ctx).Info("tokenomics executed", "subspace",
			tokenomics.SubspaceID, "contractAddress", tokenomics.ContractAddress)

		if err != nil {
			k.Logger(ctx).Error("ERROR", err)
			fmt.Println("[!] error: ", err.Error())
		}

		return false
	})
}
