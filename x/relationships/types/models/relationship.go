package models

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
)

// Relationship is the struct of a relationship.
// It represent the concept of "follow" of traditional social networks.
type Relationship struct {
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Subspace  string         `json:"subspace" yaml:"subspace"`
}

func NewRelationship(recipient sdk.AccAddress, subspace string) Relationship {
	return Relationship{
		Recipient: recipient,
		Subspace:  subspace,
	}
}

// String implement fmt.Stringer
func (r Relationship) String() string {
	out := "Relationship:"
	out += fmt.Sprintf("[Recipient] %s [Subspace] %s", r.Recipient, r.Subspace)
	return out
}

// Validate implement Validator
func (r Relationship) Validate() error {
	if r.Recipient.Empty() {
		return fmt.Errorf("recipient can't be empty")
	}

	// TODO edit this when user blocks is merged
	if !posts.IsValidSubspace(r.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256")
	}

	return nil
}

// Equals allows to check whether the contents of r are the same of other
func (r Relationship) Equals(other Relationship) bool {
	return r.Recipient.Equals(other.Recipient) &&
		r.Subspace == other.Subspace
}

type Relationships []Relationship

// String implement fmt.Stringer
func (rs Relationships) String() string {
	out := "Relationships:\n"
	for _, rel := range rs {
		out += fmt.Sprintf("%s\n", rel.String())
	}
	return out
}
