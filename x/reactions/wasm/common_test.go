package wasm_test

import (
	"encoding/json"
	"testing"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v6/x/reactions/types"
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
	storeKey       storetypes.StoreKey

	k keeper.Keeper
}

func (suite *Testsuite) SetupTest() {
	// Define the store keys
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

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.k = keeper.NewKeeper(suite.cdc, keys[types.StoreKey], nil, nil, nil, nil)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(Testsuite))
}
