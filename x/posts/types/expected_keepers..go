package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type SubspacesKeeper interface {
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool
	HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission subspacestypes.Permission) bool
}
