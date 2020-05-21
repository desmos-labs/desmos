package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
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

// ValidReportsIDs checks that all reports are associated with a valid postID that correspond to an existent post
func ValidReportsIDs(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {

	}
}
