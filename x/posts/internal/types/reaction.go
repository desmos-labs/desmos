package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Reaction represents a registered reaction that can be referenced
// by its shortCode inside post reactions
type Reaction struct {
	ShortCode string
	Value     string
	Subspace  string
	Creator   sdk.AccAddress
}

// NewPostReaction returns a new PostReaction
func NewReaction(creator sdk.AccAddress, shortCode, value, subspace string) Reaction {
	return Reaction{
		ShortCode: shortCode,
		Value:     value,
		Subspace:  subspace,
		Creator:   creator,
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
	if reaction.Creator.Empty() {
		return fmt.Errorf("invalid reaction creator: %s", reaction.Creator)
	}

	if len(strings.TrimSpace(reaction.ShortCode)) == 0 {
		return fmt.Errorf("reaction short code cannot be empty or blank")
	}

	if len(strings.TrimSpace(reaction.Value)) == 0 {
		return fmt.Errorf("reaction value cannot be empty or blank")
	}

	if !ShortCodeRegEx.MatchString(reaction.ShortCode) {
		return fmt.Errorf("reaction short code must be an emoji short code")
	}

	if !URIRegEx.MatchString(reaction.Value) || !UnicodeRegEx.MatchString(reaction.Value) {
		return fmt.Errorf("reaction value should be a URL or an emoji unicode")
	}

	if !SubspaceRegEx.MatchString(reaction.Subspace) {
		return fmt.Errorf("reaction subspace must be a valid sha-256 hash")
	}

	return nil
}

// Equals returns true if reaction and other contain the same data
func (reaction Reaction) Equals(other Reaction) bool {
	return reaction.Value == other.Value &&
		reaction.ShortCode == other.ShortCode &&
		reaction.Subspace == other.Subspace &&
		reaction.Creator.Equals(other.Creator)
}

// ------------
// --- Reactions
// ------------

// PostReactions represents a slice of Reaction objects
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
