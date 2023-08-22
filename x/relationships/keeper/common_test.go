package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"

	"github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v6/x/relationships/testutil"
	"github.com/desmos-labs/desmos/v6/x/relationships/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc      codec.Codec
	ctx      sdk.Context
	storeKey storetypes.StoreKey
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
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
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
