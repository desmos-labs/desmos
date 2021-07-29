package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/subspaces/types"
)

// RelationshipsKeeper represents the expected k to deal with users relationships
type RelationshipsKeeper interface {
	HasUserBlocked(ctx sdk.Context, blocker string, blocked string, subspace string) bool
}

type SubspacesKeeper interface {
	// CheckSubspaceUserPermission checks the permission of the given user inside the subspace with the
	// given id to make sure they are able to perform operations inside it
	CheckSubspaceUserPermission(ctx sdk.Context, subspaceID string, user string) error

	IterateTokenomics(ctx sdk.Context, fn func(index int64, tokenomics types.Tokenomics) (stop bool))
}
