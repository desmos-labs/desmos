package types

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/commons"
	"time"
)

// NewSubspace returns a Subspace
func NewSubspace(creationTime time.Time, subspaceId, creator string) Subspace {
	return Subspace{
		Id:           subspaceId,
		Creator:      creator,
		CreationTime: creationTime,
	}
}

func (sub Subspace) Validate() error {
	if !commons.IsValidSubspace(sub.Id) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid sha-256 hash", sub.Id)
	}

	if sub.Creator == "" {
		return fmt.Errorf("invalid post owner: %s", sub.Creator)
	}

	return nil
}
