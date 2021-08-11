package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllSubspaces(ctx),
		k.GetAllAdmins(ctx),
		k.GetAllRegisteredUsers(ctx),
		k.GetAllBannedUsers(ctx),
		k.GetAllTokenomics(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	// Initialize the subspaces
	for _, subspace := range data.Subspaces {
		err := k.SaveSubspace(ctx, subspace, subspace.Owner)
		if err != nil {
			panic(err)
		}
	}

	// Initialize the admins
	for _, entry := range data.Admins {
		subspace, found := k.GetSubspace(ctx, entry.SubspaceID)
		if !found {
			panic(fmt.Errorf("invalid admins entry: subspace with id %s does not exist", entry.SubspaceID))
		}

		for _, admin := range entry.Users {
			err := k.AddAdminToSubspace(ctx, subspace.ID, admin, subspace.Owner)
			if err != nil {
				panic(err)
			}
		}
	}

	// Initialize the registered users
	for _, entry := range data.RegisteredUsers {
		subspace, found := k.GetSubspace(ctx, entry.SubspaceID)
		if !found {
			panic(fmt.Errorf("invalid registered user entry: subspace with id %s does not exist", entry.SubspaceID))
		}

		for _, user := range entry.Users {
			err := k.RegisterUserInSubspace(ctx, subspace.ID, user, subspace.Owner)
			if err != nil {
				panic(err)
			}
		}
	}

	// Initialize the banned users
	for _, entry := range data.BannedUsers {
		subspace, found := k.GetSubspace(ctx, entry.SubspaceID)
		if !found {
			panic(fmt.Errorf("invalid banned user entry: subspace with id %s does not exist", entry.SubspaceID))
		}

		for _, user := range entry.Users {
			err := k.BanUserInSubspace(ctx, subspace.ID, user, subspace.Owner)
			if err != nil {
				panic(err)
			}
		}
	}

	// Initialize all the tokenomics pairs
	for _, tokenomicsPair := range data.AllTokenomics {
		if err := k.SaveSubspaceTokenomics(ctx, tokenomicsPair); err != nil {
			panic(err)
		}
	}
}
