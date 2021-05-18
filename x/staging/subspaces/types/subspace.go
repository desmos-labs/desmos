package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/desmos-labs/desmos/x/commons"
)

const (
	Admin          = "admin"
	RegisteredUser = "registered-user"
	BlockedUser    = "blocked-user"
)

// NewSubspace is a constructor for the Subspace type
func NewSubspace(subspaceID, name, owner, creator string, open bool, creationTime time.Time) Subspace {
	return Subspace{
		ID:           subspaceID,
		Name:         name,
		Owner:        owner,
		Creator:      creator,
		CreationTime: creationTime,
		Open:         open,
	}
}

func (sub Subspace) WithName(name string) Subspace {
	if strings.TrimSpace(name) != "" {
		sub.Name = name
	}
	return sub
}

func (sub Subspace) WithOwner(owner string) Subspace {
	if strings.TrimSpace(owner) != "" {
		sub.Owner = owner
	}
	return sub
}

func (sub Subspace) Validate() error {
	if !commons.IsValidSubspace(sub.ID) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid sha-256 hash", sub.ID)
	}

	if strings.TrimSpace(sub.Name) == "" {
		return fmt.Errorf("subspace name cannot be empty or blank")
	}

	if sub.Owner == "" {
		return fmt.Errorf("invalid subspace owner: %s", sub.Owner)
	}

	if sub.Creator == "" {
		return fmt.Errorf("invalid subspace creator: %s", sub.Creator)
	}

	if sub.CreationTime.IsZero() {
		return fmt.Errorf("invalid subspace creation time: %s", sub.CreationTime)
	}

	if err := sub.Admins.ValidateUsers(Admin); err != nil {
		return err
	}

	if err := sub.BlockedUsers.ValidateUsers(BlockedUser); err != nil {
		return err
	}

	if err := sub.RegisteredUsers.ValidateUsers(RegisteredUser); err != nil {
		return err
	}

	return nil
}
