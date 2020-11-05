package keeper_test

import (
	"testing"
	"time"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/desmos-labs/desmos/app"
	relationshipskeeper "github.com/desmos-labs/desmos/x/relationships/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.BinaryMarshaler
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	keeper         keeper.Keeper
	storeKey       sdk.StoreKey
	pk             paramskeeper.Keeper
	rk             relationshipskeeper.Keeper
	testData       TestData
}

type TestData struct {
	postID                 string
	postOwner              string
	postCreationDate       time.Time
	postEndPollDate        time.Time
	postEndPollDateExpired time.Time
	answers                types.PollAnswers
	registeredReaction     types.RegisteredReaction
	post                   types.Post
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store keys
	postKey := sdk.NewKVStoreKey(types.StoreKey)
	suite.storeKey = postKey

	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")
	relationshipsKey := sdk.NewKVStoreKey("relationships")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(postKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	ms.MountStoreWithDB(relationshipsKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	blockTime, _ := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.ctx = sdk.NewContext(
		ms,
		tmproto.Header{ChainID: "test-chain-id", Time: blockTime},
		false,
		log.NewNopLogger(),
	)
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()
	suite.pk = paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, paramsKey, paramsTKey)
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, relationshipsKey)
	suite.keeper = keeper.NewKeeper(suite.cdc, postKey, suite.pk.Subspace(types.DefaultParamSpace), suite.rk)

	// Setup data
	suite.testData.postID = "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"
	suite.testData.postOwner = "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"

	suite.testData.postCreationDate = blockTime
	suite.testData.postEndPollDate, _ = time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.testData.postEndPollDateExpired, _ = time.Parse(time.RFC3339, "2019-01-01T01:15:00.000Z")
	suite.testData.answers = types.PollAnswers{
		types.NewPollAnswer("1", "Yes"),
		types.NewPollAnswer("2", "No"),
	}
	suite.testData.post = types.Post{
		PostID:       suite.testData.postID,
		Message:      "Post message",
		Created:      suite.testData.postCreationDate,
		LastEdited:   suite.testData.postCreationDate.Add(1),
		Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData: nil,
		Creator:      suite.testData.postOwner,
		Attachments: types.NewAttachments(
			types.NewAttachment("https://uri.com", "text/plain", []string{suite.testData.postOwner}),
		),
		PollData: &types.PollData{
			Question:              "poll?",
			ProvidedAnswers:       types.NewPollAnswers(suite.testData.answers[0], suite.testData.answers[1]),
			EndDate:               suite.testData.postEndPollDate,
			AllowsMultipleAnswers: true,
			AllowsAnswerEdits:     true,
		},
	}

	suite.testData.registeredReaction = types.NewRegisteredReaction(
		suite.testData.postOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
}
