package v0150

import (
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	v0150posts "github.com/desmos-labs/desmos/x/staging/posts/types"
)

// GenesisState contains the data of a v0.15.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post               `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts"`
	UsersPollAnswers    []UserAnswersEntry   `protobuf:"bytes,2,rep,name=users_poll_answers,json=usersPollAnswers,proto3" json:"users_poll_answers"`
	PostsReactions      []PostReactionsEntry `protobuf:"bytes,3,rep,name=posts_reactions,json=postsReactions,proto3" json:"posts_reactions"`
	RegisteredReactions []RegisteredReaction `protobuf:"bytes,4,rep,name=registered_reactions,json=registeredReactions,proto3" json:"registered_reactions"`
	Params              Params               `protobuf:"bytes,5,opt,name=params,proto3" json:"params"`
}

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

// Params contains the parameters for the posts module
type Params struct {
	MaxPostMessageLength            github_com_cosmos_cosmos_sdk_types.Int
	MaxOptionalDataFieldsNumber     github_com_cosmos_cosmos_sdk_types.Int
	MaxOptionalDataFieldValueLength github_com_cosmos_cosmos_sdk_types.Int
}
