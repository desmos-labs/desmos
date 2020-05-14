package reactions

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
)

// ---------------
// --- PostReaction
// ---------------

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	Shortcode string         `json:"shortcode" yaml:"shortcode"` // Shortcode of the reaction
	Value     string         `json:"value" yaml:"value"`         // Value of the reaction
	Owner     sdk.AccAddress `json:"owner" yaml:"owner"`         // Creator that has created the reaction
}

// NewPostReaction returns a new PostReaction
func NewPostReaction(shortcode, value string, owner sdk.AccAddress) PostReaction {
	return PostReaction{
		Shortcode: shortcode,
		Value:     value,
		Owner:     owner,
	}
}

// String implements fmt.Stringer
func (reaction PostReaction) String() string {
	bytes, err := json.Marshal(&reaction)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Validate implements validator
func (reaction PostReaction) Validate() error {
	if reaction.Owner.Empty() {
		return fmt.Errorf("invalid reaction owner: %s", reaction.Owner)
	}

	if len(strings.TrimSpace(reaction.Value)) == 0 {
		return fmt.Errorf("reaction value cannot be empty or blank")
	}

	if !common.ShortCodeRegEx.MatchString(reaction.Shortcode) {
		//nolint - errcheck
		return fmt.Errorf("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a :")
	}

	return nil
}

// Equals returns true if reaction and other contain the same data
func (reaction PostReaction) Equals(other PostReaction) bool {
	return reaction.Value == other.Value &&
		reaction.Shortcode == other.Shortcode &&
		reaction.Owner.Equals(other.Owner)
}

// ------------
// --- PostReactions
// ------------

// PostReactions represents a slice of PostReaction objects
type PostReactions []PostReaction

// NewPostReactions allows to create a new PostReactions object from the given reactions
func NewPostReactions(reactions ...PostReaction) PostReactions {
	return reactions
}

// AppendIfMissing returns a new slice of PostReaction objects containing
// the given reaction if it wasn't already present.
// It also returns the result of the append.
func (reactions PostReactions) AppendIfMissing(other PostReaction) (PostReactions, bool) {
	for _, reaction := range reactions {
		if reaction.Equals(other) {
			return reactions, false
		}
	}
	return append(reactions, other), true
}

// ContainsReactionFrom returns true if the reactions slice contain
// a reaction from the given user having the given value, false otherwise.
// NOTE: The value can be either an emoji or a shortcode.
func (reactions PostReactions) ContainsReactionFrom(user sdk.Address, value string) bool {
	return reactions.IndexOfByUserAndValue(user, value) != -1
}

// IndexOfByUserAndValue returns the index of the reaction from the
// given user with the specified code inside the reactions slice.
// NOTE: The value can be either an emoji or a shortcode.
func (reactions PostReactions) IndexOfByUserAndValue(owner sdk.Address, value string) int {
	var reactEmoji *emoji.Emoji
	if ej, err := emoji.LookupEmoji(value); err == nil {
		reactEmoji = &ej
	} else if ej, err := emoji.LookupEmojiByCode(value); err == nil {
		reactEmoji = &ej
	}

	for index, reaction := range reactions {
		if reaction.Owner.Equals(owner) {
			if reactEmoji != nil {
				// Check the emoji value
				if reaction.Value == reactEmoji.Value {
					return index
				}

				// Check the emoji shortcodes
				for _, code := range reactEmoji.Shortcodes {
					if reaction.Shortcode == code {
						return index
					}
				}
			}

			if reactEmoji == nil {
				if value == reaction.Shortcode {
					return index
				}
			}
		}
	}
	return -1
}

// RemovePostReaction returns a new PostReactions slice not containing the
// reaction of the given user with the given value.
// If the reaction was removed properly, true is also returned. Otherwise,
// if no reaction was found, false is returned instead.
func (reactions PostReactions) RemoveReaction(user sdk.Address, value string) (PostReactions, bool) {
	index := reactions.IndexOfByUserAndValue(user, value)
	if index == -1 {
		return reactions, false
	}

	return append(reactions[:index], reactions[index+1:]...), true
}
