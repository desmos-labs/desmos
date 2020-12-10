package v0150

import (
	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
)

// Migrate accepts exported genesis state from v0.13.0 and migrates it to v0.15.0
// genesis state.
func Migrate(oldGenState v0130posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               ConvertPosts(oldGenState.Posts),
		UsersPollAnswers:    ConvertUsersPollAnswers(oldGenState.UsersPollAnswers),
		PostReactions:       ConvertPostReactions(oldGenState.PostReactions),
		RegisteredReactions: ConvertRegisteredReactions(oldGenState.RegisteredReactions),
		Params:              oldGenState.Params,
	}
}

func ConvertPosts(oldPosts []v0130posts.Post) []Post {
	posts := make([]Post, len(oldPosts))
	for index, post := range oldPosts {
		posts[index] = Post{
			PostID:         string(post.PostID),
			ParentID:       string(post.ParentID),
			Message:        post.Message,
			Created:        post.Created,
			LastEdited:     post.LastEdited,
			AllowsComments: post.AllowsComments,
			Subspace:       post.Subspace,
			OptionalData:   post.OptionalData,
			Creator:        post.Creator.String(),
			Attachments:    post.Attachments,
			PollData:       post.PollData,
		}
	}
	return posts
}

func ConvertUsersPollAnswers(oldUsersPollAnswers map[string][]v040posts.UserAnswer) []UserAnswersEntry {
	userAnswersEntries := make([]UserAnswersEntry, len(oldUsersPollAnswers))
	index := 0
	for key, value := range oldUsersPollAnswers {
		userAnswersEntries[index] = newUserAnswerEntry(key, value)
		index++
	}

	return userAnswersEntries
}

func ConvertPostReactions(oldPostReactions map[string][]v060.PostReaction) []PostReactionsEntry {
	postReactionsEntries := make([]PostReactionsEntry, len(oldPostReactions))
	index := 0
	for key, value := range oldPostReactions {
		postReactionsEntries[index] = newPostReactionEntry(key, value)
		index++
	}

	return postReactionsEntries
}

func ConvertRegisteredReactions(oldRegisteredReactions []v040posts.Reaction) []RegisteredReaction {
	registeredReactions := make([]RegisteredReaction, len(oldRegisteredReactions))
	for index, reaction := range oldRegisteredReactions {
		registeredReactions[index] = RegisteredReaction{
			ShortCode: reaction.ShortCode,
			Value:     reaction.Value,
			Subspace:  reaction.Subspace,
			Creator:   reaction.Creator.String(),
		}
	}

	return registeredReactions
}
