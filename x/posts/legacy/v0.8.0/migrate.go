package v060

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		RegisteredReactions: oldGenState.RegisteredReactions,
		Params: Params{
			MaxPostMessageLength:            sdk.NewInt(500),
			MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
			MaxOptionalDataFieldValueLength: sdk.NewInt(200),
		},
	}
}
