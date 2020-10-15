package v0130

import (
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
)

// Migrate accepts exported genesis state from v0.12.0 and migrates it to v0.13.0
// genesis state. This migration replace the old optional data map with the new struct
func Migrate(oldGenState v0120posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               ConvertPosts(oldGenState.Posts),
		UsersPollAnswers:    oldGenState.UsersPollAnswers,
		PostReactions:       oldGenState.PostReactions,
		RegisteredReactions: oldGenState.RegisteredReactions,
		Params:              oldGenState.Params,
	}
}

// ConvertPosts v0.12.0 posts into v0.13.0 posts
func ConvertPosts(oldPosts []v0120posts.Post) []Post {
	posts := make([]Post, len(oldPosts))
	for index, post := range oldPosts {
		posts[index] = Post{
			PostID:         post.PostID,
			ParentID:       post.ParentID,
			Message:        post.Message,
			Created:        post.Created,
			LastEdited:     post.LastEdited,
			AllowsComments: post.AllowsComments,
			Subspace:       post.Subspace,
			OptionalData:   ConvertOptionalData(post.OptionalData),
			Creator:        post.Creator,
			Attachments:    post.Attachments,
			PollData:       post.PollData,
		}
	}
	return posts
}

func ConvertOptionalData(oldOptionalData v040posts.OptionalData) []OptionalDataEntry {
	optionalData := make([]OptionalDataEntry, len(oldOptionalData))
	index := 0
	for key, value := range oldOptionalData {
		optionalData[index] = OptionalDataEntry{
			Key:   key,
			Value: value,
		}
		index++
	}

	return optionalData
}
