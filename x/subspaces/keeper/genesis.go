package keeper

import (
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	subspaceID, err := k.GetSubspaceID(ctx)
	if err != nil {
		panic(err)
	}

	return types.NewGenesisState(
		subspaceID,
		k.getSubspacesData(ctx),
		k.GetAllSubspaces(ctx),
		k.GetAllSections(ctx),
		k.getAllUserPermissions(ctx),
		k.GetAllUserGroups(ctx),
		k.getAllUserGroupsMembers(ctx),
	)
}

// getSubspacesData returns all the stored information for all subspaces
func (k Keeper) getSubspacesData(ctx sdk.Context) []types.SubspaceData {
	var data []types.SubspaceData
	k.IterateSubspaces(ctx, func(subspace types.Subspace) (stop bool) {
		nextSectionID, err := k.GetNextSectionID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		nextGroupID, err := k.GetNextGroupID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		data = append(data, types.NewSubspaceData(subspace.ID, nextSectionID, nextGroupID))

		return false
	})
	return data
}

// getAllUserPermissions returns all the stored user permissions for all subspaces
func (k Keeper) getAllUserPermissions(ctx sdk.Context) []types.UserPermission {
	var entries []types.UserPermission
	k.IterateUserPermissions(ctx, func(entry types.UserPermission) (stop bool) {
		entries = append(entries, entry)
		return false
	})
	return entries
}

// getAllUserGroupsMembers returns all the stored user group members
func (k Keeper) getAllUserGroupsMembers(ctx sdk.Context) []types.UserGroupMemberEntry {
	var entries []types.UserGroupMemberEntry
	k.IterateUserGroupsMembers(ctx, func(entry types.UserGroupMemberEntry) (stop bool) {
		// Skip group ID 0 to avoid exporting any member
		if entry.GroupID == 0 {
			return false
		}

		entries = append(entries, entry)
		return false
	})
	return entries
}

// --------------------------------------------------------------------------------------------------------------------

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialize the subspace id setting it to be the max id found + 1
	k.SetSubspaceID(ctx, data.InitialSubspaceID)

	// Initialize the subspaces data
	for _, entry := range data.SubspacesData {
		k.SetNextGroupID(ctx, entry.SubspaceID, entry.NextGroupID)
		k.SetNextSectionID(ctx, entry.SubspaceID, entry.NextSectionID)
	}

	// Initialize the subspaces
	for _, subspace := range data.Subspaces {
		k.SaveSubspace(ctx, subspace)
	}

	// Initialize the sections
	for _, section := range data.Sections {
		k.SaveSection(ctx, section)
	}

	// Initialize the groups with default permission PermissionNothing
	for _, group := range data.UserGroups {
		k.SaveUserGroup(ctx, group)
	}

	// Initialize the group members
	for _, entry := range data.UserGroupsMembers {
		k.AddUserToGroup(ctx, entry.SubspaceID, entry.GroupID, entry.User)
	}

	// Initialize the permissions
	for _, entry := range data.UserPermissions {
		k.SetUserPermissions(ctx, entry.SubspaceID, entry.SectionID, entry.User, entry.Permissions)
	}
}
