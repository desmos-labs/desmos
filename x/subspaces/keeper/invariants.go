package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

// RegisterInvariants registers all subspaces invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-subspaces",
		ValidSubspacesInvariant(keeper))
}

// AllInvariants runs all invariants of the module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := ValidSubspacesInvariant(k)(ctx); stop {
			return res, stop
		}
		return "Every invariant condition is fulfilled correctly", true
	}
}

// formatOutputSubspaces concatenate the subspaces given into a unique string
func formatOutputSubspaces(subspaces []types.Subspace) (outputSubspaces string) {
	for _, subspace := range subspaces {
		outputSubspaces += subspace.ID + "\n"
	}
	return outputSubspaces
}

// ValidSubspacesInvariant checks that all the subspaces are valid
func ValidSubspacesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidSubspaces []types.Subspace
		k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
			if err := subspace.Validate(); err != nil {
				invalidSubspaces = append(invalidSubspaces, subspace)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid subspaces",
			fmt.Sprintf("the following subspaces are invalid:\n %s", formatOutputSubspaces(invalidSubspaces)),
		), invalidSubspaces != nil
	}
}
