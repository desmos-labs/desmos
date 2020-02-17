package v030

import (
	"crypto/sha256"
	"encoding/hex"

	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
)

// Migrate accepts exported genesis state from v0.2.0 and migrates it to v0.3.0
// genesis state. This migration changes the way subspaces are specified from a simple
// string to a sha-256 hash
func Migrate(oldGenState v020posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:     migratePosts(oldGenState.Posts),
		Reactions: migrateReactions(oldGenState.Reactions),
	}
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
func migrateReactions(likes map[string][]v020posts.Reaction) map[string][]Reaction {
	migratedLikes := make(map[string][]Reaction, len(likes))

	for key, value := range likes {

		// Migrate the likes
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
