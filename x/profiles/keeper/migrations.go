package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v043 "github.com/cosmos/cosmos-sdk/x/auth/legacy/v043"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/gogo/protobuf/grpc"

	v210 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v210"

	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
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

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v200.MigrateStore(ctx, m.keeper.StoreKey, m.keeper.paramSubspace, m.keeper.Cdc, m.amino)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v210.MigrateStore(ctx, m.keeper.StoreKey, m.keeper.Cdc)
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	var iterErr error

	m.keeper.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {

		// If the account is a profile, migrate the profile
		if profile, ok := account.(*types.Profile); ok {
			err := m.migrateProfile(ctx, profile)
			if err != nil {
				iterErr = err
				return true
			}
			return false
		}

		// If the account is not a profile migrate it normally
		wb, err := v043.MigrateAccount(ctx, account, m.queryServer)
		if err != nil {
			iterErr = err
			return true
		}

		if wb == nil {
			return false
		}

		m.keeper.ak.SetAccount(ctx, wb)
		return false
	})

	return iterErr
}

// Migrate4to5 migrates from version 4 to 5
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	var iterErr error

	params := m.keeper.GetParams(ctx)
	blockTime := ctx.BlockTime()
	expirationTimeParam := params.AppLinks.ExpirationTime

	m.keeper.IterateApplicationLinks(ctx, func(index int64, link types.ApplicationLink) (stop bool) {
		expirationTime := types.CalculateExpirationTime(link.CreationTime, expirationTimeParam)
		// if the existent app link is expired, delete it
		if expirationTime.Before(blockTime) {
			m.keeper.deleteApplicationLinkStoreKeys(ctx, link)
			return false
		}

		link.ExpirationTime = expirationTime
		// TODO can this error be unchecked? it checks if the link is associated to a profile
		err := m.keeper.SaveApplicationLink(ctx, link)
		if err != nil {
			iterErr = err
			return true
		}
		return false
	})

	return iterErr
}

func (m Migrator) migrateProfile(ctx sdk.Context, profile *types.Profile) error {
	// Do not migrate those profiles that are not based on a VestingAccount
	vestingAcc, ok := profile.GetAccount().(exported.VestingAccount)
	if !ok {
		return nil
	}

	// Migrate the underlying vesting account
	wb, err := v043.MigrateAccount(ctx, vestingAcc, m.queryServer)
	if err != nil {
		return err
	}

	if wb == nil {
		return nil
	}

	// Serialize the underlying vesting account back into the Profile
	accAny, err := codectypes.NewAnyWithValue(wb)
	if err != nil {
		return err
	}
	profile.Account = accAny

	// Store the new Profile as a VestingProfile instead of a common Profile
	// This will grant that future operations that deal with VestingAccount instances are carried out properly
	m.keeper.ak.SetAccount(ctx, profile)
	return nil
}
