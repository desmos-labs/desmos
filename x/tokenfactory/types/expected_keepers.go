package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// BankKeeper represents a keeper that deals with x/bank
type BankKeeper interface {
	// Methods imported from bank should be defined here
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)

	HasSupply(ctx sdk.Context, denom string) bool

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
}

// AccountKeeper represents a keeper that deals with x/auth
type AccountKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
}

// SubspacesKeeper represents a keeper that deals with x/subspaces
type SubspacesKeeper interface {
	GetAllSubspaces(ctx sdk.Context) []subspacestypes.Subspace
	GetSubspace(ctx sdk.Context, subspaceID uint64) (subspace subspacestypes.Subspace, found bool)
	HasPermission(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permission subspacestypes.Permission) bool
	GetUsersWithRootPermissions(ctx sdk.Context, subspaceID uint64, permission subspacestypes.Permissions) []string
}
