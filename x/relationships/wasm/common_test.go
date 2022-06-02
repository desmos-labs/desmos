package wasm_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/desmos-labs/desmos/v3/app"
	profileskeeper "github.com/desmos-labs/desmos/v3/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	"github.com/desmos-labs/desmos/v3/x/relationships/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

func buildCreateRelationshipRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{CreateRelationship: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildDeleteRelationshipRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{DeleteRelationship: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildBlockUserRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{BlockUser: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildUnblockUserRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{UnblockUser: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildRelationshipsQueryRequest(cdc codec.Codec, query *types.QueryRelationshipsRequest) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsQuery{Relationships: cdc.MustMarshalJSON(query)})
	return bz
}

func buildBlocksQueryRequest(cdc codec.Codec, query *types.QueryBlocksRequest) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsQuery{Blocks: cdc.MustMarshalJSON(query)})
	return bz
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       sdk.StoreKey
	k              keeper.Keeper
	pk             profileskeeper.Keeper
	sk             subspaceskeeper.Keeper
}

func (suite *TestSuite) SetupTest() {
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
