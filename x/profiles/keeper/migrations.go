package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/grpc"

	v2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      Keeper
	queryServer grpc.Server
	amino       *codec.LegacyAmino
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, amino *codec.LegacyAmino, queryServer grpc.Server) Migrator {
	return Migrator{
		keeper:      keeper,
		amino:       amino,
		queryServer: queryServer,
	}
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
