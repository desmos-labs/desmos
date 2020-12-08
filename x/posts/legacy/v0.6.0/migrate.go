package v060

import (
	emoji "github.com/desmos-labs/Go-Emoji-Utils"

	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
)

// Migrate accepts exported genesis state from v0.4.0 and migrates it to v0.6.0
// genesis state. This migration replace all the old post reactions structure
// with the new one that includes shortcodes.
func Migrate(oldGenState v040posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               oldGenState.Posts,
		UsersPollAnswers:    oldGenState.UsersPollAnswers,
		PostReactions:       migratePostReactions(oldGenState.PostReactions, oldGenState.RegisteredReactions),
		RegisteredReactions: oldGenState.RegisteredReactions,
	}
}

// GetConvertedPostReaction convert the given 0.4.0 post reaction to a v0.6.0 post reaction.
func GetConvertedPostReaction(reaction v040posts.PostReaction, regReactions []v040posts.Reaction) (postReaction PostReaction) {
	// check if the given post reaction is an emoji
	if em, err := emoji.LookupEmojiByCode(reaction.Value); err != nil {
		// if not check among the registered reaction
		for _, regReaction := range regReactions {
			if regReaction.ShortCode == reaction.Value {
				postReaction = PostReaction{
					Owner:     reaction.Owner,
					Shortcode: reaction.Value,
					Value:     regReaction.Value,
				}
			}
		}
	} else {
		postReaction = PostReaction{
			Owner:     reaction.Owner,
			Shortcode: reaction.Value,
			Value:     em.Value,
		}
	}

	return postReaction
}

// migratePostReactions migrate all the given v0.4.0 post reactions to the new v0.6.0 post reactions.
// Each migrated post reactions preserve the owner, the value (which is a shortcode) is now contained inside the shortcode field
// while the value field contains the real value of the reaction itself (that can be an emoji or a URI in case of custom reactions).
func migratePostReactions(reactionsMap map[string][]v040posts.PostReaction, regReactions []v040posts.Reaction) map[string][]PostReaction {
	newReactMap := make(map[string][]PostReaction, len(reactionsMap))
	for postID, reactions := range reactionsMap {
		newPostReactions := make([]PostReaction, len(reactions))
		for index, reaction := range reactions {
			newPostReactions[index] = GetConvertedPostReaction(reaction, regReactions)
		}
		newReactMap[postID] = newPostReactions
	}
	return newReactMap
}
