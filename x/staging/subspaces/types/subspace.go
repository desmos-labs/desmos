package types

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/commons"
	"time"
)

// NewSubspace returns a Subspace
func NewSubspace(creationTime time.Time, subspaceId, name, creator string) Subspace {
	return Subspace{
		ID:           subspaceId,
		Name:         name,
		Creator:      creator,
		CreationTime: creationTime,
	}
}

func (sub Subspace) Validate() error {
	if !commons.IsValidSubspace(sub.ID) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid sha-256 hash", sub.ID)
	}

	if sub.Creator == "" {
		return fmt.Errorf("invalid subspace creator: %s", sub.Creator)
	}

	return nil
}
