package reactions

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
)

// Reaction represents a registered reaction that can be referenced
// by its shortCode inside post reactions
type Reaction struct {
	ShortCode string         `json:"shortcode" yaml:"shortcode"`
	Value     string         `json:"value" yaml:"value"`
	Subspace  string         `json:"subspace" yaml:"subspace"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewReaction returns a new Reaction
func NewReaction(creator sdk.AccAddress, shortCode, value, subspace string) Reaction {
	return Reaction{
		ShortCode: shortCode,
		Value:     value,
		Subspace:  subspace,
		Creator:   creator,
	}
}

// Validate implements validator
func (reaction Reaction) Validate() error {
	if reaction.Creator.Empty() {
		return fmt.Errorf("invalid reaction creator: %s", reaction.Creator)
	}

	if !common.ShortCodeRegEx.MatchString(reaction.ShortCode) {
		return fmt.Errorf("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'")
	}

	if !common.URIRegEx.MatchString(reaction.Value) {
		return fmt.Errorf("reaction value should be a URL")
	}

	if !common.Sha256RegEx.MatchString(reaction.Subspace) {
		return fmt.Errorf("reaction subspace must be a valid sha-256 hash")
	}

	if _, found := common.GetEmojiByShortCodeOrValue(reaction.ShortCode); found {
		return fmt.Errorf("reaction has emoji shortcode: %s", reaction.ShortCode)
	}

	return nil
}

// IsEmoji checks whether the value is an emoji or an emoji unicode
func IsEmoji(value string) bool {

	_, err := emoji.LookupEmoji(value)
	if err == nil {
		return true
	}

	trimmed := strings.TrimPrefix(value, "U+")
	emo := emoji.Emojis[trimmed]
	return len(emo.Key) != 0
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

// Reactions represents a slice of Reaction objects
type Reactions []Reaction

// NewReactions allows to create a Reactions object given a list of reactions
func NewReactions(reactions ...Reaction) Reactions {
	return reactions
}

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
