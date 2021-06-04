package keeper_test

// NOLINT

import (
	"testing"
	"time"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	profileskeeper "github.com/desmos-labs/desmos/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/desmos-labs/desmos/app"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	postskeeper "github.com/desmos-labs/desmos/x/staging/posts/keeper"
	"github.com/desmos-labs/desmos/x/staging/reports/keeper"
	"github.com/desmos-labs/desmos/x/staging/reports/types"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Marshaler
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              keeper.Keeper
	storeKey       sdk.StoreKey
	postsKeeper    postskeeper.Keeper

	stakingKeeper stakingkeeper.Keeper
	IBCKeeper     *ibckeeper.Keeper
	testData      TestData
}

type TestData struct {
	postID       string
	creationDate time.Time
	creator      string
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, paramstypes.StoreKey, profilestypes.StoreKey)
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
		ms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, nil)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Define keepers
	pk := paramskeeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)

	ak := authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		pk.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	capabilityKeeper := capabilitykeeper.NewKeeper(suite.cdc, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	ScopedProfilesKeeper := capabilityKeeper.ScopeToModule(types.ModuleName)

	scopedIBCKeeper := capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	IBCKeeper := ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibchost.StoreKey],
		pk.Subspace(ibchost.ModuleName),
		suite.stakingKeeper,
		scopedIBCKeeper,
	)

	rk := profileskeeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		pk.Subspace(profilestypes.DefaultParamsSpace),
		ak,
		IBCKeeper.ChannelKeeper,
		&IBCKeeper.PortKeeper,
		ScopedProfilesKeeper,
	)

	suite.postsKeeper = postskeeper.NewKeeper(
		suite.cdc,
		keys[types.StoreKey],
		pk.Subspace(types.DefaultParamsSpace),
		rk,
	)
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.postsKeeper)

	// Setup data
	date, _ := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.testData = TestData{
		postID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		creationDate: date,
		creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
