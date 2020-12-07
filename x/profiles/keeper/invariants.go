package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-profile",
		ValidProfileInvariant(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := ValidProfileInvariant(k)(ctx); stop {
			return res, stop
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

// formatOutputProfiles prepare invalid profiles to be displayed correctly
func formatOutputProfiles(invalidProfiles []types.Profile) (outputProfiles string) {
	outputProfiles = "Invalid profiles:\n"
	for _, invalidProfile := range invalidProfiles {
		outputProfiles += fmt.Sprintf(
			"[DTag]: %s, [Creator]: %s\n",
			invalidProfile.Dtag,
			invalidProfile.Creator,
		)
	}
	return outputProfiles
}

// ValidProfileInvariant checks that all registered profiles have a non-empty dtag and a non-empty creator
func ValidProfileInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidProfiles []types.Profile
		k.IterateProfiles(ctx, func(_ int64, profile types.Profile) (stop bool) {
			if err := profile.Validate(); err != nil {
				invalidProfiles = append(invalidProfiles, profile)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid profiles",
			fmt.Sprintf("The following list contains invalid profiles:\n %s",
				formatOutputProfiles(invalidProfiles)),
		), invalidProfiles != nil
	}
}
