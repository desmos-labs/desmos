package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Reaction represents a registered reaction that can be referenced
// by its shortcode inside post reactions
type Reaction struct {
	ShortCode string
	Value     string
}

// NewPostReaction returns a new PostReaction
func NewReaction(shortCode, value string) Reaction {
	return Reaction{
		ShortCode: shortCode,
		Value:     value,
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
	if len(strings.TrimSpace(reaction.ShortCode)) == 0 {
		return fmt.Errorf("reaction shortCode cannot empty or blank")
	}

	if len(strings.TrimSpace(reaction.Value)) == 0 {
		return fmt.Errorf("reaction value cannot empty or blank")
	}

	return nil
}

// Equals returns true if reaction and other contain the same data
func (reaction Reaction) Equals(other Reaction) bool {
	return reaction.Value == other.Value &&
		reaction.ShortCode == other.ShortCode
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
