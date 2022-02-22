package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	profilesv2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"

	v1 "github.com/desmos-labs/desmos/v2/x/relationships/legacy/v1"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
	pk     profilesv2.Keeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, pk profilesv2.Keeper) Migrator {
	return Migrator{
		keeper: keeper,
		pk:     pk,
	}
}

// Migrate0To1 migrates from version 0 to 1.
func (m Migrator) Migrate0To1(ctx sdk.Context) error {
	return v1.MigrateStore(ctx, m.pk, m.keeper.storeKey, m.keeper.cdc)
}
