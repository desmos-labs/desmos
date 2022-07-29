package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	v2 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v3"
	v4 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v4"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      Keeper
	authzKeeper authzkeeper.Keeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, authzKeeper authzkeeper.Keeper) Migrator {
	return Migrator{
		keeper:      keeper,
		authzKeeper: authzKeeper,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.authzKeeper, m.keeper.cdc)
}
