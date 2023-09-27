package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

// BankKeeper represents a keeper that deals with x/bank
type BankKeeper interface {
	// Methods imported from bank should be defined here
	GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx context.Context, denomMetaData banktypes.Metadata)

	HasSupply(ctx context.Context, denom string) bool

	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	HasBalance(ctx context.Context, addr sdk.AccAddress, amt sdk.Coin) bool
}

// AccountKeeper represents a keeper that deals with x/auth
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
}

// SubspacesKeeper represents a keeper that deals with x/subspaces
type SubspacesKeeper interface {
	GetAllSubspaces(ctx sdk.Context) []subspacestypes.Subspace
	GetSubspace(ctx sdk.Context, subspaceID uint64) (subspace subspacestypes.Subspace, found bool)
	HasPermission(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permission subspacestypes.Permission) bool
	GetUsersWithRootPermissions(ctx sdk.Context, subspaceID uint64, permission subspacestypes.Permissions) []string
}
