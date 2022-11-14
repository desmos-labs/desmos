package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) auth.ModuleAccountI

	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) auth.AccountI
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) auth.AccountI
	SetAccount(ctx sdk.Context, acc auth.AccountI)
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
