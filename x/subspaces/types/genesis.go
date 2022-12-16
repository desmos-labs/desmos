package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new genesis state
func NewGenesisState(
	initialSubspaceID uint64,
	subspacesData []SubspaceData,
	subspaces []Subspace,
	sections []Section,
	userPermissions []UserPermission,
	userGroups []UserGroup,
	userGroupMembers []UserGroupMemberEntry,
	userGrants []UserGrant,
	groupGrants []GroupGrant,
) *GenesisState {
	return &GenesisState{
		InitialSubspaceID: initialSubspaceID,
		SubspacesData:     subspacesData,
		Subspaces:         subspaces,
		Sections:          sections,
		UserPermissions:   userPermissions,
		UserGroups:        userGroups,
		UserGroupsMembers: userGroupMembers,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		1,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	// Make sure the initial subspace id is valid
	if data.InitialSubspaceID == 0 {
		return fmt.Errorf("invalid initial subspace id: %d", data.InitialSubspaceID)
	}

	// Validate the subspaces data
	for _, entry := range data.SubspacesData {
		if containsDuplicatedSubspaceData(data.SubspacesData, entry) {
			return fmt.Errorf("duplicated subspace data for id: %d", entry.SubspaceID)
		}

		err := entry.Validate()
		if err != nil {
			return err
		}
	}

	// Validate the subspace
	for _, subspace := range data.Subspaces {
		if containsDuplicatedSubspace(data.Subspaces, subspace) {
			return fmt.Errorf("duplicated subspace: %d", subspace.ID)
		}

		err := subspace.Validate()
		if err != nil {
			return err
		}
	}

	for _, section := range data.Sections {
		if containsDuplicatedSection(data.Sections, section) {
			return fmt.Errorf("duplicated section: subspace id %d, section id %d", section.SubspaceID, section.ID)
		}

		err := section.Validate()
		if err != nil {
			return err
		}
	}

	// Validate the user permissions
	for _, entry := range data.UserPermissions {
		if containsDuplicatedUserPermission(data.UserPermissions, entry) {
			return fmt.Errorf("duplicated user permission: subspace id %d, user %s", entry.SubspaceID, entry.User)
		}

		err := entry.Validate()
		if err != nil {
			return err
		}
	}

	// Validate the user groups
	for _, group := range data.UserGroups {
		if containsDuplicatedGroups(data.UserGroups, group) {
			return fmt.Errorf("duplicated group for subspace %d and group %d", group.SubspaceID, group.ID)
		}

		err := group.Validate()
		if err != nil {
			return err
		}
	}

	// Validate the group members
	for _, entry := range data.UserGroupsMembers {
		if containsDuplicatedMembersEntries(data.UserGroupsMembers, entry) {
			return fmt.Errorf("duplicated user group members entry for group %d within subspace %d", entry.GroupID, entry.SubspaceID)
		}

		err := entry.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// containsDuplicatedSubspaceData tells whether the given entries slice contains two or more
// data for the subspace with the same id of the given data
func containsDuplicatedSubspaceData(entries []SubspaceData, data SubspaceData) bool {
	var count = 0
	for _, s := range entries {
		if s.SubspaceID == data.SubspaceID {
			count++
		}
	}
	return count > 1
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

// containsDuplicatedSection tells whether the given sections slice contains two or more
// sections with the same id for the same subspace
func containsDuplicatedSection(sections []Section, section Section) bool {
	var count = 0
	for _, s := range sections {
		if s.SubspaceID == section.SubspaceID && s.ID == section.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedUserPermission tells whether the given entries slice contains two or more
// entries for the same user and subspace section
func containsDuplicatedUserPermission(entries []UserPermission, entry UserPermission) bool {
	var count = 0
	for _, e := range entries {
		if e.SubspaceID == entry.SubspaceID && e.SectionID == entry.SectionID && e.User == entry.User {
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
func containsDuplicatedMembersEntries(entries []UserGroupMemberEntry, entry UserGroupMemberEntry) bool {
	var count = 0
	for _, e := range entries {
		if e.SubspaceID == entry.SubspaceID && e.GroupID == entry.GroupID && e.User == entry.User {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// NewSubspaceData returns a new SubspaceData instance
func NewSubspaceData(subspaceID uint64, nextSectionID uint32, nextGroupID uint32) SubspaceData {
	return SubspaceData{
		SubspaceID:    subspaceID,
		NextGroupID:   nextGroupID,
		NextSectionID: nextSectionID,
	}
}

// Validate implements fmt.Validator
func (data SubspaceData) Validate() error {
	if data.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", data.NextSectionID)
	}

	if data.NextSectionID == 0 {
		return fmt.Errorf("invalid initial section id: %d", data.NextSectionID)
	}

	if data.NextGroupID == 0 {
		return fmt.Errorf("invalid initial group id: %d", data.NextSectionID)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// NewUserGroupMemberEntry returns a new UserGroupMemberEntry instance
func NewUserGroupMemberEntry(subspaceID uint64, groupID uint32, user string) UserGroupMemberEntry {
	return UserGroupMemberEntry{
		SubspaceID: subspaceID,
		GroupID:    groupID,
		User:       user,
	}
}

// Validate implements fmt.Validator
func (entry UserGroupMemberEntry) Validate() error {
	if entry.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", entry.SubspaceID)
	}

	if entry.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", entry.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(entry.User)
	if err != nil {
		return fmt.Errorf("invalid user address: %s", err)
	}

	return nil
}
