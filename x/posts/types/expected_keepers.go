package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SubspacesKeeper represents a keeper that deals with subspaces
type SubspacesKeeper interface {
	// HasSubspace tells whether the subspace with the given id exists or not
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool

	// HasPermission tells whether the given user has the provided permission inside the subspace with the specified id
	HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission subspacestypes.Permission) bool
}

// RelationshipsKeeper represents a keeper that deals with relationships
type RelationshipsKeeper interface {
	// HasUserBlocked tells whether the given blocker has blocked the user inside the provided subspace
	HasUserBlocked(ctx sdk.Context, blocker, user string, subspaceID uint64) bool

	// HasRelationship tells whether the relationship between the user and counterparty exists for the given subspace
	HasRelationship(ctx sdk.Context, user, counterparty string, subspaceID uint64) bool
}
