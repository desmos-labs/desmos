package v0150

import (
	v0150posts "github.com/desmos-labs/desmos/x/staging/posts/types"
)

// GenesisState contains the data of a v0.15.0 genesis state for the posts module
type GenesisState = v0150posts.GenesisState

func FindUserAnswerEntryForPostID(state GenesisState, postID string) (bool, *UserAnswersEntry) {
	for _, entry := range state.UsersPollAnswers {
		if entry.PostId == postID {
			return true, &entry
		}
	}
	return false, nil
}

func FindPostReactionEntryForPostID(state GenesisState, postID string) (bool, *PostReactionsEntry) {
	for _, entry := range state.PostsReactions {
		if entry.PostId == postID {
			return true, &entry
		}
	}
	return false, nil
}

// ----------------------------------------------------------------------------------------------------------------

type Post = v0150posts.Post
type OptionalDataEntry = v0150posts.OptionalDataEntry
type Attachment = v0150posts.Attachment
type PollData = v0150posts.PollData
type PollAnswer = v0150posts.PollAnswer

// ----------------------------------------------------------------------------------------------------------------

type UserAnswersEntry = v0150posts.UserAnswersEntry
type UserAnswer = v0150posts.UserAnswer

// ----------------------------------------------------------------------------------------------------------------

type PostReactionsEntry = v0150posts.PostReactionsEntry
type PostReaction = v0150posts.PostReaction

// ----------------------------------------------------------------------------------------------------------------

type RegisteredReaction = v0150posts.RegisteredReaction

// ----------------------------------------------------------------------------------------------------------------

type Params = v0150posts.Params
