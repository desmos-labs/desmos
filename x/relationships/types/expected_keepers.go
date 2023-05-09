package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type SubspacesKeeper interface {
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool
	GetAllSubspaces(ctx sdk.Context) []subspacestypes.Subspace
}
