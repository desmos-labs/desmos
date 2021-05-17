package types

import (
	"fmt"
	"strings"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
)

// NewSubspace returns a Subspace
func NewSubspace(creationTime time.Time, subspaceID, name, creator string) Subspace {
	return Subspace{
		ID:           subspaceID,
		Name:         name,
		Owner:        creator,
		CreationTime: creationTime,
	}
}

func (sub Subspace) Validate() error {
	if !commons.IsValidSubspace(sub.ID) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid sha-256 hash", sub.ID)
	}

	if strings.TrimSpace(sub.Name) == "" {
		return sdkerrors.Wrap(ErrInvalidSubspaceName, "subspace name cannot be empty or blank")
	}

	if sub.Owner == "" {
		return fmt.Errorf("invalid subspace owner: %s", sub.Owner)
	}

	return nil
}
