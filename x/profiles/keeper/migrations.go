package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/gogo/protobuf/grpc"

	v4 "github.com/desmos-labs/desmos/v3/x/profiles/legacy/v4"
	v5 "github.com/desmos-labs/desmos/v3/x/profiles/legacy/v5"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      Keeper
	ak          authkeeper.AccountKeeper
	queryServer grpc.Server
}

// NewMigrator returns a new Migrator
func NewMigrator(ak authkeeper.AccountKeeper, keeper Keeper, queryServer grpc.Server) Migrator {
	return Migrator{
		keeper:      keeper,
		ak:          ak,
		queryServer: queryServer,
	}
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.ak, m.keeper.storeKey, m.keeper.legacyAmino, m.keeper.cdc)
}

// Migrate5To6 migrates from version 5 to 6.
func (m Migrator) Migrate5To6(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc, m.keeper.legacyAmino)
}
