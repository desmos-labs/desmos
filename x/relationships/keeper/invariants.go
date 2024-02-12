package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-user-blocks", ValidUserBlocksInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-relationships", ValidRelationshipsInvariant(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, broken := ValidUserBlocksInvariant(k)(ctx)
		if broken {
			return res, broken
		}

		res, broken = ValidRelationshipsInvariant(k)(ctx)
		if broken {
			return res, broken
		}

		return "Every invariant condition is fulfilled correctly", false
	}
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserBlocksInvariant checks that all created user blocks have been created by a user with a profile
// and they do not have the same user as creator and recipient
func ValidUserBlocksInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidBlocks []types.UserBlock
		k.IterateUsersBlocks(ctx, func(index int64, block types.UserBlock) (stop bool) {
			if block.Blocker == block.Blocked {
				invalidBlocks = append(invalidBlocks, block)
			}
			return false
		})

		broken := len(invalidBlocks) != 0
		return sdk.FormatInvariant(types.ModuleName, "invalid user blocks",
			formatOutputBlocks(invalidBlocks)), broken
	}
}

// formatOutputProfiles prepares the given invalid user blocks to be displayed correctly
func formatOutputBlocks(invalidBlocks []types.UserBlock) (outputBlocks string) {
	outputBlocks = "The following list contains invalid user blocks:\n"
	for _, block := range invalidBlocks {
		outputBlocks += fmt.Sprintf(
			"[Blocker]: %s, [Blocked]: %s, [SubspaceID]: %d\n",
			block.Blocker, block.Blocked, block.SubspaceID,
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
			if relationship.Creator == relationship.Counterparty {
				invalidRelationships = append(invalidRelationships, relationship)
			}
			return false
		})

		broken := len(invalidRelationships) != 0
		return sdk.FormatInvariant(types.ModuleName, "invalid relationships",
			formatOutputRelationships(invalidRelationships)), broken
	}
}

// formatOutputRelationships prepares the given invalid relationships to be displayed correctly
func formatOutputRelationships(relationships []types.Relationship) (output string) {
	output = "The following list contains invalid relationships:\n"
	for _, relationship := range relationships {
		output += fmt.Sprintf(
			"[Creator]: %s, [Counterparty]: %s, [SubspaceID]: %d\n",
			relationship.Creator, relationship.Counterparty, relationship.SubspaceID,
		)
	}
	return output
}
