package v0150

import (
	"strconv"

	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
)

const (
	ModuleName = "posts"
)

// Migrate accepts exported genesis state from v0.13.0 and migrates it to v0.15.0
// genesis state.
func Migrate(oldGenState v0130posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               ConvertPosts(oldGenState.Posts),
		UsersPollAnswers:    ConvertUsersPollAnswers(oldGenState.UsersPollAnswers),
		PostsReactions:      ConvertPostReactions(oldGenState.PostReactions),
		RegisteredReactions: ConvertRegisteredReactions(oldGenState.RegisteredReactions),
		Params:              Params(oldGenState.Params),
	}
}

// ----------------------------------------------------------------------------------------------------------------

func ConvertPosts(oldPosts []v0130posts.Post) []Post {
	posts := make([]Post, len(oldPosts))
	for index, post := range oldPosts {
		posts[index] = Post{
			PostId:         post.PostId,
			ParentId:       post.ParentId,
			Message:        post.Message,
			Created:        post.Created,
			LastEdited:     post.LastEdited,
			AllowsComments: post.AllowsComments,
			Subspace:       post.Subspace,
			OptionalData:   convertOptionalData(post.OptionalData),
			Creator:        post.Creator.String(),
			Attachments:    convertAttachments(post.Attachments),
			PollData:       convertPollData(post.PollData),
		}
	}
	return posts
}

func convertOptionalData(old []v0130posts.OptionalDataEntry) []OptionalDataEntry {
	data := make([]OptionalDataEntry, len(old))
	for index, entry := range old {
		data[index] = OptionalDataEntry(entry)
	}
	return data
}

func convertAttachments(old []v0130posts.Attachment) []Attachment {
	attachments := make([]Attachment, len(old))
	for index, attachment := range old {
		tags := make([]string, len(attachment.Tags))
		for index, tag := range attachment.Tags {
			tags[index] = tag.String()
		}

		attachments[index] = Attachment{
			URI:      attachment.URI,
			MimeType: attachment.MimeType,
			Tags:     tags,
		}
	}
	return attachments
}

func convertPollData(old *v0130posts.PollData) *PollData {
	if old == nil {
		return nil
	}

	answers := make([]PollAnswer, len(old.ProvidedAnswers))
	for index, answer := range old.ProvidedAnswers {
		answers[index] = PollAnswer{
			ID:   strconv.FormatUint(answer.ID, 10),
			Text: answer.Text,
		}
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

func ConvertUsersPollAnswers(oldUsersPollAnswers map[string][]v0130posts.UserAnswer) []UserAnswersEntry {
	userAnswersEntries := make([]UserAnswersEntry, len(oldUsersPollAnswers))
	index := 0
	for key, value := range oldUsersPollAnswers {
		userAnswersEntries[index] = creteUserAnswerEntry(key, value)
		index++
	}

	return userAnswersEntries
}

// creteUserAnswerEntry create a new userAnswerEntry from the old data from genesis
func creteUserAnswerEntry(postID string, oldUsersAnswers []v0130posts.UserAnswer) UserAnswersEntry {
	userAnswers := make([]UserAnswer, len(oldUsersAnswers))

	for index, oldUserAnswers := range oldUsersAnswers {
		answers := make([]string, len(oldUserAnswers.Answers))
		for index, id := range oldUserAnswers.Answers {
			answers[index] = strconv.FormatUint(id, 10)
		}

		userAnswers[index] = UserAnswer{
			User:    oldUserAnswers.User.String(),
			Answers: answers,
		}
	}

	return UserAnswersEntry{
		PostId:      postID,
		UserAnswers: userAnswers,
	}
}

// ----------------------------------------------------------------------------------------------------------------

func ConvertPostReactions(oldPostReactions map[string][]v0130posts.PostReaction) []PostReactionsEntry {
	postReactionsEntries := make([]PostReactionsEntry, len(oldPostReactions))
	index := 0
	for key, value := range oldPostReactions {
		postReactionsEntries[index] = createPostReactionEntry(key, value)
		index++
	}

	return postReactionsEntries
}

// createPostReactionEntry create a new PostReactionEntry from the old data from genesis
func createPostReactionEntry(postID string, oldPostReactions []v0130posts.PostReaction) PostReactionsEntry {
	reactions := make([]PostReaction, len(oldPostReactions))

	for index, oldPostReaction := range oldPostReactions {
		reactions[index] = PostReaction{
			ShortCode: oldPostReaction.Shortcode,
			Value:     oldPostReaction.Value,
			Owner:     oldPostReaction.Owner.String(),
		}
	}

	return PostReactionsEntry{
		PostId:    postID,
		Reactions: reactions,
	}
}

// ----------------------------------------------------------------------------------------------------------------

func ConvertRegisteredReactions(oldRegisteredReactions []v0130posts.RegisteredReaction) []RegisteredReaction {
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
