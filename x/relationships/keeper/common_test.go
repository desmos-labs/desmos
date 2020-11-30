package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type TestData struct {
	user      string
	otherUser string
}

type KeeperTestSuite struct {
	suite.Suite

	cdc          codec.Marshaler
	legacyAmino  *codec.LegacyAmino
	ctx          sdk.Context
	keeper       keeper.Keeper
	storeKey     sdk.StoreKey
	paramsKeeper paramskeeper.Keeper
	testData     TestData
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store keys
	relationshipsKey := sdk.NewKVStoreKey("storedBlocks")
	suite.storeKey = relationshipsKey

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(relationshipsKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAmino = app.MakeCodecs()
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.storeKey)

	suite.testData.user = "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	suite.testData.otherUser = "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
}
