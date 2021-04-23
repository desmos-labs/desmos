package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// RelationshipsKeeper represents the expected k to deal with users relationships
type RelationshipsKeeper interface {
	HasUserBlocked(ctx sdk.Context, blocker string, blocked string, subspace string) bool
}
