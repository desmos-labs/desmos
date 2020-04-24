package v040

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/supply"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// MigratePostReactions takes a map of v0.3.0 Reaction objects and migrates
// them to a v0.4.0 map of Reaction objects.
func MigratePostReactions(
	postReactions map[string][]v030posts.Reaction, posts []v030posts.Post,
) (map[string][]PostReaction, error) {
	migratedReactions := make(map[string][]PostReaction, len(postReactions))

	for key, value := range postReactions {

		// Migrate the postReactions
		reactions := make([]PostReaction, len(value))
		for index, reaction := range value {
			value, err := GetReactionShortCodeFromValue(reaction.Value)
			if err != nil {
				return nil, err
			}

			reactions[index] = PostReaction{
				Owner: reaction.Owner,
				Value: value,
			}
		}

		postID, err := ConvertID(key, posts)
		if err != nil {
			return nil, err
		}

		migratedReactions[string(postID)] = RemoveDuplicatedReactions(reactions)
	}

	return migratedReactions, nil
}

// RemoveDuplicatedReactions removes all the duplicated reactions present inside
// the given slice, returning a new one without such duplicates
func RemoveDuplicatedReactions(reactions []PostReaction) (reacts []PostReaction) {

	for _, reaction := range reactions {

		exists := false
		for _, r := range reacts {
			if r.Value == reaction.Value && r.Owner.Equals(reaction.Owner) {
				exists = true
				break
			}
		}

		if !exists {
			reacts = append(reacts, reaction)
		}
	}

	return reacts
}

// GetReactionsToRegister takes the list of posts that exist and the map of all the
// added reactions and returns a list of reactions that should be registered.
func GetReactionsToRegister(
	posts []Post, postReactions map[string][]PostReaction,
) (reactionsToRegister []Reaction, error error) {

	for postID, reactions := range postReactions {

		for _, reaction := range reactions {
			post, err := getPostWithID(posts, PostID(postID))
			if err != nil {
				return nil, err
			}

			if !containsReactionWithCodeForSubspace(reactionsToRegister, reaction.Value, post.Subspace) {
				// nolint: errcheck
				reactionEmoji, _ := emoji.LookupEmojiByCode(reaction.Value)

				reactionsToRegister = append(reactionsToRegister, Reaction{
					ShortCode: reaction.Value,
					Value:     reactionEmoji.Value,
					Subspace:  post.Subspace,
					Creator:   supply.NewModuleAddress(ModuleName),
				})
			}
		}

	}

	return reactionsToRegister, nil
}

// getPostWithID returns the post having the specified id from the given posts list
func getPostWithID(posts []Post, id PostID) (Post, error) {
	for _, post := range posts {
		if post.PostID == id {
			return post, nil
		}
	}

	return Post{}, fmt.Errorf("post with id %s does not exist in list", id)
}

// containsReactionWithCodeForSubspace returns true if the reactions list contains a reaction having
// the specified code, or false otherwise
func containsReactionWithCodeForSubspace(reactions []Reaction, code, subspace string) bool {
	for _, reaction := range reactions {
		if reaction.ShortCode == code && reaction.Subspace == subspace {
			return true
		}
	}

	return false
}
