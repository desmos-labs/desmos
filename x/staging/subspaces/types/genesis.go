package types

import "fmt"

// NewUsersEntry allows to build a new UsersEntry instance
func NewUsersEntry(subspaceID string, users []string) UsersEntry {
	return UsersEntry{
		SubspaceID: subspaceID,
		Users:      users,
	}
}

// NewGenesisState creates a new genesis state
func NewGenesisState(subspaces []Subspace, admins, registeredUsers, bannedUsers []UsersEntry) *GenesisState {
	return &GenesisState{
		Subspaces:       subspaces,
		Admins:          admins,
		RegisteredUsers: registeredUsers,
		BannedUsers:     bannedUsers,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, nil)
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
			return fmt.Errorf("duplicated subspace: %s", subspace.ID)
		}
	}

	for _, entry := range data.Admins {
		if isEntryDuplicated(data.Admins, entry) {
			return fmt.Errorf("duplicated admins entry for subspace with id %s", entry.SubspaceID)
		}

		for _, admin := range entry.Users {
			if containsDuplicatedValue(entry.Users, admin) {
				return fmt.Errorf("duplicated admin for subspace with id %s: %s", entry.SubspaceID, admin)
			}
		}
	}

	for _, entry := range data.RegisteredUsers {
		if isEntryDuplicated(data.RegisteredUsers, entry) {
			return fmt.Errorf("duplicated registered users entry for subspace with id %s", entry.SubspaceID)
		}

		for _, user := range entry.Users {
			if containsDuplicatedValue(entry.Users, user) {
				return fmt.Errorf("duplicated registered user for subspace with id %s: %s", entry.SubspaceID, user)
			}
		}
	}

	for _, entry := range data.BannedUsers {
		if isEntryDuplicated(data.BannedUsers, entry) {
			return fmt.Errorf("duplicated banned users entry for subspace with id %s", entry.SubspaceID)
		}

		for _, user := range entry.Users {
			if containsDuplicatedValue(entry.Users, user) {
				return fmt.Errorf("duplicated banned user for subspace with id %s: %s", entry.SubspaceID, user)
			}
		}
	}

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

func isEntryDuplicated(entries []UsersEntry, entry UsersEntry) bool {
	var count = 0
	for _, e := range entries {
		if e.SubspaceID == entry.SubspaceID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedValue tells whether the given entries slice contains duplicates of the provided value
func containsDuplicatedValue(slice []string, value string) bool {
	var count = 0
	for _, e := range slice {
		if e == value {
			count++
		}
	}
	return count > 1
}
