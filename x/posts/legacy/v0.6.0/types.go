package v060

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
)

// GenesisState contains the data of a v0.6.0 genesis state for the posts module
type GenesisState struct {
	Posts               []v040posts.Post                  `json:"posts"`
	UsersPollAnswers    map[string][]v040posts.UserAnswer `json:"users_poll_answers"`
	PostReactions       map[string][]PostReaction         `json:"post_reactions"`
	RegisteredReactions []v040posts.Reaction              `json:"registered_reactions"`
}

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	Owner     sdk.AccAddress `json:"owner" yaml:"owner"`         // Creator that has created the reaction
	Shortcode string         `json:"shortcode" yaml:"shortcode"` // Shortcode of the reaction
	Value     string         `json:"value" yaml:"value"`         // Value of the reaction
}
