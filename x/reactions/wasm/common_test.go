package wasm_test

import (
	"encoding/json"
	"testing"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	profileskeeper "github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

	postskeeper "github.com/desmos-labs/desmos/v4/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"

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
	"github.com/desmos-labs/desmos/v4/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
	relationshipskeeper "github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func buildAddReactionRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReactionsMsg{AddReaction: &raw})
	return bz
}

func buildRemoveReactionRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReactionsMsg{RemoveReaction: &raw})
	return bz
}

func buildAddRegisteredReactionRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReactionsMsg{AddRegisteredReaction: &raw})
	return bz
}

func buildEditRegisteredReactionRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReactionsMsg{EditRegisteredReaction: &raw})
	return bz
}

func buildRemoveRegisteredReactionRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReactionsMsg{RemoveRegisteredReaction: &raw})
	return bz
}

func buildSetReactionsParamsRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReactionsMsg{SetReactionsParams: &raw})
	return bz
}

func buildReactionsQueryRequest(cdc codec.Codec, query *types.QueryReactionsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReactionsQuery{Reactions: &raw})
	return bz
}

func buildReactionQueryRequest(cdc codec.Codec, query *types.QueryReactionRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReactionsQuery{Reaction: &raw})
	return bz
}

func buildRegisteredReactionsQueryRequest(cdc codec.Codec, query *types.QueryRegisteredReactionsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReactionsQuery{RegisteredReactions: &raw})
	return bz
}

func buildRegisteredReactionQueryRequest(cdc codec.Codec, query *types.QueryRegisteredReactionRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReactionsQuery{RegisteredReaction: &raw})
	return bz
}

func buildReactionsParamsQueryRequest(cdc codec.Codec, query *types.QueryReactionsParamsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReactionsQuery{ReactionsParams: &raw})
	return bz
}

type Testsuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       sdk.StoreKey

	ak profileskeeper.Keeper
	rk relationshipskeeper.Keeper
	pk postskeeper.Keeper
	sk subspaceskeeper.Keeper
	k  keeper.Keeper
}

func (suite *Testsuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, authtypes.StoreKey,
		profilestypes.StoreKey, relationshipstypes.StoreKey,
		subspacestypes.StoreKey, poststypes.StoreKey,
		types.StoreKey,
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

	authKeeper := authkeeper.NewAccountKeeper(suite.cdc, keys[authtypes.StoreKey], paramsKeeper.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, app.GetMaccPerms())
	suite.ak = profileskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[profilestypes.StoreKey], paramsKeeper.Subspace(profilestypes.DefaultParamsSpace), authKeeper, suite.rk, nil, nil, nil)
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey], authKeeper)
	suite.pk = postskeeper.NewKeeper(
		suite.cdc,
		keys[poststypes.StoreKey],
		paramsKeeper.Subspace(poststypes.DefaultParamsSpace),
		suite.pk,
		suite.sk,
		suite.rk,
	)
	suite.k = keeper.NewKeeper(suite.cdc, keys[types.StoreKey], suite.ak, suite.sk, suite.rk, suite.pk)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(Testsuite))
}
