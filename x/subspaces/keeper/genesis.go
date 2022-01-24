package keeper

import (
	"github.com/desmos-labs/desmos/v2/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllSubspaces(ctx),
		k.GetAllUserGroups(ctx),
		k.GetAllPermissions(ctx),
	)
}

// GetAllPermissions returns all the stored permissions for all subspaces
func (k Keeper) GetAllPermissions(ctx sdk.Context) []types.ACLEntry {
	var entries []types.ACLEntry
	k.IterateSubspaces(ctx, func(index int64, subspace types.Subspace) (stop bool) {
		k.IteratePermissions(ctx, subspace.ID, func(index int64, target string, permission types.Permission) (stop bool) {
			entries = append(entries, types.NewACLEntry(subspace.ID, target, permission))
			return false
		})
		return false
	})

	return entries
}

// GetAllUserGroups returns the information (name and members) for all the groups of all the subspaces
func (k Keeper) GetAllUserGroups(ctx sdk.Context) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateSubspaces(ctx, func(index int64, subspace types.Subspace) (stop bool) {
		k.IterateSubspaceGroups(ctx, subspace.ID, func(index int64, groupName string) (stop bool) {

			var members []string
			k.IterateGroupMembers(ctx, subspace.ID, groupName, func(index int64, member sdk.AccAddress) (stop bool) {
				members = append(members, member.String())
				return false
			})

			groups = append(groups, types.NewUserGroup(subspace.ID, groupName, members))
			return false
		})

		return false
	})
	return groups
}

// --------------------------------------------------------------------------------------------------------------------

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialize the subspaces
	for _, subspace := range data.Subspaces {
		k.SaveSubspace(ctx, subspace)
	}

	// Initialize the groups with default permission PermissionNothing
	// The real permission will be set later when initializing the various permissions
	for _, group := range data.UserGroups {
		k.SaveUserGroup(ctx, group.SubspaceID, group.Name, types.PermissionNothing)

		// Initialize the members
		for _, member := range group.Members {
			userAddr, err := sdk.AccAddressFromBech32(member)
			if err != nil {
				panic(err)
			}

			err = k.AddUserToGroup(ctx, group.SubspaceID, group.Name, userAddr)
			if err != nil {
				panic(err)
			}
		}
	}

	// Initialize the permissions
	for _, entry := range data.ACL {
		k.SetPermissions(ctx, entry.SubspaceID, entry.Target, entry.Permissions)
	}
}
