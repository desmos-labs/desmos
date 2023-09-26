package wasm_test

import (
	"encoding/json"
	"testing"

	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/posts/keeper"

	"cosmossdk.io/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

func buildCreatePostRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.PostsMsg{CreatePost: &raw})
	return bz
}

func buildEditPostRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.PostsMsg{EditPost: &raw})
	return bz
}

func buildDeletePostRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.PostsMsg{DeletePost: &raw})
	return bz
}

func buildAddPostAttachmentRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.PostsMsg{AddPostAttachment: &raw})
	return bz
}

func buildRemovePostAttachmentRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.PostsMsg{RemovePostAttachment: &raw})
	return bz
}

func buildAnswerPollRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.PostsMsg{AnswerPoll: &raw})
	return bz
}

func buildSubspacePostsQueryRequest(cdc codec.Codec, query *types.QuerySubspacePostsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.PostsQuery{SubspacePosts: &raw})
	return bz
}

func buildIncomingDtagTransferQueryRequest(cdc codec.Codec, query *types.QuerySectionPostsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.PostsQuery{SectionPosts: &raw})
	return bz
}

func buildPostQueryRequest(cdc codec.Codec, query *types.QueryPostRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.PostsQuery{Post: &raw})
	return bz
}

func buildAppLinksQueryRequest(cdc codec.Codec, query *types.QueryPostAttachmentsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.PostsQuery{PostAttachments: &raw})
	return bz
}

func buildPollAnswersQueryRequest(cdc codec.Codec, query *types.QueryPollAnswersRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.PostsQuery{PollAnswers: &raw})
	return bz
}

type TestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey storetypes.StoreKey
	k        *keeper.Keeper
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	// Define store keys
	keys := storetypes.NewMemoryStoreKeys(types.StoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		nil,
		nil,
		nil,
		authtypes.NewModuleAddress("gov").String(),
	)
}
