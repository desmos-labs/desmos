package keeper

import (
	"bytes"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// RegisterInvariants registers all reports invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-stored-ids",
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
func formatOutputIDs(ids []string) (outputIDs string) {
	return strings.Join(ids, "\n")
}

// ValidReportsIDs checks that all reports are associated with a valid post id that corresponds to an existent post
func ValidReportsIDs(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidIDs []string
		store := ctx.KVStore(k.storeKey)
		iterator := sdk.KVStorePrefixIterator(store, types.ReportsStorePrefix)
		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			postID := string(bytes.TrimPrefix(iterator.Key(), types.ReportsStorePrefix))
			if !poststypes.IsValidPostID(postID) {
				invalidIDs = append(invalidIDs, postID)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "invalid reports",
			fmt.Sprintf("The following list contains invalid postIDs:\n %s",
				formatOutputIDs(invalidIDs))), invalidIDs != nil
	}
}
