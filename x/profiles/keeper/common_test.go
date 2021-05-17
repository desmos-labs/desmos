package keeper_test

import (
	"testing"
	"time"

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

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
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

	stakingKeeper    stakingkeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	testData         TestData
}

type TestData struct {
	user      string
	otherUser string
	profile   *types.Profile
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, capabilitytypes.StoreKey)
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

	suite.capabilityKeeper = capabilitykeeper.NewKeeper(suite.cdc, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	ScopedProfilesKeeper := suite.capabilityKeeper.ScopeToModule(types.ModuleName)

	scopedIBCKeeper := suite.capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	suite.IBCKeeper = ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibchost.StoreKey],
		suite.paramsKeeper.Subspace(ibchost.ModuleName),
		suite.stakingKeeper,
		scopedIBCKeeper,
	)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.IBCKeeper.ChannelKeeper,
		&suite.IBCKeeper.PortKeeper,
		ScopedProfilesKeeper,
	)

	// Set test data
	suite.testData.user = "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	suite.testData.otherUser = "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"

	addr, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

	// Create the base account and set inside the auth keeper.
	// This is done in order to make sure that when we try to create a profile using the above address, the profile
	// can be created properly. Not storing the base account would end up in the following error since it's null:
	// "the given account cannot be serialized using Protobuf"
	baseAcc := authtypes.NewBaseAccountWithAddress(addr)
	suite.ak.SetAccount(suite.ctx, baseAcc)

	suite.testData.profile, err = types.NewProfile(
		"dtag",
		"test-user",
		"biography",
		types.NewPictures(
			"https://shorturl.at/adnX3",
			"https://shorturl.at/cgpyF",
		),
		time.Time{},
		baseAcc,
	)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) CheckProfileNoError(profile *types.Profile, err error) *types.Profile {
	suite.Require().NoError(err)
	return profile
}
