package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v4/app"

	"github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v4/x/relationships/testutil"
	"github.com/desmos-labs/desmos/v4/x/relationships/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc      codec.Codec
	ctx      sdk.Context
	storeKey sdk.StoreKey
	k        keeper.Keeper
	sk       *testutil.MockSubspacesKeeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, _ = app.MakeCodecs()

	// Mocks initializations
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.sk = testutil.NewMockSubspacesKeeper(ctrl)

	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.sk)
}
