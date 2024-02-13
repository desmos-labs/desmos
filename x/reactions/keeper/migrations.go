package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v7/x/reactions/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v7/x/reactions/legacy/v3"
	"github.com/desmos-labs/desmos/v7/x/reactions/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k  Keeper
	sk types.SubspacesKeeper
	pk types.PostsKeeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, sk types.SubspacesKeeper, pk types.PostsKeeper) Migrator {
	return Migrator{
		k:  keeper,
		sk: sk,
		pk: pk,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.k.storeKey, m.sk, m.pk, m.k.cdc)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.k.storeKey, m.pk, m.k.cdc)
}
