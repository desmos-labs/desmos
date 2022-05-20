package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisSubspace returns a new GenesisSubspace instance
func NewGenesisSubspace(subspace Subspace, initialGroupID uint32) GenesisSubspace {
	return GenesisSubspace{
		Subspace:       subspace,
		InitialGroupID: initialGroupID,
	}
}

// Validate returns an error if something is wrong within the subspace data
func (subspace GenesisSubspace) Validate() error {
	if subspace.InitialGroupID == 0 {
		return fmt.Errorf("invalid initial group id: %d", subspace.InitialGroupID)
	}

	return subspace.Subspace.Validate()
}

// -------------------------------------------------------------------------------------------------------------------

// NewUserGroupMembersEntry returns a new UserGroupMembersEntry instance
func NewUserGroupMembersEntry(subspaceID uint64, groupID uint32, members []string) UserGroupMembersEntry {
	return UserGroupMembersEntry{
		SubspaceID: subspaceID,
		GroupID:    groupID,
		Members:    members,
	}
}

// Validate returns an error if something is wrong within the entry data
func (entry UserGroupMembersEntry) Validate() error {
	if entry.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", entry.SubspaceID)
	}

	if entry.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", entry.GroupID)
	}

	for _, user := range entry.Members {
		_, err := sdk.AccAddressFromBech32(user)
		if err != nil {
			return fmt.Errorf("invalid user address: %s", user)
		}
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// NewGenesisState creates a new genesis state
func NewGenesisState(
	initialSubspaceID uint64, subspaces []GenesisSubspace, userPermissions []UserPermission,
	userGroups []UserGroup, userGroupMembers []UserGroupMembersEntry,
) *GenesisState {
	return &GenesisState{
		InitialSubspaceID: initialSubspaceID,
		Subspaces:         subspaces,
		UserPermissions:   userPermissions,
		UserGroups:        userGroups,
		UserGroupsMembers: userGroupMembers,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(1, nil, nil, nil, nil)
}

// -------------------------------------------------------------------------------------------------------------------

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	// Make sure the initial subspace id is valid
	if data.InitialSubspaceID <= uint64(len(data.Subspaces)) {
		return fmt.Errorf("invalid initial subspace id: %d", data.InitialSubspaceID)
	}

	// Validate the subspace
	for _, subspace := range data.Subspaces {
		err := subspace.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedSubspace(data.Subspaces, subspace) {
			return fmt.Errorf("duplicated subspace: %d", subspace.Subspace.ID)
		}
	}

	// Validate the ACL entries
	for _, entry := range data.UserPermissions {
		err := entry.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedUserPermission(data.UserPermissions, entry) {
			return fmt.Errorf("duplicated ACL entry for subspace %d and user %s", entry.SubspaceID, entry.User)
		}

		// Make sure the associated subspace exists
		subspace, found := findSubspace(data.Subspaces, entry.SubspaceID)
		if !found {
			return fmt.Errorf("invalid ACL entry: subspace %d not found", subspace.Subspace.ID)
		}
	}

	// Validate the user groups
	groupsCount := map[uint64]int{}
	for _, group := range data.UserGroups {
		err := group.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedGroups(data.UserGroups, group) {
			return fmt.Errorf("duplicated group for subspace %d and group %d", group.SubspaceID, group.ID)
		}

		// Increment the groups count for this subspace
		groupsCount[group.SubspaceID]++
	}

	// Make sure each subspace has a correct initial group id based on the number of groups inside that subspace
	for subspaceID, count := range groupsCount {
		genSub, found := findSubspace(data.Subspaces, subspaceID)
		if !found {
			return fmt.Errorf("invalid group id: subspace %d not found", subspaceID)
		}

		if genSub.InitialGroupID <= uint32(count) {
			return fmt.Errorf("invalid initial group id for subspace %d: %d", genSub.Subspace.ID, genSub.InitialGroupID)
		}
	}

	// Validate the group members
	for _, entry := range data.UserGroupsMembers {
		err := entry.Validate()
		if err != nil {
			return err
		}

		if containsDuplicatedMembersEntries(data.UserGroupsMembers, entry) {
			return fmt.Errorf("duplicated user group members entry for group %d within subspace %d", entry.GroupID, entry.SubspaceID)
		}

		// Make sure the associated subspace exists
		_, found := findGroup(data.UserGroups, entry.SubspaceID, entry.GroupID)
		if !found {
			return fmt.Errorf("invalid group members entry: group %d for subspace %d not found",
				entry.GroupID, entry.SubspaceID)
		}
	}

	return nil
}

// findSubspace searches the subspace with the given id inside the provided slice
func findSubspace(subspaces []GenesisSubspace, subspaceID uint64) (genSub GenesisSubspace, found bool) {
	for _, subspace := range subspaces {
		if subspace.Subspace.ID == subspaceID {
			return subspace, true
		}
	}
	return GenesisSubspace{}, false
}

// findGroup searches the group for the group having the given id and subspace id inside the given slice
func findGroup(groups []UserGroup, subspaceID uint64, groupID uint32) (group UserGroup, found bool) {
	for _, group := range groups {
		if group.SubspaceID == subspaceID && group.ID == groupID {
			return group, true
		}
	}
	return UserGroup{}, false
}

// containsDuplicatedSubspace tells whether the given subspaces slice contains two or more
// subspaces with the same id of the given subspace
func containsDuplicatedSubspace(subspaces []GenesisSubspace, subspace GenesisSubspace) bool {
	var count = 0
	for _, s := range subspaces {
		if s.Subspace.ID == subspace.Subspace.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedUserPermission tells whether the given entries slice contains two or more
// entries for the same user and subspace
func containsDuplicatedUserPermission(entries []UserPermission, entry UserPermission) bool {
	var count = 0
	for _, e := range entries {
		if e.SubspaceID == entry.SubspaceID && e.User == entry.User {
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
		if g.SubspaceID == group.SubspaceID && g.ID == group.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedMembersEntries tells whether the given entries slice contains two or more
// entries for the same subspace and group id
func containsDuplicatedMembersEntries(entries []UserGroupMembersEntry, entry UserGroupMembersEntry) bool {
	var count = 0
	for _, e := range entries {
		if e.SubspaceID == entry.SubspaceID && e.GroupID == entry.GroupID {
			count++
		}
	}
	return count > 1
}
