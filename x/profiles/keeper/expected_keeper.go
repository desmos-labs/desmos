package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

type SubspacesKeeper interface {
	// CheckSubspaceUserPermission checks the permission of the given user inside the subspace with the
	// given id to make sure they are able to perform operations inside it
	CheckSubspaceUserPermission(ctx sdk.Context, subspaceID string, user string) error
}
