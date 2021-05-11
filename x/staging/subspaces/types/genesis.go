package types

import "fmt"

// NewGenesisState creates a new genesis state
func NewGenesisState(subspaces []Subspace, adminsEntries []SubspaceAdminsEntry, blockedUsers []BlockedUsersEntry) *GenesisState {
	return &GenesisState{
		Subspaces:          subspaces,
		SubspaceAdmins:     adminsEntries,
		BlockedToPostUsers: blockedUsers,
	}
}

func NewAdminsEntries(subspaceId string, admins Users) SubspaceAdminsEntry {
	return SubspaceAdminsEntry{
		SubspaceId: subspaceId,
		Admins:     admins,
	}
}

func NewBlockedUsersEntry(subspaceId string, users Users) BlockedUsersEntry {
	return BlockedUsersEntry{
		SubspaceId: subspaceId,
		Users:      users,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, subspace := range data.Subspaces {
		if err := subspace.Validate(); err != nil {
			return err
		}
	}

	for _, adminsEntry := range data.SubspaceAdmins {
		if !doesSubspaceExists(data.Subspaces, adminsEntry.SubspaceId) {
			return fmt.Errorf("the subspace with id: %s does not exist", adminsEntry.SubspaceId)
		}
	}

	for _, blockedUsersEntry := range data.BlockedToPostUsers {
		if !doesSubspaceExists(data.Subspaces, blockedUsersEntry.SubspaceId) {
			return fmt.Errorf("the subspace with id: %s does not exist", blockedUsersEntry.SubspaceId)
		}
	}

	return nil
}

// doesSubspaceExists check if the subspaces array contains a subspace with the given id
func doesSubspaceExists(subspaces []Subspace, id string) bool {
	for _, sub := range subspaces {
		if sub.Id == id {
			return true
		}
	}
	return false
}
