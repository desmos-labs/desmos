package types

import "fmt"

// NewACLEntry returns a new ACLEntry instance
func NewACLEntry(subspaceID uint64, target string, permissions Permission) ACLEntry {
	return ACLEntry{
		SubspaceId:  subspaceID,
		Target:      target,
		Permissions: permissions,
	}
}

// NewUserGroup returns a new UserGroup instance
func NewUserGroup(subspaceID uint64, groupName string, members []string) UserGroup {
	return UserGroup{
		SubspaceID: subspaceID,
		Name:       groupName,
		Members:    members,
	}
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
		if containsDuplicatedSubspace(data.Subspaces, subspace) {
			return fmt.Errorf("duplicated subspace: %d", subspace.ID)
		}
	}

	// TODO

	return nil
}

// containsDuplicatedSubspace tells whether the given subspaces slice contain duplicates of the provided subspace
func containsDuplicatedSubspace(subspaces []Subspace, subspace Subspace) bool {
	var count = 0
	for _, s := range subspaces {
		if s.Equal(subspace) {
			count++
		}
	}
	return count > 1
}
