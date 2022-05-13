package keeper_test

import (
	"testing"

	relationshipskeeper "github.com/desmos-labs/desmos/v3/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v3/x/relationships/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"

	"github.com/desmos-labs/desmos/v3/x/posts/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v3/app"
)

type KeeperTestsuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey sdk.StoreKey
	k        keeper.Keeper

	sk subspaceskeeper.Keeper
	rk relationshipskeeper.Keeper
	pk paramskeeper.Keeper
}

func (suite *KeeperTestsuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, relationshipstypes.StoreKey, subspacestypes.StoreKey, paramstypes.StoreKey)
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

	suite.pk = paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)

	// Define k
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey])
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.pk.Subspace(types.DefaultParamsSpace),
		suite.sk,
		suite.rk,
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestsuite))
}
