package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	v4 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v4"
	v5 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v5"
	v6 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v6"
	v7 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v7"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
	ak     authkeeper.AccountKeeper
}

// NewMigrator returns a new Migrator
func NewMigrator(ak authkeeper.AccountKeeper, keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
		ak:     ak,
	}
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.ak, m.keeper.storeKey, m.keeper.legacyAmino, m.keeper.cdc)
}

// Migrate5to6 migrates from version 5 to 6.
func (m Migrator) Migrate5to6(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc, m.keeper.legacyAmino)
}

// Migrate6to7 migrates from version 6 to 7.
func (m Migrator) Migrate6to7(ctx sdk.Context) error {
	return v6.MigrateStore(ctx, m.keeper.ak, m.keeper.storeKey, m.keeper.legacyAmino, m.keeper.cdc)
}

// Migrate7to8 migrates from version 7 to 8.
func (m Migrator) Migrate7to8(ctx sdk.Context) error {
	return v7.MigrateStore(ctx, m.keeper.paramSubspace)
}
