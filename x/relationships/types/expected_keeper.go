package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SubspacesKeeper interface {
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool
}
