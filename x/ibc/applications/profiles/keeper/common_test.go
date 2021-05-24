package keeper_test

import (
	"testing"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/keeper"
	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/desmos-labs/desmos/app"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Marshaler
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       sdk.StoreKey
	k              keeper.Keeper
	ak             authkeeper.AccountKeeper
	paramsKeeper   paramskeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey, paramstypes.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, sdk.StoreTypeTransient, memDB)
	}
	for _, memKey := range memKeys {
		ms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.paramsKeeper = paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)

	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		suite.paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		suite.cdc,
		keys[banktypes.StoreKey],
		suite.ak,
		suite.paramsKeeper.Subspace(banktypes.ModuleName),
		map[string]bool{},
	)

	stakingKeeper := stakingkeeper.NewKeeper(
		suite.cdc,
		keys[stakingtypes.StoreKey],
		suite.ak,
		bankKeeper,
		suite.paramsKeeper.Subspace(stakingtypes.ModuleName),
	)

	capabilityKeeper := capabilitykeeper.NewKeeper(
		suite.cdc,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	scopedIBCKeeper := capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	ibcKeeper := ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibchost.StoreKey],
		suite.paramsKeeper.Subspace(ibchost.ModuleName),
		stakingKeeper,
		scopedIBCKeeper,
	)

	scopedIBCProfilesKeeper := capabilityKeeper.ScopeToModule(types.ModuleName)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		scopedIBCProfilesKeeper,
	)
}
