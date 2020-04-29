package v040

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/supply"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
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

// GetReactionShortCodeFromValue retrieves the shortcode that should
// be associated to the reaction having the given value
func GetReactionShortCodeFromValue(originalValue string) (string, error) {
	// Make the value lowercase and replace - dividers with _
	value := strings.ToLower(originalValue)
	value = strings.ReplaceAll(value, "️", "")
	value = strings.ReplaceAll(value, "-", "_")

	// Try to get any present emoji by considering the value as the emoji itself
	if presentEmojis := emoji.FindAll(value); len(presentEmojis) > 0 {
		return presentEmojis[0].Match.(emoji.Emoji).Shortcodes[0], nil
	}

	value = strings.Split(value, " ")[0]
	// nolint: gocritic
	if value == "like" {
		value = ":heart:"
	} else if value == "true" || value == "q" || strings.Contains(value, "nice") || strings.Contains(value, "well") {
		value = ":+1:"
	} else if value == ":grinning_face_with_star_eyes:" {
		value = ":star-struck:"
	} else if value == ":grinning_face_with_one_large_and_one_small_eye:" {
		value = ":zany_face:"
	} else if value == ":lion_face:" {
		value = ":lion:"
	} else if value == ":star_struck:" {
		value = ":star-struck:"
	} else if value == ":money_mouth_face:" {
		value = ":money-mouth_face:"
	}

	// Try to get the emoji by considering the value as the shortcode
	foundEmoji, err := emoji.LookupEmojiByCode(value)
	if err != nil {
		return "", err
	}

	return foundEmoji.Shortcodes[0], nil
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
