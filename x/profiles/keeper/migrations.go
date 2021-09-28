package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v200 "github.com/desmos-labs/desmos/x/profiles/legacy/v200"
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
