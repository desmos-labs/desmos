package v050

import (
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
)

// Migrate accepts exported genesis state from v0.4.0 and migrates it to v0.5.0
// genesis state. This migration removed the Unicode emoji that was registered as reactions in v0.4.0
func Migrate(oldGenState v040posts.GenesisState) v040posts.GenesisState {
	return v040posts.GenesisState{
		Posts:               oldGenState.Posts,
		UsersPollAnswers:    oldGenState.UsersPollAnswers,
		PostReactions:       oldGenState.PostReactions,
		RegisteredReactions: GetReactionsToRegister(oldGenState.RegisteredReactions),
	}
}

// GetReactionsToRegister takes the list of reactions that were registered in v0.4.0 and
// returns a new list without the registered Unicode emojis.
func GetReactionsToRegister(oldRegisteredReactions []v040posts.Reaction) (reactionsToRegister []v040posts.Reaction) {
	for _, reaction := range oldRegisteredReactions {
		if _, err := emoji.LookupEmojiByCode(reaction.ShortCode); err != nil {
			reactionsToRegister = append(reactionsToRegister, reaction)
		}
	}
	return reactionsToRegister
}
