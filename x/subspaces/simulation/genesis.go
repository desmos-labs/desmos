package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/feegrant"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for subspaces
func RandomizeGenState(simState *module.SimulationState) {
	subspaces := randomSubspaces(simState.Rand, simState.Accounts)
	sections := randomSections(simState.Rand, subspaces)
	groups := randomUserGroups(simState.Rand, subspaces)
	members := randomUserGroupsMembers(simState.Rand, simState.Accounts, groups)
	acl := randomACL(simState.Rand, simState.Accounts, subspaces)
	initialSubspaceID, subspacesData := getSubspacesDataEntries(subspaces, sections, groups)
	grants := append(randomUserGrants(simState.Rand, simState.Accounts, subspaces), randomGroupGrants(simState.Rand, simState.Accounts, groups)...)

	// Create the genesis and sanitize it
	subspacesGenesis := types.NewGenesisState(initialSubspaceID, subspacesData, subspaces, sections, acl, groups, members, grants)
	subspacesGenesis = sanitizeGenesis(subspacesGenesis)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(subspacesGenesis)
}

// randomSubspaces generates a random slice of subspaces
func randomSubspaces(r *rand.Rand, accs []simtypes.Account) (subspaces []types.Subspace) {
	subspacesNumber := r.Intn(100)
	subspaces = make([]types.Subspace, subspacesNumber)
	for index := 0; index < subspacesNumber; index++ {
		randData := GenerateRandomSubspace(r, accs)
		subspaces[index] = types.NewSubspace(
			uint64(index+1),
			randData.Name,
			randData.Description,
			randData.Treasury,
			randData.Owner,
			randData.Creator,
			time.Now(),
		)
	}
	return subspaces
}

// randomSections generates a random slice of sections
func randomSections(r *rand.Rand, subspaces []types.Subspace) (sections []types.Section) {
	if len(subspaces) == 0 {
		return nil
	}

	sectionsNumber := r.Intn(20)
	sections = make([]types.Section, sectionsNumber)
	for i := 0; i < sectionsNumber; i++ {
		subspace := RandomSubspace(r, subspaces)

		// Generate a random section
		sections[i] = types.NewSection(
			subspace.ID,
			uint32(i)+1,
			0,
			RandomSectionName(r),
			RandomSectionDescription(r),
		)
	}
	return sections
}

// randomUserGroups generates a random slice of user group details
func randomUserGroups(r *rand.Rand, subspaces []types.Subspace) []types.UserGroup {
	if len(subspaces) == 0 {
		return nil
	}

	groupsNumber := r.Intn(30)

	groups := make([]types.UserGroup, groupsNumber)
	for i := 0; i < groupsNumber; i++ {
		subspace := RandomSubspace(r, subspaces)
		groupID := uint32(i + 1)

		// Get a random permission
		permission := RandomPermission(r, validPermissions)

		// Build the group details
		groups[i] = types.NewUserGroup(subspace.ID, 0, groupID, RandomName(r), RandomDescription(r), permission)
	}

	return groups
}

// randomUserGroupsMembers generates a random slice of user group members
func randomUserGroupsMembers(r *rand.Rand, accounts []simtypes.Account, groups []types.UserGroup) []types.UserGroupMemberEntry {
	if len(groups) == 0 {
		return nil
	}

	var membersEntries []types.UserGroupMemberEntry
	for _, group := range groups {
		for i := 0; i < r.Intn(10); i++ {
			account, _ := simtypes.RandomAcc(r, accounts)
			membersEntries = append(membersEntries, types.NewUserGroupMemberEntry(group.SubspaceID, group.ID, account.Address.String()))
		}
	}
	return membersEntries
}

// getSubspacesDataEntries returns the subspace data entries based on the given data
func getSubspacesDataEntries(
	subspaces []types.Subspace, sections []types.Section, groups []types.UserGroup,
) (initialSubspaceID uint64, subspacesData []types.SubspaceData) {
	maxSubspaceID := uint64(0)
	initialSectionID := map[uint64]uint32{}
	initialGroupIDS := map[uint64]uint32{}
	for _, subspace := range subspaces {
		if subspace.ID > initialSubspaceID {
			maxSubspaceID = subspace.ID
		}

		// Get the max section id
		maxSectionID := uint32(0)
		for _, section := range sections {
			if section.SubspaceID == subspace.ID && section.ID > maxSectionID {
				maxSectionID = section.ID
			}
		}
		initialSectionID[subspace.ID] = maxSectionID + 1

		// Get the max group id
		maxGroupID := uint32(0)
		for _, group := range groups {
			if group.SubspaceID == subspace.ID && group.ID > maxGroupID {
				maxGroupID = group.ID
			}
		}
		initialGroupIDS[subspace.ID] = maxGroupID + 1
	}

	subspacesData = make([]types.SubspaceData, len(subspaces))
	for i, subspace := range subspaces {
		subspacesData[i] = types.NewSubspaceData(subspace.ID, initialSectionID[subspace.ID], initialGroupIDS[subspace.ID])
	}

	return maxSubspaceID + 1, subspacesData
}

