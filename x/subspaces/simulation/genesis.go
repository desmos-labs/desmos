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
	groups, members := randomUserGroups(simState.Rand, simState.Accounts, subspaces)
	acl := randomACL(simState.Rand, simState.Accounts, subspaces)
	initialSubspaceID, genSubspaces := getInitialIDs(subspaces, groups)

	// Create the genesis and sanitize it
	subspacesGenesis := types.NewGenesisState(initialSubspaceID, genSubspaces, acl, groups, members)
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
func randomUserGroups(
	r *rand.Rand, accounts []simtypes.Account, subspaces []types.Subspace,
) (groups []types.UserGroup, membersEntries []types.UserGroupMembersEntry) {
	groupsNumber := r.Intn(30)

	groups = make([]types.UserGroup, groupsNumber)
	membersEntries = make([]types.UserGroupMembersEntry, groupsNumber)

	for i := 0; i < groupsNumber; i++ {
		subspace := RandomSubspace(r, subspaces)
		groupID := uint32(i + 1)

		// Get a random permission
		permission := RandomPermission(r, []types.Permission{
			types.PermissionNothing,
			types.PermissionWrite,
			types.PermissionManageGroups,
			types.PermissionEverything,
		})

		// Build the group details
		groups[i] = types.NewUserGroup(subspace.ID, groupID, RandomName(r), RandomDescription(r), permission)

		// Get a random number of members
		membersNumber := r.Intn(5)
		members := make([]string, membersNumber)
		for j := 0; j < membersNumber; j++ {
			members[j] = RandomAccount(r, accounts).Address.String()
		}
		members = sanitizeStrings(members)

		// Build the members details
		membersEntries[i] = types.NewUserGroupMembersEntry(subspace.ID, groupID, members)
	}

	return groups, membersEntries
}

// getInitialIDs returns the initial subspace id and various initial group ids given the slice of subspaces and groups
func getInitialIDs(
	subspaces []types.Subspace, groups []types.UserGroup,
) (initialSubspaceID uint64, genSubspaces []types.GenesisSubspace) {
	initialGroupIDS := map[uint64]uint32{}
	for _, subspace := range subspaces {
		if subspace.ID > initialSubspaceID {
			initialSubspaceID = subspace.ID
		}

		// Get the max group id
		maxGroupID := uint32(0)
		for _, group := range groups {
			if group.SubspaceID == subspace.ID && group.ID > maxGroupID {
				maxGroupID = group.ID
			}
		}

		// Get the initial group id for this subspace
		initialGroupIDS[subspace.ID] = maxGroupID + 1
	}

	genSubspaces = make([]types.GenesisSubspace, len(subspaces))
	for i, subspace := range subspaces {
		genSubspaces[i] = types.NewGenesisSubspace(subspace, initialGroupIDS[subspace.ID])
	}

	return initialSubspaceID, genSubspaces
}

// randomACL generates a random slice of ACL entries
func randomACL(r *rand.Rand, accounts []simtypes.Account, subspaces []types.Subspace) (entries []types.ACLEntry) {
	aclEntriesNumber := r.Intn(40)
	entries = make([]types.ACLEntry, aclEntriesNumber)
	for index := 0; index < aclEntriesNumber; index++ {
		subspace := RandomSubspace(r, subspaces)
		target := RandomAccount(r, accounts).Address.String()

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
		genesis.InitialSubspaceID,
		sanitizeSubspaces(genesis.Subspaces),
		sanitizeACLEntry(genesis.ACL),
		sanitizeUserGroups(genesis.UserGroups),
		genesis.UserGroupsMembers,
	)
}

// sanitizeSubspaces sanitizes the given slice by removing all the double subspaces
func sanitizeSubspaces(slice []types.GenesisSubspace) []types.GenesisSubspace {
	ids := map[uint64]int{}
	for _, value := range slice {
		ids[value.Subspace.ID] = 1
	}

	var unique []types.GenesisSubspace
	for id := range ids {
	SubspaceLoop:
		for _, subspace := range slice {
			if id == subspace.Subspace.ID {
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
		entries[fmt.Sprintf("%d%s", value.SubspaceID, value.User)] = 1
	}

	var unique []types.ACLEntry
	for id := range entries {
	EntryLoop:
		for _, entry := range slice {
			if id == fmt.Sprintf("%d%s", entry.SubspaceID, entry.User) {
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
