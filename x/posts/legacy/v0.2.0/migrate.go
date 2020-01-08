package v0_2_0

import (
	"strings"

	v010posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.1.0"
)

// Migrate accepts exported genesis state from v0.1.0 and migrates it to v0.2.0
// genesis state. This migration changes the data that are saved for each post
// moving the external reference into the arbitrary data map and adding the new
// subspace field
func Migrate(oldGenState v010posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:     migratePosts(oldGenState.Posts),
		Reactions: migrateLikes(oldGenState.Likes),
	}
}

// migratePosts takes a slice of v0.1.0 Post object and migrates them to v0.2.0 Post
// The following changes are performed:
// - Each external_reference Post value (if not empty or blank) is put inside the optional_data map using
//   external_reference as the value's associated key
// - Post subspaces are left empty so that they can be properly set after the migration has been completed
func migratePosts(posts []v010posts.Post) []Post {
	migratedPosts := make([]Post, len(posts))

	// Migrate the posts
	for index, oldPost := range posts {

		optionalData := map[string]string{}
		if len(strings.TrimSpace(oldPost.ExternalReference)) != 0 {
			optionalData["external_reference"] = oldPost.ExternalReference
		}

		migratedPosts[index] = Post{
			PostID:         PostID(oldPost.PostID),
			ParentID:       PostID(oldPost.ParentID),
			Message:        oldPost.Message,
			Created:        oldPost.Created,
			LastEdited:     oldPost.LastEdited,
			AllowsComments: oldPost.AllowsComments,
			Subspace:       "",
			OptionalData:   optionalData,
			Owner:          oldPost.Owner,
		}
	}

	return migratedPosts
}

// migrateLikes takes a map of v0.1.0 Like objects and migrates them to a v0.2.0 map of Reaction objects.
// Each Like is migrated to a Reaction object by preserving the associated PostID, Creator and Created values
// but setting its value to "like".
func migrateLikes(likes map[string][]v010posts.Like) map[string][]Reaction {
	migratedLikes := make(map[string][]Reaction, len(likes))

	for key, value := range likes {

		// Migrate the likes
		reactions := make([]Reaction, len(value))
		for index, like := range value {
			reactions[index] = Reaction{
				Created: like.Created,
				Owner:   like.Owner,
				Value:   "like",
			}
		}

		migratedLikes[key] = reactions
	}

	return migratedLikes
}
