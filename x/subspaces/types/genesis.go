package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewACLEntry returns a new ACLEntry instance
func NewACLEntry(subspaceID uint64, target string, permissions Permission) ACLEntry {
	return ACLEntry{
		SubspaceID:  subspaceID,
		Target:      target,
		Permissions: permissions,
	}
}

// Validate returns an error if something is wrong within the entry data
func (entry ACLEntry) Validate() error {
	if entry.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", entry.SubspaceID)
	}

	return nil
}

// NewUserGroup returns a new UserGroup instance
func NewUserGroup(subspaceID uint64, groupName string, members []string) UserGroup {
	return UserGroup{
		SubspaceID: subspaceID,
		Name:       groupName,
		Members:    members,
	}
}

// Validate returns an error if something is wrong within the group data
func (group UserGroup) Validate() error {
	if group.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", group.SubspaceID)
	}

	if group.Name == "" {
		return fmt.Errorf("invalid group name: %s", group.Name)
	}

	for _, member := range group.Members {
		_, err := sdk.AccAddressFromBech32(member)
		if err != nil {
			return fmt.Errorf("invalid group member address: %s", member)
		}
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// NewGenesisState creates a new genesis state
func NewGenesisState(subspaces []Subspace, acl []ACLEntry, userGroups []UserGroup) *GenesisState {
	return &GenesisState{
		Subspaces:  subspaces,
		ACL:        acl,
		UserGroups: userGroups,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, subspace := range data.Subspaces {
		err := subspace.Validate()
		if err != nil {
			return err
		}
	}

	for _, subspace := range data.Subspaces {
		err := subspace.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedSubspace(data.Subspaces, subspace) {
			return fmt.Errorf("duplicated subspace: %d", subspace.ID)
		}
	}

	for _, entry := range data.ACL {
		err := entry.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedACLEntry(data.ACL, entry) {
			return fmt.Errorf("duplicated ACL entry for subspace %d: %s", entry.SubspaceID, entry.Target)
		}
	}

	for _, group := range data.UserGroups {
		err := group.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedGroups(data.UserGroups, group) {
			return fmt.Errorf("duplicated group for subspace %d: %s", group.SubspaceID, group.Name)
		}
	}

	return nil
}

// containsDuplicatedSubspace tells whether the given subspaces slice contains two or more
// subspaces with the same id of the given subspace
func containsDuplicatedSubspace(subspaces []Subspace, subspace Subspace) bool {
	var count = 0
	for _, s := range subspaces {
		if s.ID == subspace.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedACLEntry tells whether the given entries slice contains two or more
// entries for the same target and subspace
func containsDuplicatedACLEntry(entries []ACLEntry, entry ACLEntry) bool {
	var count = 0
	for _, e := range entries {
		if e.SubspaceID == entry.SubspaceID && e.Target == entry.Target {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedGroups tells whether the given groups slice contains two or more
// groups for the same subspace having the same name
func containsDuplicatedGroups(groups []UserGroup, group UserGroup) bool {
	var count = 0
	for _, g := range groups {
		if g.SubspaceID == group.SubspaceID && g.Name == group.Name {
			count++
		}
	}
	return count > 1
}
