package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllSubspaces(ctx),
		k.GetSubspaceAdminsEntry(ctx),
		k.GetBlockedToPostUsers(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {

	for _, subspace := range data.Subspaces {
		if err := subspace.Validate(); err != nil {
			panic(err)
		}

		if err := k.SaveSubspace(ctx, subspace); err != nil {
			panic(err)
		}
	}

	for _, adminsEntry := range data.SubspaceAdmins {
		for _, admin := range adminsEntry.Admins.Users {
			if err := k.AddAdminToSubspace(ctx, adminsEntry.SubspaceId, admin); err != nil {
				panic(err)
			}
		}

	}

	for _, blockedUsersEntry := range data.BlockedToPostUsers {
		for _, userToBlock := range blockedUsersEntry.Users.Users {
			if err := k.BlockPostsForUser(ctx, userToBlock, blockedUsersEntry.SubspaceId); err != nil {
				panic(err)
			}
		}
	}

}
