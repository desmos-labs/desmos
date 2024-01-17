package types

import (
	context "context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
)

// DONTCOVER

type AccountKeeper interface {
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
	SetAccount(ctx context.Context, acc sdk.AccountI)
}

type AuthzKeeper interface {
	SaveGrant(ctx context.Context, grantee, granter sdk.AccAddress, authorization authztypes.Authorization, expiration *time.Time) error
	DeleteGrant(ctx context.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) error
}
