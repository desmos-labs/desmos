package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v6/x/posts/keeper"
	"github.com/desmos-labs/desmos/v6/x/posts/testutil"
	"github.com/desmos-labs/desmos/v6/x/posts/types"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	db "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey storetypes.StoreKey
	k        *keeper.Keeper

	ctrl *gomock.Controller
	ak   *testutil.MockProfilesKeeper
	sk   *testutil.MockSubspacesKeeper
	rk   *testutil.MockRelationshipsKeeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey)
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

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Mocks initializations
	suite.ctrl = gomock.NewController(suite.T())

	suite.sk = testutil.NewMockSubspacesKeeper(suite.ctrl)
	suite.rk = testutil.NewMockRelationshipsKeeper(suite.ctrl)
	suite.ak = testutil.NewMockProfilesKeeper(suite.ctrl)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.ak,
		suite.sk,
		suite.rk,
		authtypes.NewModuleAddress("gov").String(),
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}
