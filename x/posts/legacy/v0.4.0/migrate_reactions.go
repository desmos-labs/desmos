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
	migratedLikes := make(map[string][]PostReaction, len(postReactions))

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

		migratedLikes[string(postID)] = reactions
	}

	return migratedLikes, nil
}

// GetReactionsToRegister takes the list of posts that exist and the map of all the
// added reactions and returns a list of reactions that should be registered.
func GetReactionsToRegister(
	posts []Post, postReactions map[string][]PostReaction,
) (reactionsToRegister []Reaction, error error) {

	for postID, reactions := range postReactions {

		for _, reaction := range reactions {

			if !containsReactionWithCode(reactionsToRegister, reaction.Value) {
				post, err := getPostWithId(posts, PostID(postID))
				if err != nil {
					return nil, err
				}

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

// getPostWithId returns the post having the specified id from the given posts list
func getPostWithId(posts []Post, id PostID) (Post, error) {
	for _, post := range posts {
		if post.PostID == id {
			return post, nil
		}
	}

	return Post{}, fmt.Errorf("post with id %s does not exist in list", id)
}

// containsReactionWithCode returns true if the reactions list contains a reaction having
// the specified code, or false otherwise
func containsReactionWithCode(reactions []Reaction, code string) bool {
	for _, reaction := range reactions {
		if reaction.ShortCode == code {
			return true
		}
	}

	return false
}
