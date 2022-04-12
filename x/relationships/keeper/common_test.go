package keeper_test

import (
	"testing"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	profileskeeper "github.com/desmos-labs/desmos/v3/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v3/app"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"

	"github.com/desmos-labs/desmos/v3/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       sdk.StoreKey
	k              keeper.Keeper
	pk             profileskeeper.Keeper
	sk             subspaceskeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, capabilitytypes.StoreKey,

		types.StoreKey, profilestypes.StoreKey, subspacestypes.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

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
	for _, memKey := range memKeys {
		ms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, nil)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	paramsKeeper := paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)
	authKeeper := authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	suite.pk = profileskeeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		keys[profilestypes.StoreKey],
		paramsKeeper.Subspace(profilestypes.DefaultParamsSpace),
		authKeeper,
		suite.k,
		nil,
		nil,
		nil,
	)
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey])
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.sk)
}
