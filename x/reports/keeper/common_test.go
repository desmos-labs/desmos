package keeper_test

// NOLINT

import (
	"testing"
	"time"

	relationshipskeeper "github.com/desmos-labs/desmos/x/relationships/keeper"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/desmos-labs/desmos/app"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
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
	postID       string
	creationDate time.Time
	creator      string
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store keys
	postsKey := sdk.NewKVStoreKey(poststypes.StoreKey)
	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")
	reportsKey := sdk.NewKVStoreKey(types.StoreKey)
	relationshipsKey := sdk.NewKVStoreKey("relationships")
	suite.storeKey = reportsKey

	// create an in-memory db for stored
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(postsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(reportsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	ms.MountStoreWithDB(relationshipsKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// define keepers
	rk := relationshipskeeper.NewKeeper(suite.cdc, relationshipsKey)
	paramsKeeper := paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, paramsKey, paramsTKey)
	suite.postsKeeper = postskeeper.NewKeeper(suite.cdc, postsKey, paramsKeeper.Subspace("poststypes"), rk)
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.postsKeeper)

	// setup data
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
