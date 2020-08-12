package v040

import (
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
)

// Migrate accepts exported genesis state from v0.3.0 and migrates it to v0.4.0
// genesis state. This migration changes the way posts IDs are specified from a simple
// uint64 to a sha-256 hashed string
func Migrate(oldGenState v030posts.GenesisState) GenesisState {

	posts := MigratePosts(oldGenState.Posts)

	answers, err := MigrateUsersAnswers(oldGenState.PollAnswers, oldGenState.Posts)
	if err != nil {
		panic(err)
	}

	reactions, err := MigratePostReactions(oldGenState.Reactions, oldGenState.Posts)
	if err != nil {
		panic(err)
	}

	registeredReactions, err := GetReactionsToRegister(posts, reactions)
	if err != nil {
		panic(err)
	}

	return GenesisState{
		Posts:               posts,
		UsersPollAnswers:    answers,
		PostReactions:       reactions,
		RegisteredReactions: registeredReactions,
	}
}

// ComputeParentID get the post related to the given parentID if exists and returns it computed RelationshipID.
// Returns "" otherwise
func ComputeParentID(posts []v030posts.Post, parentID v030posts.PostID) PostID {
	if parentID == v030posts.PostID(uint64(0)) {
		return ""
	}

	for _, post := range posts {
		if post.PostID == parentID {
			return ComputeID(post.Created, post.Creator, post.Subspace)
		}
	}

	//it should never reach this
	return ""
}

// MigratePosts takes a slice of v0.3.0 Post object and migrates them to v0.4.0 Post.
// For each post, its id is converted from an uint64 representation to a SHA-256 string representation.
func MigratePosts(posts []v030posts.Post) []Post {
	migratedPosts := make([]Post, len(posts))

	// Migrate the posts
	for index, oldPost := range posts {
		migratedPosts[index] = Post{
			PostID:         ComputeID(oldPost.Created, oldPost.Creator, oldPost.Subspace),
			ParentID:       ComputeParentID(posts, oldPost.ParentID),
			Message:        oldPost.Message,
			Created:        oldPost.Created,
			LastEdited:     oldPost.LastEdited,
			AllowsComments: oldPost.AllowsComments,
			Subspace:       oldPost.Subspace,
			OptionalData:   OptionalData(oldPost.OptionalData),
			Creator:        oldPost.Creator,
			Medias:         MigrateMedias(oldPost.Medias),
			PollData:       MigratePollData(oldPost.PollData),
		}
	}

	return migratedPosts
}

// ConvertID take the given v030 post RelationshipID and convert it to a v040 post RelationshipID
func ConvertID(id string, posts []v030posts.Post) (postID PostID, error error) {
	for _, post := range posts {
		convertedID, err := v030posts.ParsePostID(id)
		if err != nil {
			return "", err
		}

		if post.PostID == convertedID {
			postID = ComputeID(post.Created, post.Creator, post.Subspace)
		}
	}
	return postID, nil
}

// MigrateMedias takes the given v030 medias and converts them into a v040 medias array
func MigrateMedias(medias []v030posts.PostMedia) []PostMedia {
	newMedias := make([]PostMedia, len(medias))
	for index, media := range medias {
		newMedias[index] = PostMedia{
			URI:      media.URI,
			MimeType: media.MimeType,
		}
	}
	return newMedias
}

// MigrateMedias takes the given v030 poll data and converts it into a v040 poll data
func MigratePollData(data *v030posts.PollData) *PollData {
	if data == nil {
		return nil
	}

	answers := make([]PollAnswer, len(data.ProvidedAnswers))
	for index, answer := range data.ProvidedAnswers {
		answers[index] = PollAnswer{
			ID:   AnswerID(answer.ID),
			Text: answer.Text,
		}
	}

	return &PollData{
		Question:              data.Question,
		ProvidedAnswers:       answers,
		EndDate:               data.EndDate,
		Open:                  data.Open,
		AllowsMultipleAnswers: data.AllowsMultipleAnswers,
		AllowsAnswerEdits:     data.AllowsAnswerEdits,
	}
}

// MigrateUsersAnswers takes a slice of v0.3.0 UsersAnswers object and migrates them to v0.4.0 UserAnswers
func MigrateUsersAnswers(
	usersAnswersMap map[string][]v030posts.UserAnswer, posts []v030posts.Post,
) (map[string][]UserAnswer, error) {
	migratedUsersAnswers := make(map[string][]UserAnswer, len(usersAnswersMap))

	//Migrate the users answers
	for key, value := range usersAnswersMap {

		newUserAnswers := make([]UserAnswer, len(value))
		for index, userAnswers := range value {
			migratedAnswersIDs := make([]AnswerID, len(userAnswers.Answers))

			for index, answerID := range userAnswers.Answers {
				migratedAnswersIDs[index] = AnswerID(answerID)
			}
			newUserAnswers[index] = UserAnswer{
				Answers: migratedAnswersIDs,
				User:    userAnswers.User,
			}
		}

		postID, err := ConvertID(key, posts)
		if err != nil {
			return nil, err
		}

		migratedUsersAnswers[string(postID)] = newUserAnswers
	}

	return migratedUsersAnswers, nil
}
