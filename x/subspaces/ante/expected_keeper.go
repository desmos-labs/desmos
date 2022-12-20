package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// DONTCOVER

// AccountKeeper represents the expected keeper used to interact with x/auth
type AccountKeeper interface {
	GetParams(ctx sdk.Context) authtypes.Params
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// BankKeeper represents the expected keeper used to interact with x/bank
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// FeegrantKeeper represents the expected keeper used to interact with x/feegrant
type FeegrantKeeper interface {
	UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error
}

// SubspacesKeeper represents the expected keeper used to interact with x/subspaces
type SubspacesKeeper interface {
	UseGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) bool
}
