package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"

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
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              keeper.Keeper
	storeKey       storetypes.StoreKey

	ak          authkeeper.AccountKeeper
	authzKeeper authzkeeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, authtypes.StoreKey, authzkeeper.StoreKey)
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

	// Dependencies initialization
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
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.ak, suite.authzKeeper, authtypes.NewModuleAddress("gov").String())
}

func (suite *KeeperTestSuite) getAllGrantsInExpiringQueue(ctx sdk.Context) []types.Grant {
	store := ctx.KVStore(suite.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ExpiringAllowanceQueuePrefix)
	defer iterator.Close()

	var grants []types.Grant
	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		suite.cdc.MustUnmarshal(store.Get(types.ParseAllowanceKeyFromExpiringKey(iterator.Key())), &grant)
		grants = append(grants, grant)
	}

	return grants
}
