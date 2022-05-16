package keeper

import (
	"fmt"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants registers all subspaces invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	// TODO Add better cheks
	// - Next section id
	// - next group id
	// - has section
	// - has group
	// - section id <= next section id
	// - group id <= next group id
	ir.RegisterRoute(types.ModuleName, "valid-subspaces",
		ValidSubspacesInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-groups",
		ValidUserGroupsInvariant(keeper))
}

// AllInvariants runs all invariants of the module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, broken := ValidSubspacesInvariant(k)(ctx)
		if broken {
			return res, true
		}

		res, broken = ValidUserGroupsInvariant(k)(ctx)
		if broken {
			return res, true
		}

		return "Every invariant condition is fulfilled correctly", false
	}
}

// --------------------------------------------------------------------------------------------------------------------

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

// formatOutputSubspaces concatenate the subspaces given into a unique string
func formatOutputSubspaces(subspaces []types.Subspace) (outputSubspaces string) {
	for _, subspace := range subspaces {
		outputSubspaces += fmt.Sprintf("%d\n", subspace.ID)
	}
	return outputSubspaces
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserGroupsInvariant checks that all the subspaces are valid
func ValidUserGroupsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidUserGroups []types.UserGroup
		k.IterateUserGroups(ctx, func(_ int64, group types.UserGroup) (stop bool) {
			err := group.Validate()
			if err != nil {
				// The group is not valid
				invalidUserGroups = append(invalidUserGroups, group)
			}

			if !k.HasSubspace(ctx, group.SubspaceID) {
				// The subspace for this group does not exist anymore
				invalidUserGroups = append(invalidUserGroups, group)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user groups",
			fmt.Sprintf("the following user groups are invalid:\n %s", formatOutputUserGroups(invalidUserGroups)),
		), invalidUserGroups != nil
	}
}

// formatOutputUserGroups concatenate the subspaces given into a unique string
func formatOutputUserGroups(groups []types.UserGroup) (outputUserGroups string) {
	for _, group := range groups {
		outputUserGroups += fmt.Sprintf("%d\n", group.ID)
	}
	return outputUserGroups
}
