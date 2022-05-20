package keeper

import (
	"fmt"

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
		k.GetGenesisSubspaces(ctx, k.GetAllSubspaces(ctx)),
		k.GetAllPermissions(ctx),
		k.GetAllUserGroups(ctx),
		k.GetUserAllGroupsMembers(ctx),
	)
}

// GetGenesisSubspaces maps the given subspaces to the corresponding GenesisSubspace instance
func (k Keeper) GetGenesisSubspaces(ctx sdk.Context, subspaces []types.Subspace) []types.GenesisSubspace {
	if subspaces == nil {
		return nil
	}

	genesisSubspaces := make([]types.GenesisSubspace, len(subspaces))
	for i, subspace := range subspaces {
		groupID, err := k.GetGroupID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		genesisSubspaces[i] = types.NewGenesisSubspace(subspace, groupID)
	}
	return genesisSubspaces
}

// GetAllPermissions returns all the stored permissions for all subspaces
func (k Keeper) GetAllPermissions(ctx sdk.Context) []types.UserPermission {
	var entries []types.UserPermission
	k.IterateSubspaces(ctx, func(index int64, subspace types.Subspace) (stop bool) {
		k.IterateSubspacePermissions(ctx, subspace.ID, func(index int64, entry types.UserPermission) (stop bool) {
			entries = append(entries, entry)
			return false
		})
		return false
	})
	return entries
}

// GetAllUserGroups returns the information (name and members) for all the groups of all the subspaces
func (k Keeper) GetAllUserGroups(ctx sdk.Context) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateUserGroups(ctx, func(index int64, group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// GetUserAllGroupsMembers returns all the UserGroupMembersEntry
func (k Keeper) GetUserAllGroupsMembers(ctx sdk.Context) []types.UserGroupMembersEntry {
	var entries []types.UserGroupMembersEntry
	k.IterateUserGroups(ctx, func(index int64, group types.UserGroup) (stop bool) {
		// Skip group ID 0 to avoid exporting any member
		if group.ID == 0 {
			return false
		}

		// Get the group members
		members := k.GetGroupMembers(ctx, group.SubspaceID, group.ID)
		membersAddr := make([]string, len(members))
		for i, member := range members {
			membersAddr[i] = member.String()
		}

		entries = append(entries, types.NewUserGroupMembersEntry(group.SubspaceID, group.ID, membersAddr))
		return false
	})
	return entries
}

// --------------------------------------------------------------------------------------------------------------------

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialize the subspaces
	for _, subspaceData := range data.Subspaces {
		k.SaveSubspace(ctx, subspaceData.Subspace)
		k.SetGroupID(ctx, subspaceData.Subspace.ID, subspaceData.InitialGroupID)
	}

	// Initialize the subspace id setting it to be the max id found + 1
	k.SetSubspaceID(ctx, data.InitialSubspaceID)

	// Initialize the groups with default permission PermissionNothing
	// The real permission will be set later when initializing the various permissions
	for _, group := range data.UserGroups {
		k.SaveUserGroup(ctx, group)
	}

	// Initialize the group members
	for _, entry := range data.UserGroupsMembers {
		// Skip group ID 0 since it's the default group and no user should be here
		if entry.GroupID == 0 {
			continue
		}

		// Initialize the members
		for _, member := range entry.Members {
			userAddr, err := sdk.AccAddressFromBech32(member)
			if err != nil {
				panic(err)
			}

			err = k.AddUserToGroup(ctx, entry.SubspaceID, entry.GroupID, userAddr)
			if err != nil {
				panic(err)
			}
		}
	}

	// Initialize the permissions
	for _, entry := range data.UserPermissions {
		userAddr, err := sdk.AccAddressFromBech32(entry.User)
		if err != nil {
			panic(fmt.Errorf("invalid user address: %s", entry.User))
		}
		k.SetUserPermissions(ctx, entry.SubspaceID, userAddr, entry.Permissions)
	}
}
