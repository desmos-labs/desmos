package v5

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DONTCOVER

type AccountKeeper interface {
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
	SetAccount(ctx context.Context, acc sdk.AccountI)
}
