package keeper

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-reports-ids",
		ValidReportsIDs(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := ValidReportsIDs(k)(ctx); stop {
			return res, stop
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

// formatOutputIDs concatenate the ids given into a unique string
func formatOutputIDs(ids posts.PostIDs) (outputIDs string) {
	for _, id := range ids {
		outputIDs += id.String() + "\n"
	}
	return outputIDs
}

// ValidReportsIDs checks that all reports are associated with a valid postID that correspond to an existent post
func ValidReportsIDs(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidIDs posts.PostIDs
		store := ctx.KVStore(k.StoreKey)
		iterator := sdk.KVStorePrefixIterator(store, types.ReportsStorePrefix)
		defer iterator.Close()
		for ; iterator.Valid(); iterator.Next() {
			postID := posts.PostID(bytes.TrimPrefix(iterator.Key(), types.ReportsStorePrefix))
			if valid := postID.Valid(); !valid {
				invalidIDs = append(invalidIDs, postID)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "invalid reports' IDs",
			fmt.Sprintf("The following list contains invalid postIDs:\n %s",
				formatOutputIDs(invalidIDs))), invalidIDs != nil
	}
}
