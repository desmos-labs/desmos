package v0130

import (
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
)

const (
	ModuleName = "posts"
)

// Migrate accepts exported genesis state from v0.12.0 and migrates it to v0.13.0
// genesis state. This migration replace the old optional data map with the new struct
func Migrate(oldGenState v0120posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               ConvertPosts(oldGenState.Posts),
		UsersPollAnswers:    ConvertUserPollAnswers(oldGenState.UsersPollAnswers),
		PostReactions:       ConvertPostReactions(oldGenState.PostReactions),
		RegisteredReactions: ConvertRegisteredReactions(oldGenState.RegisteredReactions),
		Params:              Params(oldGenState.Params),
	}
}

// ----------------------------------------------------------------------------------------------------------------

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
			Attachments:    ConvertAttachments(post.Attachments),
			PollData:       ConvertPollData(post.PollData),
		}
	}
	return posts
}

func ConvertOptionalData(oldOptionalData v0120posts.OptionalData) []OptionalDataEntry {
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

func ConvertAttachments(old []v0120posts.Attachment) []Attachment {
	attachments := make([]Attachment, len(old))
	for index, attachment := range old {
		attachments[index] = Attachment(attachment)
	}
	return attachments
}

func ConvertPollData(old *v0120posts.PollData) *PollData {
	if old == nil {
		return nil
	}

	answers := make([]PollAnswer, len(old.ProvidedAnswers))
	for index, answer := range old.ProvidedAnswers {
		answers[index] = PollAnswer(answer)
	}

	return &PollData{
		Question:              old.Question,
		ProvidedAnswers:       answers,
		EndDate:               old.EndDate,
		AllowsMultipleAnswers: old.AllowsMultipleAnswers,
		AllowsAnswerEdits:     old.AllowsAnswerEdits,
	}
}

// ----------------------------------------------------------------------------------------------------------------

func ConvertUserPollAnswers(old map[string][]v0120posts.UserAnswer) map[string][]UserAnswer {
	newAnswers := make(map[string][]UserAnswer, len(old))
	for key, value := range old {
		answers := make([]UserAnswer, len(value))
		for index, answer := range value {
			answers[index] = UserAnswer(answer)
		}

		newAnswers[key] = answers
	}
	return newAnswers
}

// ----------------------------------------------------------------------------------------------------------------

func ConvertPostReactions(old map[string][]v0120posts.PostReaction) map[string][]PostReaction {
	newReactions := make(map[string][]PostReaction, len(old))
	for key, value := range old {
		reactions := make([]PostReaction, len(value))
		for index, reaction := range value {
			reactions[index] = PostReaction(reaction)
		}
		newReactions[key] = reactions
	}
	return newReactions
}

// ----------------------------------------------------------------------------------------------------------------

func ConvertRegisteredReactions(old []v0120posts.RegisteredReaction) []RegisteredReaction {
	reactions := make([]RegisteredReaction, len(old))
	for index, reaction := range old {
		reactions[index] = RegisteredReaction(reaction)
	}
	return reactions
}
