package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v4/x/posts/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v4/x/posts/legacy/v3"
	v4 "github.com/desmos-labs/desmos/v4/x/posts/legacy/v4"
	v5 "github.com/desmos-labs/desmos/v4/x/posts/legacy/v5"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k  Keeper
	sk types.SubspacesKeeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, sk types.SubspacesKeeper) Migrator {
	return Migrator{
		k:  keeper,
		sk: sk,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.k.storeKey, m.k.paramsSubspace, m.sk)
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
