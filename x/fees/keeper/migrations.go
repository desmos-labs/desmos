package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v4/x/fees/legacy/v2"
	"github.com/desmos-labs/desmos/v4/x/fees/types"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k              Keeper
	legacySubspace types.ParamsSubspace
}

// NewMigrator returns a new Migrator
func NewMigrator(k Keeper, legacySubspace types.ParamsSubspace) Migrator {
	return Migrator{
		k:              k,
		legacySubspace: legacySubspace,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.k.storeKey, m.legacySubspace, m.k.cdc)
}
