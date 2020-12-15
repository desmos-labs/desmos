package keeper_test

import (
	"testing"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	relationshipskeeper "github.com/desmos-labs/desmos/x/relationships/keeper"

	"github.com/desmos-labs/desmos/app"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

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
	rk             relationshipskeeper.Keeper
	paramsKeeper   paramskeeper.Keeper
	testData       TestData
}

type TestData struct {
	user      string
	otherUser string
	profile   types.Profile
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store keys
	profileKey := sdk.NewKVStoreKey("profiles")
	suite.storeKey = profileKey

	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")
	relationshipsKey := sdk.NewKVStoreKey("relationships")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(profileKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	ms.MountStoreWithDB(relationshipsKey, sdk.StoreTypeIAVL, memDB)

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, relationshipsKey)
	suite.paramsKeeper = paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, paramsKey, paramsTKey)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.paramsKeeper.Subspace(types.DefaultParamspace),
		suite.rk,
	)

	// setup Data
	// nolint - errcheck
	suite.testData.user = "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	suite.testData.otherUser = "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
	suite.testData.profile = types.Profile{
		Dtag: "dtag",
		Bio:  "biography",
		Pictures: types.NewPictures(
			"https://shorturl.at/adnX3",
			"https://shorturl.at/cgpyF",
		),
		Creator: suite.testData.user,
	}
}

func (suite *KeeperTestSuite) RequireErrorsEqual(expected, actual error) {
	if expected != nil {
		suite.Require().Error(actual)
		suite.Require().Equal(expected.Error(), actual.Error())
	} else {
		suite.Require().NoError(actual)
	}
}
