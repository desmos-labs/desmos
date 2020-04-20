package v040

import (
	"fmt"
	"sort"

	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
)

// Migrate accepts exported genesis state from v0.3.0 and migrates it to v0.4.0
// genesis state. This migration changes the way posts IDs are specified from a simple
// uint64 to a sha-256 hashed string
func Migrate(oldGenState v030posts.GenesisState) GenesisState {
	posts, postReactions := RemoveDoublePosts(oldGenState.Posts, oldGenState.Reactions)
	posts, postReactions = SquashPostIDs(posts, postReactions)

	return GenesisState{
		Posts:               migratePosts(posts),
		UsersPollAnswers:    migrateUsersAnswers(oldGenState.PollAnswers, posts),
		PostReactions:       migratePostReactions(postReactions, posts),
		RegisteredReactions: []Reaction{},
	}
}

// RemoveDoublePosts removes all the duplicated posts that have the same contents from the given posts slice.
// All the reactions associated with the removed posts are moved to the post with the lower id.
func RemoveDoublePosts(posts []v030posts.Post, postReactions map[string][]v030posts.Reaction) ([]v030posts.Post, map[string][]v030posts.Reaction) {
	var newPosts []v030posts.Post
	newReactions := make(map[string][]v030posts.Reaction, len(postReactions))

	for index, post := range posts {

		var conflictingID v030posts.PostID
		for _, other := range posts[0:index] {
			if post.ConflictsWith(other) {

				// If the post has the same contents we just skip it and move all the postReactions to the original one
				if post.ContentsEquals(other) {
					conflictingID = other.PostID
					break
				}

				panic(fmt.Errorf("post with id %s and that with id %s are conflicting but do not contains the same data",
					post.PostID, other.PostID))
			}
		}

		postReactions := postReactions[post.PostID.String()]
		if conflictingID > 0 {
			originalReactions := newReactions[conflictingID.String()]
			newReactions[conflictingID.String()] = append(originalReactions, postReactions...)
		} else {
			newPosts = append(newPosts, post)
			if len(postReactions) > 0 {
				newReactions[post.PostID.String()] = postReactions
			}
		}
	}

	return newPosts, newReactions
}

// SquashPostIDs iterates over all the given posts and reactions and squashes the post IDs
// to be all subsequent to each other so that no IDs are skipped
func SquashPostIDs(posts []v030posts.Post, postReactions map[string][]v030posts.Reaction) ([]v030posts.Post, map[string][]v030posts.Reaction) {
	// Sort the posts for easier management
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].PostID < posts[j].PostID
	})

	newPosts := make([]v030posts.Post, len(posts))
	newReactions := make(map[string][]v030posts.Reaction)

	postIDsReplacements := make(map[string]v030posts.PostID)

	for index, post := range posts {
		oldPostID := post.PostID
		newPostID := v030posts.PostID(index + 1)

		post.PostID = newPostID
		postIDsReplacements[oldPostID.String()] = newPostID

		// Change the parent id if previously replaced
		newParentID := postIDsReplacements[post.ParentID.String()]
		if newParentID > 0 {
			post.ParentID = newParentID
		}

		newPosts[index] = post

		// Update the postReactions reference if some reaction exists for this post
		postReactions := postReactions[oldPostID.String()]
		if len(postReactions) > 0 {
			newReactions[newPostID.String()] = postReactions
		}
	}

	return newPosts, newReactions
}

// ComputeParentID get the post related to the given parentID if exists and returns it computed ID.
// Returns "" otherwise
func ComputeParentID(posts []v030posts.Post, parentID v030posts.PostID) PostID {
	if parentID == v030posts.PostID(uint64(0)) {
		return ""
	}

	for _, post := range posts {
		if post.ParentID == parentID {
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
