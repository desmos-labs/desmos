package v2

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
func NewUserGroup(subspaceID uint64, id uint32, name, description string, permissions Permission) UserGroup {
	return UserGroup{
		SubspaceID:  subspaceID,
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
		"Default",
		"This is a default user group which all users are automatically part of",
		PermissionNothing,
	)
}

// Validate returns an error if something is wrong within the group data
func (group UserGroup) Validate() error {
	if group.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", group.SubspaceID)
	}

	if strings.TrimSpace(group.Name) == "" {
		return fmt.Errorf("invalid group name: %s", group.Name)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// GroupUpdate contains all the data that can be updated about a group.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type GroupUpdate struct {
	Name        string
	Description string
}

// NewGroupUpdate builds a new SubspaceUpdate instance containing the given data
func NewGroupUpdate(name, description string) *GroupUpdate {
	return &GroupUpdate{
		Name:        name,
		Description: description,
	}
}

// Update updates the fields of a given group without validating it.
// Before storing the updated group, a validation with Validate() should
// be performed.
func (group UserGroup) Update(update *GroupUpdate) UserGroup {
	if update.Name == DoNotModify {
		update.Name = group.Name
	}

	if update.Description == DoNotModify {
		update.Description = group.Description
	}

	return NewUserGroup(
		group.SubspaceID,
		group.ID,
		update.Name,
		update.Description,
		group.Permissions,
	)
}

// --------------------------------------------------------------------------------------------------------------------

// NewPermissionDetailUser returns a new PermissionDetail for the user with the given address and permission value
func NewPermissionDetailUser(user string, permission Permission) PermissionDetail {
	return PermissionDetail{
		Sum: &PermissionDetail_User_{
			User: &PermissionDetail_User{
				User:       user,
				Permission: permission,
			},
		},
	}
}

// NewPermissionDetailGroup returns a new PermissionDetail for the user with the given id and permission value
func NewPermissionDetailGroup(groupID uint32, permission Permission) PermissionDetail {
	return PermissionDetail{
		Sum: &PermissionDetail_Group_{
			Group: &PermissionDetail_Group{
				GroupID:    groupID,
				Permission: permission,
			},
		},
	}
}
