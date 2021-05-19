package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v0160 "github.com/desmos-labs/desmos/x/profiles/legacy/v0160"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrators
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	params := v0160.MigrateParams(m.keeper.paramSubspace, ctx)
	m.keeper.SetParams(ctx, params)

	return nil
}
