package v030

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"

	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
)

// Migrate accepts exported genesis state from v0.2.0 and migrates it to v0.3.0
// genesis state. This migration changes the way subspaces are specified from a simple
// string to a sha-256 hash
func Migrate(oldGenState v020posts.GenesisState) GenesisState {
	posts, reactions := RemoveDoublePosts(oldGenState.Posts, oldGenState.Reactions)
	posts, reactions = SquashPostIDs(posts, reactions)

	return GenesisState{
		Posts:       migratePosts(posts),
		PollAnswers: nil,
		Reactions:   migrateReactions(reactions),
	}
}

// RemoveDoublePosts removes all the duplicated posts that have the same contents from the given posts slice.
// All the reactions associated with the removed posts are moved to the post with the lower id.
func RemoveDoublePosts(posts []v020posts.Post, reactions map[string][]v020posts.Reaction) ([]v020posts.Post, map[string][]v020posts.Reaction) {
	var newPosts []v020posts.Post
	newReactions := make(map[string][]v020posts.Reaction, len(reactions))

	for index, post := range posts {

		var conflictingID v020posts.PostID
		for _, other := range posts[0:index] {
			if post.ConflictsWith(other) {

				// If the post has the same contents we just skip it and move all the reactions to the original one
				if post.ContentsEquals(other) {
					conflictingID = other.PostID
					break
				}

				panic(fmt.Errorf("post with id %s and that with id %s are conflicting but do not contains the same data",
					post.PostID, other.PostID))
			}
		}

		postReactions := reactions[post.PostID.String()]
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
func SquashPostIDs(posts []v020posts.Post, reactions map[string][]v020posts.Reaction) ([]v020posts.Post, map[string][]v020posts.Reaction) {
	// Sort the posts for easier management
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].PostID < posts[j].PostID
	})

	newPosts := make([]v020posts.Post, len(posts))
	newReactions := make(map[string][]v020posts.Reaction)

	postIDsReplacements := make(map[string]v020posts.PostID)

	for index, post := range posts {
		oldPostID := post.PostID
		newPostID := v020posts.PostID(index + 1)

		post.PostID = newPostID
		postIDsReplacements[oldPostID.String()] = newPostID

		// Change the parent id if previously replaced
		newParentID := postIDsReplacements[post.ParentID.String()]
		if newParentID > 0 {
			post.ParentID = newParentID
		}

		newPosts[index] = post

		// Update the reactions reference if some reaction exists for this post
		postReactions := reactions[oldPostID.String()]
		if len(postReactions) > 0 {
			newReactions[newPostID.String()] = postReactions
		}
	}

	return newPosts, newReactions
}

// migratePosts takes a slice of v0.2.0 Post object and migrates them to v0.3.0 Post
// For each post, if the subspace is a valid sha-256 hash it is preserved the way it is, otherwise it is
// converted by hashing it.
func migratePosts(posts []v020posts.Post) []Post {
	migratedPosts := make([]Post, len(posts))

	// Migrate the posts
	for index, oldPost := range posts {
		subspace := oldPost.Subspace
		if !SubspaceRegEx.MatchString(subspace) {
			hash := sha256.Sum256([]byte(subspace))
			subspace = hex.EncodeToString(hash[:])
		}

		migratedPosts[index] = Post{
			PostID:         PostID(oldPost.PostID),
			ParentID:       PostID(oldPost.ParentID),
			Message:        oldPost.Message,
			Created:        oldPost.Created,
			LastEdited:     oldPost.LastEdited,
			AllowsComments: oldPost.AllowsComments,
			Subspace:       subspace,
			OptionalData:   OptionalData(oldPost.OptionalData),
			Creator:        oldPost.Creator,
		}
	}

	return migratedPosts
}

// migrateReactions takes a map of v0.2.0 Reaction objects and migrates them to a v0.3.0 map of Reaction objects.
func migrateReactions(reactions map[string][]v020posts.Reaction) map[string][]Reaction {
	migratedLikes := make(map[string][]Reaction, len(reactions))

	for key, value := range reactions {

		// Migrate the reactions
		reactions := make([]Reaction, len(value))
		for index, reaction := range value {
			reactions[index] = Reaction{
				Owner: reaction.Owner,
				Value: reaction.Value,
			}
		}

		migratedLikes[key] = reactions
	}

	return migratedLikes
}
