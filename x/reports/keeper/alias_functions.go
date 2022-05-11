package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// HasSubspace tells whether the subspace with the given id exists or not
func (k Keeper) HasSubspace(ctx sdk.Context, subspaceID uint64) bool {
	return k.sk.HasSubspace(ctx, subspaceID)
}

// HasPermission tells whether the given user has the provided permission inside the subspace with the specified id
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission subspacestypes.Permission) bool {
	return k.sk.HasPermission(ctx, subspaceID, user, permission)
}

// HasUserBlocked tells whether the given blocker has blocked the user inside the provided subspace
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, user string, subspaceID uint64) bool {
	return k.rk.HasUserBlocked(ctx, blocker, user, subspaceID)
}
