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
		k.getSubspacesData(ctx),
		k.GetAllSubspaces(ctx),
		k.GetAllSections(ctx),
		k.getAllUserPermissions(ctx),
		k.GetAllUserGroups(ctx),
		k.getUserAllGroupsMembers(ctx),
	)
}

// getSubspacesData maps the given subspaces to the corresponding GenesisSubspace instance
func (k Keeper) getSubspacesData(ctx sdk.Context) []types.SubspaceData {
	var data []types.SubspaceData
	k.IterateSubspaces(ctx, func(index int64, subspace types.Subspace) (stop bool) {
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
	k.IterateUserPermissions(ctx, func(index int64, subspaceID uint64, sectionID uint32, user sdk.AccAddress, permission types.Permission) (stop bool) {
		entries = append(entries, types.NewUserPermission(subspaceID, sectionID, user.String(), permission))
		return false
	})
	return entries
}

// getUserAllGroupsMembers returns all the UserGroupMembersEntry
func (k Keeper) getUserAllGroupsMembers(ctx sdk.Context) []types.UserGroupMembersEntry {
	var entries []types.UserGroupMembersEntry
	k.IterateUserGroups(ctx, func(index int64, group types.UserGroup) (stop bool) {
		// Skip group ID 0 to avoid exporting any member
		if group.ID == 0 {
			return false
		}

		// Get the group members
		members := k.GetUserGroupMembers(ctx, group.SubspaceID, group.ID)
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
		k.SetUserPermissions(ctx, entry.SubspaceID, entry.SectionID, userAddr, entry.Permissions)
	}
}
