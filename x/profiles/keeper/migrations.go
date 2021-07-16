package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// DONTCOVER

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
