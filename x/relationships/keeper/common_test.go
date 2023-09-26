package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
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

	ctrl *gomock.Controller
	sk   *testutil.MockSubspacesKeeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := storetypes.NewKVStoreKeys(types.StoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, _ = app.MakeCodecs()

	// Mocks initializations
	suite.ctrl = gomock.NewController(suite.T())

	suite.sk = testutil.NewMockSubspacesKeeper(suite.ctrl)

	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.sk)
}

func (suite *KeeperTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}
