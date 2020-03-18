package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- PostReaction
// ---------------

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	Owner sdk.AccAddress `json:"owner"` // User that has created the reaction
	Value string         `json:"value"` // PostReaction of the reaction
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
// a reaction from the given user having the given value, false otherwise
func (reactions PostReactions) ContainsReactionFrom(user sdk.Address, value string) bool {
	return reactions.IndexOfByUserAndValue(user, value) != -1
}

// IndexOfByUserAndValue returns the index of the reaction from the
// given user with the specified value inside the reactions slice.
func (reactions PostReactions) IndexOfByUserAndValue(owner sdk.Address, value string) int {
	for index, reaction := range reactions {
		if reaction.Owner.Equals(owner) && reaction.Value == value {
			return index
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
