package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/desmos-labs/desmos/app"
	ibctesting "github.com/desmos-labs/desmos/testing"
	"github.com/desmos-labs/desmos/x/links/keeper"
	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	// for default test
	cdc              codec.BinaryMarshaler
	legacyAminoCdc   *codec.LegacyAmino
	ctx              sdk.Context
	storeKey         sdk.StoreKey
	k                keeper.Keeper
	paramsKeeper     paramskeeper.Keeper
	stakingKeeper    stakingkeeper.Keeper
	accountKeeper    authkeeper.AccountKeeper
	IBCKeeper        *ibckeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	testData         TestData

	// for ibc test
	coordinator *ibctesting.Coordinator
	chainA      *ibctesting.TestChain
	chainB      *ibctesting.TestChain
	queryClient types.QueryClient
}

type TestData struct {
	user      string
	otherUser string
	link      types.Link
}

func (suite *KeeperTestSuite) SetupTest() {

	// define store keys
	linkKey := sdk.NewKVStoreKey(types.StoreKey)
	suite.storeKey = linkKey
	accountKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	ibchostKey := sdk.NewKVStoreKey(ibchost.StoreKey)
	capabilityKey := sdk.NewKVStoreKey(capabilitytypes.StoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	paramsTKey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(linkKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(accountKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(ibchostKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	ms.MountStoreWithDB(capabilityKey, sdk.StoreTypeIAVL, memDB)

	for _, memKey := range memKeys {
		ms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, nil)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.paramsKeeper = paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, paramsKey, paramsTKey)
	suite.capabilityKeeper = capabilitykeeper.NewKeeper(suite.cdc, capabilityKey, memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := suite.capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	suite.IBCKeeper = ibckeeper.NewKeeper(
		suite.cdc,
		ibchostKey,
		suite.paramsKeeper.Subspace(ibchost.ModuleName),
		suite.stakingKeeper,
		scopedIBCKeeper,
	)

	maccPerms := map[string][]string{}

	suite.accountKeeper = authkeeper.NewAccountKeeper(
		suite.cdc,
		accountKey,
		suite.paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	scopedLinksKeeper := suite.capabilityKeeper.ScopeToModule(types.ModuleName)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.IBCKeeper.ChannelKeeper,
		&suite.IBCKeeper.PortKeeper,
		scopedLinksKeeper,
		suite.accountKeeper,
	)

	// setup Data
	suite.testData.user = "desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk"
	suite.testData.otherUser = "desmos1488h84vd9rc0dmwxx9gzskmymwr7afcemegt9q"
	suite.testData.link = types.NewLink("desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk", "cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn")
}

func (suite *KeeperTestSuite) SetupIBCTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), suite.chainA.App.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.chainA.App.LinksKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	k := new(KeeperTestSuite)
	suite.Run(t, k)
}

func (suite *KeeperTestSuite) RequireErrorsEqual(expected, actual error) {
	if expected != nil {
		suite.Require().Error(actual)
		suite.Require().Equal(expected.Error(), actual.Error())
	} else {
		suite.Require().NoError(actual)
	}
}
