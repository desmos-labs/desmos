package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/reactions/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v7/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// RegisterInvariants registers all reactions invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-subspaces",
		ValidSubspacesInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-registered-reactions",
		ValidRegisteredReactionsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-reactions",
		ValidReactionsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-reactions-params",
		ValidReactionsParamsInvariant(keeper))
}

// --------------------------------------------------------------------------------------------------------------------

// ValidSubspacesInvariant checks that all the subspaces have a valid registered reaction id and reaction params
func ValidSubspacesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (message string, broken bool) {
		var invalidSubspaces []subspacestypes.Subspace
		k.sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
			invalid := false

			// Make sure the next registered reaction id exists
			if !k.HasNextRegisteredReactionID(ctx, subspace.ID) {
				invalid = true
			}

			// Make sure the reactions params exist
			if !k.HasSubspaceReactionsParams(ctx, subspace.ID) {
				invalid = true
			}

			if invalid {
				invalidSubspaces = append(invalidSubspaces, subspace)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid subspaces",
			fmt.Sprintf("the following subspaces are invalid:\n%s", subspaceskeeper.FormatOutputSubspaces(invalidSubspaces)),
		), invalidSubspaces != nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

// ValidRegisteredReactionsInvariant checks that all the stored registered reactions are valid
func ValidRegisteredReactionsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidReactions []types.RegisteredReaction
		k.IterateRegisteredReactions(ctx, func(reaction types.RegisteredReaction) (stop bool) {
			invalid := false

			// Make sure the subspace exists
			if !k.HasSubspace(ctx, reaction.SubspaceID) {
				invalid = true
			}

			nextRegisteredReactionID, err := k.GetNextRegisteredReactionID(ctx, reaction.SubspaceID)
			if err != nil {
				invalid = true
			}

			// Make sure the registered reaction id is always less than the next one
			if reaction.ID >= nextRegisteredReactionID {
				invalid = true
			}

			// Validate the reaction
			err = reaction.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidReactions = append(invalidReactions, reaction)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid registered reactions",
			fmt.Sprintf("the following registered reactions are invalid:\n%s", formatOutputRegisteredReactions(invalidReactions)),
		), invalidReactions != nil
	}
}

// formatOutputRegisteredReactions concatenates the given registered reactions information into a string
func formatOutputRegisteredReactions(reactions []types.RegisteredReaction) (output string) {
	for _, reaction := range reactions {
		output += fmt.Sprintf("SuspaceID: %d, RegisteredReactionID: %d\n", reaction.SubspaceID, reaction.ID)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidReactionsInvariant checks that all the stored reactions are valid
func ValidReactionsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidReactions []types.Reaction
		k.IterateReactions(ctx, func(reaction types.Reaction) (stop bool) {
			invalid := false

			// Make sure the subspace exists
			if !k.HasSubspace(ctx, reaction.SubspaceID) {
				invalid = true
			}

			// Make sure the post exists
			if !k.HasPost(ctx, reaction.SubspaceID, reaction.PostID) {
				invalid = true
			}

			nextReactionID, err := k.GetNextReactionID(ctx, reaction.SubspaceID, reaction.PostID)
			if err != nil {
				invalid = true
			}

			// Make sure the reaction id is always less than the next one
			if reaction.ID >= nextReactionID {
				invalid = true
			}

			// Validate the reaction
			err = reaction.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidReactions = append(invalidReactions, reaction)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid reactions",
			fmt.Sprintf("the following reactions are invalid:\n%s", formatOutputReactions(invalidReactions)),
		), invalidReactions != nil
	}
}

// formatOutputReactions concatenates the given reactions information into a string
func formatOutputReactions(reactions []types.Reaction) (output string) {
	for _, reaction := range reactions {
		output += fmt.Sprintf("SuspaceID: %d, PostID: %d, ReactionID: %d\n",
			reaction.SubspaceID, reaction.PostID, reaction.ID)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidReactionsParamsInvariant checks that all the stored reactions params are valid
func ValidReactionsParamsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidParams []types.SubspaceReactionsParams
		k.IterateReactionsParams(ctx, func(params types.SubspaceReactionsParams) (stop bool) {
			invalid := false

			// Make sure the subspace exists
			if !k.HasSubspace(ctx, params.SubspaceID) {
				invalid = true
			}

			// Validate the params
			err := params.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidParams = append(invalidParams, params)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid reactions params",
			fmt.Sprintf("the following reactions params are invalid:\n%s", formatOutputParams(invalidParams)),
		), invalidParams != nil
	}
}

// formatOutputParams concatenates the given reactions params information into a string
func formatOutputParams(reasons []types.SubspaceReactionsParams) (output string) {
	for _, reason := range reasons {
		output += fmt.Sprintf("SuspaceID: %d\n", reason.SubspaceID)
	}
	return output
}
