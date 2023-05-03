package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v5/app"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              keeper.Keeper
	paramsKeeper   paramskeeper.Keeper
	storeKey       storetypes.StoreKey

	ak          authkeeper.AccountKeeper
	authzKeeper authzkeeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, paramstypes.StoreKey, authtypes.StoreKey, authzkeeper.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Dependencies initialization
	suite.paramsKeeper = paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])
	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)
	suite.authzKeeper = authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], suite.cdc, &baseapp.MsgServiceRouter{}, suite.ak)

	// Define keeper
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.ak, suite.authzKeeper)
}