// randomACL generates a random slice of ACL entries
func randomACL(r *rand.Rand, accounts []simtypes.Account, subspaces []types.Subspace) (entries []types.UserPermission) {
	if len(subspaces) == 0 {
		return nil
	}

	aclEntriesNumber := r.Intn(40)
	entries = make([]types.UserPermission, aclEntriesNumber)
	for index := 0; index < aclEntriesNumber; index++ {
		subspace := RandomSubspace(r, subspaces)
		account, _ := simtypes.RandomAcc(r, accounts)

		// Get a random permission
		permission := RandomPermission(r, validPermissions)

		// Crete the entry
		entries[index] = types.NewUserPermission(subspace.ID, 0, account.Address.String(), permission)
	}

	return entries
}

// randomUserGrants returns randomly generated user grants
func randomUserGrants(r *rand.Rand, accounts []simtypes.Account, subspaces []types.Subspace) []types.Grant {
	if len(subspaces) == 0 {
		return nil
	}
	grantsNumber := r.Intn(30)
	grants := make([]types.Grant, grantsNumber)
	for i := 0; i < grantsNumber; {
		subspace := RandomSubspace(r, subspaces)
		granter, _ := simtypes.RandomAcc(r, accounts)
		grantee, _ := simtypes.RandomAcc(r, accounts)
		if granter.Address.String() == grantee.Address.String() {
			continue
		}
		grant, _ := types.NewGrant(subspace.ID, granter.Address.String(), types.NewUserGrantee(grantee.Address.String()), &feegrant.BasicAllowance{})
		if !containsGrant(grants, grant) {
			grants[i] = grant
			i++
		}
	}
	return grants
}

// randomGroupGrants returns randomly generated group grants
func randomGroupGrants(r *rand.Rand, accounts []simtypes.Account, groups []types.UserGroup) []types.Grant {
	if len(groups) == 0 {
		return nil
	}
	grantsNumber := r.Intn(30)
	grants := make([]types.Grant, grantsNumber)
	for i := 0; i < grantsNumber; {
		group := RandomGroup(r, groups)
		granter, _ := simtypes.RandomAcc(r, accounts)
		grant, _ := types.NewGrant(group.SubspaceID, granter.Address.String(), types.NewGroupGrantee(group.ID), &feegrant.BasicAllowance{})
		if !containsGrant(grants, grant) {
			grants[i] = grant
			i++
		}
	}
	return grants
}

func containsGrant(slice []types.Grant, grant types.Grant) bool {
	for _, g := range slice {
		if g.SubspaceID == grant.SubspaceID && g.Granter == grant.Granter && g.Grantee.Equal(grant.Grantee) {
			return true
		}
	}
	return false
}

// --------------------------------------------------------------------------------------------------------------------

// sanitizeGenesis sanitizes the given genesis by removing all the double subspaces,
// groups or ACL entries that might be there
func sanitizeGenesis(genesis *types.GenesisState) *types.GenesisState {
	return types.NewGenesisState(
		genesis.InitialSubspaceID,
		genesis.SubspacesData,
		genesis.Subspaces,
		genesis.Sections,
		sanitizeUserPermissions(genesis.UserPermissions),
		sanitizeUserGroups(genesis.UserGroups),
		genesis.UserGroupsMembers,
		genesis.Grants,
	)
}

// sanitizeUserGroups sanitizes the given slice by removing all the double groups
func sanitizeUserGroups(slice []types.UserGroup) []types.UserGroup {
	groups := map[string]bool{}
	for _, value := range slice {
		groups[fmt.Sprintf("%d%s", value.SubspaceID, value.Name)] = true
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
func sanitizeUserPermissions(slice []types.UserPermission) []types.UserPermission {
	entries := map[string]bool{}
	for _, value := range slice {
		entries[fmt.Sprintf("%d%s", value.SubspaceID, value.User)] = true
	}

	var unique []types.UserPermission
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
