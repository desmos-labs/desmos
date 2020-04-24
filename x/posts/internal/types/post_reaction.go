package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// ---------------
// --- PostReaction
// ---------------

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	Owner sdk.AccAddress `json:"owner"` // Creator that has created the reaction
	Value string         `json:"value"` // Value of the reaction, either an emoji or a shortcode
}

// NewPostReaction returns a new PostReaction
func NewPostReaction(value string, owner sdk.AccAddress) PostReaction {
	return PostReaction{
		Value: value,
		Owner: owner,
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

	return nil
}

// Equals returns true if reaction and other contain the same data
func (reaction PostReaction) Equals(other PostReaction) bool {
	return reaction.Value == other.Value &&
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
	reactEmoji, err := emoji.LookupEmoji(value)
	isEmoji := err == nil

	for index, reaction := range reactions {
		if reaction.Owner.Equals(owner) {
			// The given value is a shortcode, so check only that
			if !isEmoji && reaction.Value == value {
				return index
			}

			// The given value is an emoji, so we need to check is any of its shortcode match this reaction value
			if isEmoji {
				for _, code := range reactEmoji.Shortcodes {
					if reaction.Value == code {
						return index
					}
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
