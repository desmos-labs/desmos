package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
)

// DONTCOVER

type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) sdk.AccountI
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	SetAccount(ctx sdk.Context, acc sdk.AccountI)
}

type AuthzKeeper interface {
	SaveGrant(ctx sdk.Context, grantee, granter sdk.AccAddress, authorization authztypes.Authorization, expiration *time.Time) error
	DeleteGrant(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) error
}
