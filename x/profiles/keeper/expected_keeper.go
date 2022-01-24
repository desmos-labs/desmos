package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// SubspacesKeeper represents the expected keeper used to interact with subspaces
type SubspacesKeeper interface {
	// HasSubspace tells if the subspace with the given id exists
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool

	// GetAllSubspaces returns all the subspaces stored
	GetAllSubspaces(ctx sdk.Context) []subspacestypes.Subspace
}
