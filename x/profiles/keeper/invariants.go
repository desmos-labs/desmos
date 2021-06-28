package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-profiles", ValidProfilesInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-blocks", ValidUserBlocksInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-relationships", ValidRelationshipsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-dtag-transfer-requests", ValidDTagTransferRequests(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-chain-links", ValidChainLinks(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-application-links", ValidApplicationLinks(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := ValidProfilesInvariant(k)(ctx); stop {
			return res, stop
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

// ValidProfilesInvariant checks that all registered Profiles have a non-empty DTag and a non-empty creator
func ValidProfilesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidProfiles []*types.Profile
		k.IterateProfiles(ctx, func(_ int64, profile *types.Profile) (stop bool) {
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

// formatOutputProfiles prepare invalid Profiles to be displayed correctly
func formatOutputProfiles(invalidProfiles []*types.Profile) (outputProfiles string) {
	outputProfiles = "Invalid profiles:\n"
	for _, invalidProfile := range invalidProfiles {
		outputProfiles += fmt.Sprintf(
			"[DTag]: %s, [Creator]: %s\n",
			invalidProfile.DTag, invalidProfile.GetAddress().String(),
		)
	}
	return outputProfiles
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserBlocksInvariant checks that all created user blocks have been created by a user with a profile
// and they do not have the same user as creator and recipient
func ValidUserBlocksInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidBlocks []types.UserBlock
		k.IterateBlocks(ctx, func(index int64, block types.UserBlock) (stop bool) {
			if !k.HasProfile(ctx, block.Blocker) || block.Blocker == block.Blocked {
				invalidBlocks = append(invalidBlocks, block)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user blocks",
			fmt.Sprintf("The following list contains invalid user blocks:\n %s",
				formatOutputBlocks(invalidBlocks)),
		), invalidBlocks != nil
	}
}

// formatOutputProfiles prepares the given invalid user blocks to be displayed correctly
func formatOutputBlocks(invalidBlocks []types.UserBlock) (outputBlocks string) {
	outputBlocks = "Invalid users blocks:\n"
	for _, block := range invalidBlocks {
		outputBlocks += fmt.Sprintf(
			"[Blocker]: %s, [Blocked]: %s, [Subspace]: %s\n",
			block.Blocker, block.Blocked, block.Subspace,
		)
	}
	return outputBlocks
}

// --------------------------------------------------------------------------------------------------------------------

// ValidRelationshipsInvariant checks that all relationships are associated with a creator that has a profile
// and they do not have the same user as creator and recipient
func ValidRelationshipsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidRelationships []types.Relationship
		k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
			if !k.HasProfile(ctx, relationship.Creator) || relationship.Creator == relationship.Recipient {
				invalidRelationships = append(invalidRelationships, relationship)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid relationships",
			fmt.Sprintf("The following list contains invalid relationships:\n %s",
				formatOutputRelationships(invalidRelationships)),
		), invalidRelationships != nil
	}
}

// formatOutputRelationships prepares the given invalid relationships to be displayed correctly
func formatOutputRelationships(relationships []types.Relationship) (output string) {
	output = "Invalid relationships:\n"
	for _, relationship := range relationships {
		output += fmt.Sprintf(
			"[Creator]: %s, [Recipient]: %s, [Subspace]: %s\n",
			relationship.Creator, relationship.Recipient, relationship.Subspace,
		)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidDTagTransferRequests checks that all DTag transfer requests are associated with a recipient that has a profile
// and they have not been made from the same user towards the same user
func ValidDTagTransferRequests(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidDTagTransferRequests []types.DTagTransferRequest
		k.IterateDTagTransferRequests(ctx, func(index int64, request types.DTagTransferRequest) (stop bool) {
			if !k.HasProfile(ctx, request.Receiver) || request.Sender == request.Receiver {
				invalidDTagTransferRequests = append(invalidDTagTransferRequests, request)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid dtag transfer requests",
			fmt.Sprintf("The following list contains invalid DTag transfer requests:\n %s",
				formatOutputDTagTransferRequests(invalidDTagTransferRequests)),
		), invalidDTagTransferRequests != nil
	}
}

// formatOutputDTagTransferRequests prepares the given invalid DTag transfer requests to be displayed correctly
func formatOutputDTagTransferRequests(requests []types.DTagTransferRequest) (output string) {
	output = "Invalid DTag transfer requests:\n"
	for _, request := range requests {
		output += fmt.Sprintf(
			"[Sender]: %s, [Receiver]: %s\n",
			request.Sender, request.Receiver,
		)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidChainLinks checks that all chain links are associated with a user that has a profile
func ValidChainLinks(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidChainLinks []types.ChainLink
		k.IterateChainLinks(ctx, func(index int64, link types.ChainLink) (stop bool) {
			if !k.HasProfile(ctx, link.User) {
				invalidChainLinks = append(invalidChainLinks, link)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid chain links",
			fmt.Sprintf("The following list contains invalid chain links:\n %s",
				formatOutputChainLinks(invalidChainLinks)),
		), invalidChainLinks != nil
	}
}

// formatOutputChainLinks prepares the given invalid chain links to be displayed correctly
func formatOutputChainLinks(links []types.ChainLink) (output string) {
	output = "Invalid chain links:\n"
	for _, link := range links {
		address := link.Address.GetCachedValue().(types.AddressData)
		output += fmt.Sprintf(
			"[User]: %s, [Chain]: %s, [Address]: %s\n",
			link.User, link.ChainConfig.Name, address.GetAddress(),
		)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidApplicationLinks checks that all application links are associated with a user that has a profile
func ValidApplicationLinks(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidApplicationLinks []types.ApplicationLink
		k.IterateApplicationLinks(ctx, func(index int64, link types.ApplicationLink) (stop bool) {
			if !k.HasProfile(ctx, link.User) {
				invalidApplicationLinks = append(invalidApplicationLinks, link)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid application links",
			fmt.Sprintf("The following list contains invalid application links:\n %s",
				formatOutputApplicationLinks(invalidApplicationLinks)),
		), invalidApplicationLinks != nil
	}
}

// formatOutputApplicationLinks prepares the given invalid application links to be displayed correctly
func formatOutputApplicationLinks(links []types.ApplicationLink) (output string) {
	output = "Invalid application links:\n"
	for _, link := range links {
		output += fmt.Sprintf(
			"[User]: %s, [Application]: %s, [Username]: %s\n",
			link.User, link.Data.Application, link.Data.Username,
		)
	}
	return output
}
