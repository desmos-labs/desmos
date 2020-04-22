package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "hash256-post-id",
		Hash256PostIDInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "comments-date",
		ValidCommentsDateInvariants(keeper))
	ir.RegisterRoute(types.ModuleName, "post-reactions",
		ValidPostForReactionsInvariants(keeper))
	ir.RegisterRoute(types.ModuleName, "post-poll-answers",
		ValidPollForPollAnswersInvariants(keeper))
}

//TODO not sure if this should be used or can be deleted

// AllInvariants runs all invariants of the module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := Hash256PostIDInvariant(k)(ctx); stop {
			return res, stop
		}

		if res, stop := ValidCommentsDateInvariants(k)(ctx); stop {
			return res, stop
		}

		if res, stop := ValidPollForPollAnswersInvariants(k)(ctx); stop {
			return res, stop
		}

		if res, stop := ValidPostForReactionsInvariants(k)(ctx); stop {
			return res, stop
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

// formatOutputIDs concatenate the ids given into a unique string
func formatOutputIDs(ids types.PostIDs) (outputIDs string) {
	for _, id := range ids {
		outputIDs += id.String() + "\n"
	}
	return outputIDs
}

// Hash256PostIDInvariant checks that the all posts have a SHA256 ID
func Hash256PostIDInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPostIDs types.PostIDs
		k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
			if !post.PostID.Valid() {
				invalidPostIDs = append(invalidPostIDs, post.PostID)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid post IDs",
			fmt.Sprintf("The following post IDs are invalid:\n %s", formatOutputIDs(invalidPostIDs)),
		), invalidPostIDs != nil
	}
}

// ValidCommentsDateInvariants checks that comments creation date is always greater than parent creation date
func ValidCommentsDateInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidCommentsIDs types.PostIDs
		k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
			if post.ParentID.Valid() {
				parentPost, _ := k.GetPost(ctx, post.ParentID)
				if post.Created.Before(parentPost.Created) {
					invalidCommentsIDs = append(invalidCommentsIDs, post.PostID)
				}
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "comments dates before parent post date",
			fmt.Sprintf("The following post IDs referred to posts which are invalid comments "+
				"having creation date before parent post creation date:\n %s",
				formatOutputIDs(invalidCommentsIDs)),
		), invalidCommentsIDs != nil
	}
}

// ValidPostForReactionsInvariants checks that the post related to the reactions is valid and exists
func ValidPostForReactionsInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPostIDs types.PostIDs
		reactions := k.GetReactions(ctx)
		for key := range reactions {
			postID := types.PostID(key)
			if _, found := k.GetPost(ctx, postID); !found {
				invalidPostIDs = append(invalidPostIDs, postID)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "posts reactions refers to non existing posts",
			fmt.Sprintf("The following post IDs referred to posts that did not exist:\n %s",
				formatOutputIDs(invalidPostIDs)),
		), invalidPostIDs != nil
	}
}

// ValidPollForPollAnswersInvariants check that the poll answers are referred to a valid post's poll
func ValidPollForPollAnswersInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPostIDs types.PostIDs
		answers := k.GetPollAnswersMap(ctx)
		for key := range answers {
			postID := types.PostID(key)
			if post, found := k.GetPost(ctx, postID); !found || (found && post.PollData == nil) {
				invalidPostIDs = append(invalidPostIDs, postID)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "poll answers refers to posts without poll",
			fmt.Sprintf("The following post IDs referred to posts that either not exists or has no poll associated:\n %s",
				formatOutputIDs(invalidPostIDs)),
		), invalidPostIDs != nil
	}
}
