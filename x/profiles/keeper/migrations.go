package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/gogo/protobuf/grpc"

	"github.com/desmos-labs/desmos/v3/x/profiles/legacy/v1beta1"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      Keeper
	ak          authkeeper.AccountKeeper
	queryServer grpc.Server
	amino       *codec.LegacyAmino
}

// NewMigrator returns a new Migrator
func NewMigrator(ak authkeeper.AccountKeeper, keeper Keeper, amino *codec.LegacyAmino, queryServer grpc.Server) Migrator {
	return Migrator{
		keeper:      keeper,
		ak:          ak,
		amino:       amino,
		queryServer: queryServer,
	}
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v1beta1.MigrateStore(ctx, m.ak, m.keeper.storeKey, m.amino, m.keeper.cdc)
}
