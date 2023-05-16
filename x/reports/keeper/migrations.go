package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v5/x/reports/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v5/x/reports/legacy/v3"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k  Keeper
	sk types.SubspacesKeeper

	legacySubspace types.ParamsSubspace
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, sk types.SubspacesKeeper, legacySubspace types.ParamsSubspace) Migrator {
	return Migrator{
		k:              keeper,
		sk:             sk,
		legacySubspace: legacySubspace,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.k.storeKey, m.legacySubspace, m.sk)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.k.storeKey, m.legacySubspace, m.k.cdc)
}
