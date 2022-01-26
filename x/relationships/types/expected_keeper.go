package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ProfileKeeper interface {
	HasProfile(ctx sdk.Context, user string) bool
}

type SubspacesKeeper interface {
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool
}
