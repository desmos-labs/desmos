package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for subspaces
func RandomizeGenState(simState *module.SimulationState) {
	subspaces := randomSubspaces(simState.Rand, simState.Accounts)
	groups := randomUserGroups(simState.Rand, simState.Accounts, subspaces)
	acl := randomACL(simState.Rand, simState.Accounts, subspaces, groups)

	// Create the genesis and sanitize it
	subspacesGenesis := types.NewGenesisState(subspaces, groups, acl)
	subspacesGenesis = sanitizeGenesis(subspacesGenesis)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(subspacesGenesis)
}

// randomSubspaces returns randomly generated genesis account
func randomSubspaces(r *rand.Rand, accs []simtypes.Account) (subspaces []types.Subspace) {
	subspacesNumber := r.Intn(100)
	subspaces = make([]types.Subspace, subspacesNumber)
	for index := 0; index < subspacesNumber; index++ {
		subspaces[index] = GenerateRandomSubspace(r, accs)
	}
	return subspaces
}

// randomUserGroups generates random slice of user group details
func randomUserGroups(r *rand.Rand, accounts []simtypes.Account, subspaces []types.Subspace) (groups []types.UserGroup) {
	groupsNumber := r.Intn(30)
	groups = make([]types.UserGroup, groupsNumber)
	for i := 0; i < groupsNumber; i++ {
		subspace, _ := RandomSubspace(r, subspaces)

		// Get a random number of members
		membersNumber := r.Intn(5)
		members := make([]string, membersNumber)
		for j := 0; j < membersNumber; j++ {
			members[j] = RandomAccount(r, accounts).Address.String()
		}
		members = sanitizeStrings(members)

		// Build the group details
		groups[i] = types.NewUserGroup(subspace.ID, RandomName(r), members)
	}

	return groups
}

// randomACL generates a random slice of ACL entries
func randomACL(r *rand.Rand, accounts []simtypes.Account, subspaces []types.Subspace, groups []types.UserGroup) (entries []types.ACLEntry) {
	aclEntriesNumber := r.Intn(40)
	entries = make([]types.ACLEntry, aclEntriesNumber)
	for index := 0; index < aclEntriesNumber; index++ {
		subspace, _ := RandomSubspace(r, subspaces)
		target := RandomAccount(r, accounts).Address.String()

		// 50% of chance of selecting a group rather than an account
		if r.Intn(101) <= 50 {
			target = RandomGroup(r, groups).Name
		}

		// Get a random permission
		permission := RandomPermission(r, []types.Permission{
			types.PermissionNothing,
			types.PermissionWrite,
			types.PermissionManageGroups,
			types.PermissionEverything,
		})

		// Crete the entry
		entries[index] = types.NewACLEntry(subspace.ID, target, permission)
	}

	return entries
}

// --------------------------------------------------------------------------------------------------------------------

// sanitizeGenesis sanitizes the given genesis by removing all the double subspaces,
// groups or ACL entries that might be there
func sanitizeGenesis(genesis *types.GenesisState) *types.GenesisState {
	return types.NewGenesisState(
		sanitizeSubspaces(genesis.Subspaces),
		sanitizeUserGroups(genesis.UserGroups),
		sanitizeACLEntry(genesis.ACL),
	)
}

// sanitizeSubspaces sanitizes the given slice by removing all the double subspaces
func sanitizeSubspaces(slice []types.Subspace) []types.Subspace {
	ids := map[uint64]int{}
	for _, value := range slice {
		ids[value.ID] = 1
	}

	var unique []types.Subspace
	for id := range ids {
	SubspaceLoop:
		for _, subspace := range slice {
			if id == subspace.ID {
				unique = append(unique, subspace)
				break SubspaceLoop
			}
		}
	}

	return unique
}

// sanitizeUserGroups sanitizes the given slice by removing all the double groups
func sanitizeUserGroups(slice []types.UserGroup) []types.UserGroup {
	groups := map[string]int{}
	for _, value := range slice {
		groups[fmt.Sprintf("%d%s", value.SubspaceID, value.Name)] = 1
	}

	var unique []types.UserGroup
	for id := range groups {
	EntryLoop:
		for _, group := range slice {
			if id == fmt.Sprintf("%d%s", group.SubspaceID, group.Name) {
				unique = append(unique, group)
				break EntryLoop
			}
		}
	}

	return unique
}

// sanitizeSubspaces sanitizes the given slice by removing all the double entries
func sanitizeACLEntry(slice []types.ACLEntry) []types.ACLEntry {
	entries := map[string]int{}
	for _, value := range slice {
		entries[fmt.Sprintf("%d%s", value.SubspaceID, value.Target)] = 1
	}

	var unique []types.ACLEntry
	for id := range entries {
	EntryLoop:
		for _, entry := range slice {
			if id == fmt.Sprintf("%d%s", entry.SubspaceID, entry.Target) {
				unique = append(unique, entry)
				break EntryLoop
			}
		}
	}

	return unique
}

// sanitizeStrings sanitizes the given slice by removing all duplicated values
func sanitizeStrings(slice []string) []string {
	values := map[string]int{}
	for _, value := range slice {
		values[value] = 1
	}

	count := 0
	unique := make([]string, len(values))
	for value := range values {
		unique[count] = value
		count++
	}

	return unique
}
