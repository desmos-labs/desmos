package keeper_test

import (
	"testing"

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
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/links/keeper"
	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc              codec.Marshaler
	legacyAminoCdc   *codec.LegacyAmino
	ctx              sdk.Context
	storeKey         sdk.StoreKey
	k                keeper.Keeper
	paramsKeeper     paramskeeper.Keeper
	stakingKeeper    stakingkeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	scopedIBCKeeper  capabilitykeeper.ScopedKeeper
	testData         TestData
}

type TestData struct {
	user      string
	otherUser string
	link      types.Link
}

func (suite *KeeperTestSuite) SetupTest() {
	// TO DO: ibc setting

	// define store keys
	linkKey := sdk.NewKVStoreKey("links")
	suite.storeKey = linkKey
	accountKey := sdk.NewKVStoreKey("acc")
	ibchostKey := sdk.NewKVStoreKey("ibc")
	capabilityKey := sdk.NewKVStoreKey("capability")
	memKeys := sdk.NewMemoryStoreKeys("mem_capability")
	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(linkKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(accountKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(ibchostKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	ms.MountStoreWithDB(capabilityKey, sdk.StoreTypeIAVL, nil)

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

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName: nil,
	}

	accountKeeper := authkeeper.NewAccountKeeper(
		suite.cdc,
		accountKey,
		suite.paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.IBCKeeper.ChannelKeeper,
		&suite.IBCKeeper.PortKeeper,
		scopedIBCKeeper,
		accountKeeper,
	)

	// setup Data
	suite.testData.user = "desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk"
	suite.testData.otherUser = "desmos1488h84vd9rc0dmwxx9gzskmymwr7afcemegt9q"
	suite.testData.link = types.NewLink("desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk", "cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn")
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
