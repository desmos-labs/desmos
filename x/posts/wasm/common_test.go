package wasm_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/posts/keeper"
	profileskeeper "github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/suite"

	relationshipskeeper "github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
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
	k        keeper.Keeper

	ak profileskeeper.Keeper
	sk subspaceskeeper.Keeper
	rk relationshipskeeper.Keeper
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, authtypes.StoreKey,
		profilestypes.StoreKey, relationshipstypes.StoreKey,
		subspacestypes.StoreKey, types.StoreKey,
	)
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

	paramsKeeper := paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])

	authKeeper := authkeeper.NewAccountKeeper(suite.cdc, keys[authtypes.StoreKey], authtypes.ProtoBaseAccount, app.GetMaccPerms(), "cosmos", authtypes.NewModuleAddress("gov").String())
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey], nil, nil)
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	suite.ak = profileskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[profilestypes.StoreKey], paramsKeeper.Subspace(profilestypes.DefaultParamsSpace), authKeeper, suite.rk, nil, nil, nil)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.sk,
		suite.rk,
	)
}
