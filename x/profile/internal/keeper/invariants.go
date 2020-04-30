package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
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
func formatOutputProfiles(invalidProfiles types.Profiles) (outputProfiles string) {
	outputProfiles = "Invalid profiles:\n"
	for _, invalidProfile := range invalidProfiles {
		outputProfiles += fmt.Sprintf("[Moniker]: %s, [Creator]: %s\n", invalidProfile.Moniker, invalidProfile.Creator)
	}
	return outputProfiles
}

// ValidProfileInvariant checks that all registered profiles have a non-empty moniker and a non-empty creator
func ValidProfileInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidProfiles types.Profiles
		k.IterateProfiles(ctx, func(_ int64, profile types.Profile) (stop bool) {
			if err := profile.Validate(); err != nil {
				invalidProfiles = append(invalidProfiles, profile)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid profiles",
			fmt.Sprintf("The following list contains invalid profiles that have empty moniker, empty creator or both:\n %s",
				formatOutputProfiles(invalidProfiles)),
		), invalidProfiles != nil
	}
}
