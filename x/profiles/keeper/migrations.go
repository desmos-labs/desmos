package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v0160 "github.com/desmos-labs/desmos/x/profiles/legacy/v0160"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
	amino  *codec.LegacyAmino
}

// NewMigrator returns a new Migrators
func NewMigrator(amino *codec.LegacyAmino, keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
		amino:  amino,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	params, err := v0160.MigrateParams(ctx, m.amino, m.keeper.paramSubspace)
	if err != nil {
		return err
	}

	m.keeper.SetParams(ctx, params)
	return nil
}
