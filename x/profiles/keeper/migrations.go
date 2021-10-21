package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v210 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v210"

	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
	amino  *codec.LegacyAmino
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, amino *codec.LegacyAmino) Migrator {
	return Migrator{
		keeper: keeper,
		amino:  amino,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v200.MigrateStore(ctx, m.keeper.storeKey, m.keeper.paramSubspace, m.keeper.cdc, m.amino)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v210.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
