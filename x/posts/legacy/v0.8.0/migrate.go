package v080

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
)

// Migrate accepts exported genesis state from v0.6.0 and migrates it to v0.8.0
// genesis state. This migration replace all the old post reactions structure
// with the new one that includes shortcodes.
func Migrate(oldGenState v060posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               oldGenState.Posts,
		UsersPollAnswers:    oldGenState.UsersPollAnswers,
		PostReactions:       oldGenState.PostReactions,
		RegisteredReactions: RemoveInvalidEmojiRegisteredReactions(oldGenState.RegisteredReactions),
		Params: Params{
			MaxPostMessageLength:            sdk.NewInt(500),
			MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
			MaxOptionalDataFieldValueLength: sdk.NewInt(200),
		},
	}
}

// RemoveInvalidEmojiRegisteredReactions removes all the invalid registered reactions.
// This removes all the ones that have either a shortcode which is associated to an emoji,
// or a value that is not a URL.
func RemoveInvalidEmojiRegisteredReactions(reactions []v040posts.Reaction) []v040posts.Reaction {
	var newReactions []v040posts.Reaction
	for _, reaction := range reactions {
		_, err := emoji.LookupEmojiByCode(reaction.ShortCode)
		if URIRegEx.MatchString(reaction.Value) && err != nil {
			newReactions = append(newReactions, reaction)
		}
	}
	return newReactions
}
