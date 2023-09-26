package v5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DONTCOVER

type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) sdk.AccountI
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	SetAccount(ctx sdk.Context, acc sdk.AccountI)
}
