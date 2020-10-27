package types

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewRelationship returns a new relationships with the given recipient and subspace
func NewRelationship(creator string, recipient string, subspace string) Relationship {
	return Relationship{
		Creator:   creator,
		Recipient: recipient,
		Subspace:  subspace,
	}
}

// Validate implement Validator
func (r Relationship) Validate() error {
	if len(r.Creator) == 0 {
		return fmt.Errorf("cretor cannot be empty")
	}

	if len(r.Recipient) == 0 {
		return fmt.Errorf("recipient cannot be empty")
	}

	if r.Creator == r.Recipient {
		return fmt.Errorf("creator and recipient cannot be the same user")
	}

	if !commons.IsValidSubspace(r.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewUserBlock returns a new object representing the fact that one user has blocked another one
// for a specific reason on the given subspace.
func NewUserBlock(blocker, blocked string, reason, subspace string) UserBlock {
	return UserBlock{
		Blocker:  blocker,
		Blocked:  blocked,
		Reason:   reason,
		Subspace: subspace,
	}
}

// Validate implements validator
func (ub UserBlock) Validate() error {
	if len(ub.Blocker) == 0 {
		return fmt.Errorf("blocker address cannot be empty")
	}

	if len(ub.Blocked) == 0 {
		return fmt.Errorf("the address of the blocked user cannot be empty")
	}

	if ub.Blocker == ub.Blocked {
		return fmt.Errorf("blocker and blocked addresses cannot be equals")
	}

	if !commons.IsValidSubspace(ub.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256 hash")
	}

	return nil
}
