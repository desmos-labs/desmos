package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "hash256-post-id",
		ValidPostsInvariants(keeper))
	ir.RegisterRoute(types.ModuleName, "comments-date",
		ValidCommentsDateInvariants(keeper))
	ir.RegisterRoute(types.ModuleName, "post-reactions",
		ValidPostForReactionsInvariants(keeper))
	ir.RegisterRoute(types.ModuleName, "post-poll-answers",
		ValidPollForPollAnswersInvariants(keeper))
}

// AllInvariants runs all invariants of the module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := ValidPostsInvariants(k)(ctx); stop {
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

// ValidPostsInvariants checks that the all posts are valid
func ValidPostsInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPostIDs types.PostIDs
		k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
			if post.Validate() != nil {
				invalidPostIDs = append(invalidPostIDs, post.PostID)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid posts IDs",
			fmt.Sprintf("The following posts are invalid:\n %s", formatOutputIDs(invalidPostIDs)),
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

// formatOutputReactions concatenate the reactions given into a unique string
func formatOutputReactions(reactions types.PostReactions) (outputReactions string) {
	for _, reaction := range reactions {
		outputReactions += reaction.String() + "\n"
	}
	return outputReactions
}

// ValidPostForReactionsInvariants checks that the post related to the reactions is valid and exists
func ValidPostForReactionsInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidReactions types.PostReactions
		reactions := k.GetReactions(ctx)
		for key, value := range reactions {
			postID := types.PostID(key)
			if _, found := k.GetPost(ctx, postID); !found {
				invalidReactions = append(invalidReactions, value...)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "posts reactions refers to non existing posts",
			fmt.Sprintf("The following reactions refer to posts that do not exist:\n %s",
				formatOutputReactions(invalidReactions)),
		), invalidReactions != nil
	}
}

// formatOutputPollAnswers concatenate the poll answers given into a unique string
func formatOutputPollAnswers(pollAnswers types.UserAnswers) (outputAnswers string) {
	for _, answer := range pollAnswers {
		outputAnswers += answer.String() + "\n"
	}
	return outputAnswers
}

// ValidPollForPollAnswersInvariants check that the poll answers are referred to a valid post's poll
func ValidPollForPollAnswersInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPollAnswers types.UserAnswers
		answers := k.GetPollAnswersMap(ctx)
		for key, value := range answers {
			postID := types.PostID(key)
			if post, found := k.GetPost(ctx, postID); !found || (found && post.PollData == nil) {
				invalidPollAnswers = append(invalidPollAnswers, value...)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "poll answers refers to posts without poll",
			fmt.Sprintf("The following answers refer to a post that either does not exist or has no poll associated to it:\n %s",
				formatOutputPollAnswers(invalidPollAnswers)),
		), invalidPollAnswers != nil
	}
}
