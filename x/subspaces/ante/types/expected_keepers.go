package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

// DONTCOVER

// AccountKeeper represents the expected keeper used to interact with x/auth
type AccountKeeper interface {
	GetParams(ctx context.Context) authtypes.Params
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	SetAccount(ctx context.Context, acc sdk.AccountI)
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// BankKeeper represents the expected keeper used to interact with x/bank
type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	// Required by auth AnteHandler
	IsSendEnabledCoins(ctx context.Context, coins ...sdk.Coin) error
	SendCoins(ctx context.Context, from, to sdk.AccAddress, amt sdk.Coins) error
}

// SubspacesKeeper represents the expected keeper used to interact with x/subspaces
type SubspacesKeeper interface {
	UseGrantedFees(ctx sdk.Context, subspaceID uint64, grantee sdk.AccAddress, fees sdk.Coins, msgs []sdk.Msg) bool
	GetSubspace(ctx sdk.Context, subspaceID uint64) (types.Subspace, bool)
}

// AuthDeductFeeDecorator represents the expected keeper used to interact with auth.DeductFeeDecorator
type AuthDeductFeeDecorator interface {
	AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error)
}
