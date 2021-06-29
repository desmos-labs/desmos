package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewSubspace is a constructor for the Subspace type
func NewSubspace(subspaceID, name, description, logo, owner, creator string, subspaceType SubspaceType, creationTime time.Time) Subspace {
	return Subspace{
		ID:           subspaceID,
		Name:         name,
		Description:  description,
		Logo:         logo,
		Owner:        owner,
		Creator:      creator,
		CreationTime: creationTime,
		Type:         subspaceType,
	}
}

// WithName is a decorator that will replace the subspace name with a new one
func (sub Subspace) WithName(name string) Subspace {
	if name != DoNotModify {
		sub.Name = name
	}
	return sub
}

// WithOwner is a decorator that will replace the subspace owner with a new one
func (sub Subspace) WithOwner(owner string) Subspace {
	if strings.TrimSpace(owner) != "" {
		sub.Owner = owner
	}
	return sub
}

// WithSubspaceType is a decorator that will replace the subspace type with a new one
func (sub Subspace) WithSubspaceType(subspaceType SubspaceType) Subspace {
	if subspaceType != SubspaceTypeUnspecified {
		sub.Type = subspaceType
	}
	return sub
}

// WithDescription is a decorator that will replace the subspace description with a new one
func (sub Subspace) WithDescription(description string) Subspace {
	if description != DoNotModify {
		sub.Description = description
	}
	return sub
}

// WithLogo is a decorator that will replace the subspace logo
func (sub Subspace) WithLogo(logo string) Subspace {
	if logo != DoNotModify {
		sub.Logo = logo
	}
	return sub
}

// Validate will perform some checks to ensure the subspace validity
func (sub Subspace) Validate() error {
	if !IsValidSubspace(sub.ID) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid SHA-256 hash", sub.ID)
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

	if !IsValidSubspaceType(sub.Type) {
		return fmt.Errorf("invalid subspace type: %s", sub.Type)
	}

	if strings.TrimSpace(sub.Logo) != "" && !commons.IsURIValid(sub.Logo) {
		return fmt.Errorf("invalid subspace logo uri provided")
	}

	return nil
}

// SubspaceTypeFromString convert a string in the corresponding SubspaceType
func SubspaceTypeFromString(subType string) (SubspaceType, error) {
	subspaceType, ok := SubspaceType_value[subType]
	if !ok {
		return SubspaceTypeUnspecified, fmt.Errorf("'%s' is not a valid subspace type", subType)
	}
	return SubspaceType(subspaceType), nil
}

// NormalizeSubspaceType - normalize user specified subspace type
func NormalizeSubspaceType(subType string) string {
	switch strings.ToLower(subType) {
	case "open":
		return SubspaceTypeOpen.String()
	case "close":
		return SubspaceTypeClosed.String()
	default:
		return subType
	}
}

// IsValidSubspaceType checks if the subspaceType given correspond to one of the valid ones
func IsValidSubspaceType(subspaceType SubspaceType) bool {
	if subspaceType == SubspaceTypeOpen || subspaceType == SubspaceTypeClosed {
		return true
	}
	return false
}
