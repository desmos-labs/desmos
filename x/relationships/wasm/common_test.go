package wasm_test

import (
	"encoding/json"
	"testing"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/stretchr/testify/suite"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v4/app"
	profileskeeper "github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"
	"github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/relationships/types"
)

func buildCreateRelationshipRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.RelationshipsMsg{CreateRelationship: &raw})
	return bz
}

func buildDeleteRelationshipRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.RelationshipsMsg{DeleteRelationship: &raw})
	return bz
}

func buildBlockUserRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.RelationshipsMsg{BlockUser: &raw})
	return bz
}

func buildUnblockUserRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.RelationshipsMsg{UnblockUser: &raw})
	return bz
}

func buildRelationshipsQueryRequest(cdc codec.Codec, query *types.QueryRelationshipsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.RelationshipsQuery{Relationships: &raw})
	return bz
}

func buildBlocksQueryRequest(cdc codec.Codec, query *types.QueryBlocksRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.RelationshipsQuery{Blocks: &raw})
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
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey], nil, nil)
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.sk)
}
