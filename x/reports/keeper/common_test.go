package keeper_test

// NOLINT

import (
	"testing"
	"time"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/desmos-labs/desmos/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Marshaler
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	keeper         keeper.Keeper
	storeKey       sdk.StoreKey
	postsKeeper    postskeeper.Keeper
	testData       TestData
}

type TestData struct {
	postID       poststypes.PostID
	creationDate time.Time
	creator      sdk.AccAddress
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store keys
	postsKey := sdk.NewKVStoreKey(poststypes.StoreKey)
	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")
	reportsKey := sdk.NewKVStoreKey(types.StoreKey)
	suite.storeKey = reportsKey

	// create an in-memory db for stored
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(postsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(reportsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// define keepers
	paramsKeeper := paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, paramsKey, paramsTKey)
	suite.postsKeeper = postskeeper.NewKeeper(suite.legacyAminoCdc, postsKey, paramsKeeper.Subspace("poststypes"))
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.postsKeeper)

	// setup data
	addr, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	date, _ := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.testData = TestData{
		postID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		creationDate: date,
		creator:      addr,
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
