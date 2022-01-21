package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewSubspace is a constructor for the Subspace type
func NewSubspace(subspaceID uint64, name, description, treasury, owner, creator string, creationTime time.Time) Subspace {
	return Subspace{
		ID:           subspaceID,
		Name:         name,
		Description:  description,
		Treasury:     treasury,
		Owner:        owner,
		Creator:      creator,
		CreationTime: creationTime,
	}
}

// Validate will perform some checks to ensure the subspace validity
func (sub Subspace) Validate() error {
	if sub.ID == 0 {
		return fmt.Errorf("invalid subspace id: %d", sub.ID)
	}

	if strings.TrimSpace(sub.Name) == "" {
		return fmt.Errorf("subspace name cannot be empty or blank")
	}

	if sub.Treasury != "" {
		_, err := sdk.AccAddressFromBech32(sub.Treasury)
		if err != nil {
			return fmt.Errorf("invalid treasury address: %s", sub.Treasury)
		}
	}

	_, err := sdk.AccAddressFromBech32(sub.Owner)
	if err != nil {
		return fmt.Errorf("invalid owner address: %s", sub.Owner)
	}

	_, err = sdk.AccAddressFromBech32(sub.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %s", sub.Creator)
	}

	if sub.CreationTime.IsZero() {
		return fmt.Errorf("invalid subspace creation time: %s", sub.CreationTime)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// SubspaceUpdate contains all the data that can be updated about a subspace.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type SubspaceUpdate struct {
	Name        string
	Description string
	Treasury    string
	Owner       string
}

// NewSubspaceUpdate builds a new SubspaceUpdate instance containing the given data
func NewSubspaceUpdate(name, description, treasury, owner string) *SubspaceUpdate {
	return &SubspaceUpdate{
		Name:        name,
		Description: description,
		Treasury:    treasury,
		Owner:       owner,
	}
}

// Update updates the fields of a given subspace without validating it.
// Before storing the updated subspace, a validation with Validate() should
// be performed.
func (sub Subspace) Update(update *SubspaceUpdate) Subspace {
	if update.Name == DoNotModify {
		update.Name = sub.Name
	}

	if update.Description == DoNotModify {
		update.Description = sub.Description
	}

	if update.Treasury == DoNotModify {
		update.Treasury = sub.Treasury
	}

	if update.Owner == DoNotModify {
		update.Owner = sub.Owner
	}

	return NewSubspace(
		sub.ID,
		update.Name,
		update.Description,
		update.Treasury,
		update.Owner,
		sub.Creator,
		sub.CreationTime,
	)
}
