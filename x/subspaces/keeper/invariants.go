package keeper

import (
	"fmt"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants registers all subspaces invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-subspaces",
		ValidSubspacesInvariant(keeper))
}

// AllInvariants runs all invariants of the module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, broken := ValidSubspacesInvariant(k)(ctx)
		if broken {
			return res, true
		}

		return "Every invariant condition is fulfilled correctly", false
	}
}

// formatOutputSubspaces concatenate the subspaces given into a unique string
func formatOutputSubspaces(subspaces []types.Subspace) (outputSubspaces string) {
	for _, subspace := range subspaces {
		outputSubspaces += fmt.Sprintf("%d\n", subspace.ID)
	}
	return outputSubspaces
}

// ValidSubspacesInvariant checks that all the subspaces are valid
func ValidSubspacesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidSubspaces []types.Subspace
		k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
			err := subspace.Validate()
			if err != nil {
				invalidSubspaces = append(invalidSubspaces, subspace)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid subspaces",
			fmt.Sprintf("the following subspaces are invalid:\n %s", formatOutputSubspaces(invalidSubspaces)),
		), invalidSubspaces != nil
	}
}
