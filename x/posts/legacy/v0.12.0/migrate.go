package v0120

import (
	v0100posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
)

// Migrate accepts exported genesis state from v0.10.0 and migrates it to v0.12.0
// genesis state. This migration replace all the old poll structure with the new one
// removing the Open field from it.
func Migrate(oldGenState v0100posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               ConvertPosts(oldGenState.Posts),
		UsersPollAnswers:    oldGenState.UsersPollAnswers,
		PostReactions:       oldGenState.PostReactions,
		RegisteredReactions: oldGenState.RegisteredReactions,
		Params:              oldGenState.Params,
	}
}

// ConvertPosts v0.10.0 posts into v0.12.0 posts
func ConvertPosts(oldPosts []v0100posts.Post) []Post {
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
			OptionalData:   post.OptionalData,
			Creator:        post.Creator,
			Attachments:    post.Attachments,
			PollData:       ConvertPollData(post.PollData),
		}
	}
	return posts
}

// ConvertPollData converts the given v0.10.0 PollData to a new v0.12.0 PollData
func ConvertPollData(pollData *v040posts.PollData) *PollData {
	if pollData == nil {
		return nil
	}

	return &PollData{
		Question:              pollData.Question,
		ProvidedAnswers:       pollData.ProvidedAnswers,
		EndDate:               pollData.EndDate,
		AllowsMultipleAnswers: pollData.AllowsMultipleAnswers,
		AllowsAnswerEdits:     pollData.AllowsAnswerEdits,
	}
}
