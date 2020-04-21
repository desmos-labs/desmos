package v040

import (
	"fmt"

	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
)

// Migrate accepts exported genesis state from v0.3.0 and migrates it to v0.4.0
// genesis state. This migration changes the way posts IDs are specified from a simple
// uint64 to a sha-256 hashed string
func Migrate(oldGenState v030posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               migratePosts(oldGenState.Posts),
		UsersPollAnswers:    migrateUsersAnswers(oldGenState.PollAnswers, oldGenState.Posts),
		PostReactions:       migratePostReactions(oldGenState.Reactions, oldGenState.Posts),
		RegisteredReactions: []Reaction{},
	}
}

// ComputeParentID get the post related to the given parentID if exists and returns it computed ID.
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

// migratePosts takes a slice of v0.2.0 Post object and migrates them to v0.3.0 Post
// For each post, if the subspace is a valid sha-256 hash it is preserved the way it is, otherwise it is
// converted by hashing it.
func migratePosts(posts []v030posts.Post) []Post {
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
		}
	}

	return migratedPosts
}

// migrateUsersAnswers takes a slice of v0.4.0 UsersAnswers object and migrates them to v0.4.0 UserAnswers
func migrateUsersAnswers(usersAnswersMap map[string][]v030posts.UserAnswer, posts []v030posts.Post) map[string][]UserAnswer {
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

		var postID PostID
		for _, post := range posts {
			convertedID, err := v030posts.ParsePostID(key)
			if err != nil {
				panic(fmt.Errorf("postID parsing error (migration): %s", err))
			}
			if post.PostID == convertedID {
				postID = ComputeID(post.Created, post.Creator, post.Subspace)
			}
		}

		migratedUsersAnswers[string(postID)] = newUserAnswers
	}

	return migratedUsersAnswers
}

// migrateReactions takes a map of v0.3.0 Reaction objects and migrates them to a v0.4.0 map of Reaction objects.
func migratePostReactions(postReactions map[string][]v030posts.Reaction, posts []v030posts.Post) map[string][]PostReaction {
	migratedLikes := make(map[string][]PostReaction, len(postReactions))

	for key, value := range postReactions {

		// Migrate the postReactions
		reactions := make([]PostReaction, len(value))
		for index, reaction := range value {
			reactions[index] = PostReaction{
				Owner: reaction.Owner,
				Value: reaction.Value,
			}
		}

		var postID PostID
		for _, post := range posts {
			convertedID, err := v030posts.ParsePostID(key)
			if err != nil {
				panic(fmt.Errorf("postID parsing error (migration): %s", err))
			}
			if post.PostID == convertedID {
				postID = ComputeID(post.Created, post.Creator, post.Subspace)
			}
		}

		migratedLikes[string(postID)] = reactions
	}

	return migratedLikes
}
