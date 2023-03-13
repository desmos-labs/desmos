package keeper_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v4/x/posts/keeper"
	"github.com/desmos-labs/desmos/v4/x/posts/testutil"
	"github.com/desmos-labs/desmos/v4/x/posts/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v4/app"
)

type KeeperTestsuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey sdk.StoreKey
	k        keeper.Keeper

	ak *testutil.MockProfilesKeeper
	sk *testutil.MockSubspacesKeeper
	rk *testutil.MockRelationshipsKeeper
}

func (suite *KeeperTestsuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, types.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
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

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Dependencies initializations
	paramsKeeper := paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])

	// Mocks initializations
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.sk = testutil.NewMockSubspacesKeeper(ctrl)
	suite.rk = testutil.NewMockRelationshipsKeeper(ctrl)
	suite.ak = testutil.NewMockProfilesKeeper(ctrl)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.sk,
		suite.rk,
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestsuite))
}
