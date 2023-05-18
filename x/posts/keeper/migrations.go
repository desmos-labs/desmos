package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v3"
	v4 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v4"
	v5 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v5"
	v6 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v6"
	v7 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v7"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
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
	return v3.MigrateStore(ctx, m.k.storeKey, m.k.cdc)
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.k.storeKey, m.k.cdc)
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.k.storeKey, m.k.cdc)
}

// Migrate5to6 migrates from version 5 to 6.
func (m Migrator) Migrate5to6(ctx sdk.Context) error {
	return v6.MigrateStore(ctx, m.k.storeKey, m.legacySubspace, m.k.cdc)
}

// Migrate6to7 migrates from version 6 to 7.
func (m Migrator) Migrate6to7(ctx sdk.Context) error {
	return v7.MigrateStore(ctx, m.k.storeKey, m.k.cdc)
}
