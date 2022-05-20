package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParseSubspaceID parses the given value as a subspace id, returning an error if it's invalid
func ParseSubspaceID(value string) (uint64, error) {
	if value == "" {
		return 0, nil
	}

	subspaceID, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid subspace id: %s", err)
	}
	return subspaceID, nil
}

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

// Validate implements fmt.Validator
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

// Update updates the fields of a given subspace without validating it.
// Before storing the updated subspace, a validation with Validate() should
// be performed.
func (sub Subspace) Update(update SubspaceUpdate) Subspace {
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

// SubspaceUpdate contains all the data that can be updated about a subspace.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type SubspaceUpdate struct {
	Name        string
	Description string
	Treasury    string
	Owner       string
}

// NewSubspaceUpdate builds a new SubspaceUpdate instance containing the given data
func NewSubspaceUpdate(name, description, treasury, owner string) SubspaceUpdate {
	return SubspaceUpdate{
		Name:        name,
		Description: description,
		Treasury:    treasury,
		Owner:       owner,
	}
}

// --------------------------------------------------------------------------------------------------------------------

const (
	// RootSectionID represents the id of the root section of each subspace
	RootSectionID = 0
)

// ParseSectionID parses the given value as a section id, returning an error if it's invalid
func ParseSectionID(value string) (uint32, error) {
	if value == "" {
		return 0, nil
	}

	sectionID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid group id: %s", err)
	}
	return uint32(sectionID), nil
}

// NewSection returns a new Section instance
func NewSection(subspaceID uint64, id uint32, parentID uint32, name string, description string) Section {
	return Section{
		SubspaceID:  subspaceID,
		ID:          id,
		ParentID:    parentID,
		Name:        name,
		Description: description,
	}
}

// DefaultSection returns the default section for the given subspace
func DefaultSection(subspaceID uint64) Section {
	return NewSection(
		subspaceID,
		RootSectionID,
		RootSectionID,
		"Default section",
		"This is the default subspace section",
	)
}

// Validate implements fmt.Validator
func (s Section) Validate() error {
	if s.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", s.SubspaceID)
	}

	if strings.TrimSpace(s.Name) == "" {
		return fmt.Errorf("invalid section name: %s", s.Name)
	}

	return nil
}

// Update updates the fields of a given section without validating it.
// Before storing the updated section, a validation with Validate() should
// be performed.
func (s Section) Update(update SectionUpdate) Section {
	if update.Name == DoNotModify {
		update.Name = s.Name
	}

	if update.Description == DoNotModify {
		update.Description = s.Description
	}

	return NewSection(
		s.SubspaceID,
		s.ID,
		s.ParentID,
		update.Name,
		update.Description,
	)
}

// SectionUpdate contains all the data that can be updated about a section.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type SectionUpdate struct {
	Name        string
	Description string
}

// NewSectionUpdate returns a new SectionUpdate instance
func NewSectionUpdate(name string, description string) SectionUpdate {
	return SectionUpdate{
		Name:        name,
		Description: description,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// ParseGroupID parses the given value as a group id, returning an error if it's invalid
func ParseGroupID(value string) (uint32, error) {
	if value == "" {
		return 0, nil
	}

	groupID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid group id: %s", err)
	}
	return uint32(groupID), nil
}

// NewUserGroup returns a new UserGroup instance
func NewUserGroup(subspaceID uint64, sectionID uint32, id uint32, name, description string, permissions Permission) UserGroup {
	return UserGroup{
		SubspaceID:  subspaceID,
		SectionID:   sectionID,
		ID:          id,
		Name:        name,
		Description: description,
		Permissions: permissions,
	}
}

// DefaultUserGroup returns the default user group for the given subspace
func DefaultUserGroup(subspaceID uint64) UserGroup {
	return NewUserGroup(
		subspaceID,
		0,
		0,
		"Default",
		"This is a default user group which all users are automatically part of",
		PermissionNothing,
	)
}

// Validate implements fmt.Validator
func (group UserGroup) Validate() error {
	if group.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", group.SubspaceID)
	}

	if strings.TrimSpace(group.Name) == "" {
		return fmt.Errorf("invalid group name: %s", group.Name)
	}

	return nil
}

// Update updates the fields of a given group without validating it.
// Before storing the updated group, a validation with Validate() should
// be performed.
func (group UserGroup) Update(update GroupUpdate) UserGroup {
	if update.Name == DoNotModify {
		update.Name = group.Name
	}

	if update.Description == DoNotModify {
		update.Description = group.Description
	}

	return NewUserGroup(
		group.SubspaceID,
		group.SectionID,
		group.ID,
		update.Name,
		update.Description,
		group.Permissions,
	)
}

// GroupUpdate contains all the data that can be updated about a group.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type GroupUpdate struct {
	Name        string
	Description string
}

// NewGroupUpdate builds a new SubspaceUpdate instance containing the given data
func NewGroupUpdate(name, description string) GroupUpdate {
	return GroupUpdate{
		Name:        name,
		Description: description,
	}
}
