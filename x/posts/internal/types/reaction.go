package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Reaction
// ---------------

// Reaction is a struct of a user reaction to a post
type Reaction struct {
	Created sdk.Int        `json:"created"` // Block height at which the reaction was created
	Owner   sdk.AccAddress `json:"owner"`   // User that has created the reaction
	Value   string         `json:"value"`   // Value of the reaction
}

// NewReaction returns a new Reaction
func NewReaction(value string, created int64, owner sdk.AccAddress) Reaction {
	return Reaction{
		Value:   value,
		Created: sdk.NewInt(created),
		Owner:   owner,
	}
}

// String implements fmt.Stringer
func (reaction Reaction) String() string {
	bytes, err := json.Marshal(&reaction)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Validate implements validator
func (reaction Reaction) Validate() error {
	if reaction.Owner.Empty() {
		return fmt.Errorf("invalid reaction owner: %s", reaction.Owner)
	}

	if reaction.Created.Equal(sdk.ZeroInt()) {
		return fmt.Errorf("invalid reaction creation block height: %s", reaction.Created)
	}

	if len(strings.TrimSpace(reaction.Value)) == 0 {
		return fmt.Errorf("reaction value cannot empty or blank")
	}

	return nil
}

// Equals returns true if reaction and other contain the same data
func (reaction Reaction) Equals(other Reaction) bool {
	return reaction.Value == other.Value &&
		reaction.Created == other.Created &&
		reaction.Owner.Equals(other.Owner)
}

// ------------
// --- Reactions
// ------------

// Reactions represents a slice of Reaction objects
type Reactions []Reaction

// AppendIfMissing returns a new slice of Reaction objects containing
// the given reaction if it wasn't already present.
// It also returns the result of the append.
func (reactions Reactions) AppendIfMissing(other Reaction) (Reactions, bool) {
	for _, reaction := range reactions {
		if reaction.Equals(other) {
			return reactions, false
		}
	}
	return append(reactions, other), true
}

// ContainsReactionFrom returns true if the reactions slice contain
// a reaction from the given user having the given value, false otherwise
func (reactions Reactions) ContainsReactionFrom(user sdk.Address, value string) bool {
	return reactions.IndexOfByUserAndValue(user, value) != -1
}

// IndexOfByUserAndValue returns the index of the reaction from the
// given user with the specified value inside the reactions slice.
func (reactions Reactions) IndexOfByUserAndValue(owner sdk.Address, value string) int {
	for index, reaction := range reactions {
		if reaction.Owner.Equals(owner) && reaction.Value == value {
			return index
		}
	}
	return -1
}

// RemoveReaction returns a new Reactions slice not containing the
// reaction of the given user with the given value.
// If the reaction was removed properly, true is also returned. Otherwise,
// if no reaction was found, false is returned instead.
func (reactions Reactions) RemoveReaction(user sdk.Address, value string) (Reactions, bool) {
	index := reactions.IndexOfByUserAndValue(user, value)
	if index == -1 {
		return reactions, false
	}

	return append(reactions[:index], reactions[index+1:]...), true
}
