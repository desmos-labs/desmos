package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// DONTCOVER

// AccountKeeper represents the expected keeper used to interact with x/auth
type AccountKeeper interface {
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	SetAccount(ctx sdk.Context, account authtypes.AccountI)
}
