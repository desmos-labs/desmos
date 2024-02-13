package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	db "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v7/app"
	"github.com/desmos-labs/desmos/v7/x/tokenfactory/keeper"
	"github.com/desmos-labs/desmos/v7/x/tokenfactory/testutil"
	"github.com/desmos-labs/desmos/v7/x/tokenfactory/types"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey storetypes.StoreKey
	k        keeper.Keeper

	ctrl *gomock.Controller
	bk   *testutil.MockBankKeeper
	sk   *testutil.MockSubspacesKeeper
	ak   *testutil.MockAccountKeeper
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

	suite.bk = testutil.NewMockBankKeeper(suite.ctrl)
	suite.sk = testutil.NewMockSubspacesKeeper(suite.ctrl)
	suite.ak = testutil.NewMockAccountKeeper(suite.ctrl)

	suite.k = keeper.NewKeeper(
		suite.storeKey,
		suite.cdc,
		suite.sk,
		suite.ak,
		suite.bk,
		authtypes.NewModuleAddress("gov").String(),
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}
