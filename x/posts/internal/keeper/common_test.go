package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc          *codec.Codec
	ctx          sdk.Context
	keeper       keeper.Keeper
	paramsKeeper params.Keeper
	testData     TestData
}

type TestData struct {
	postID                 types.PostID
	postOwner              sdk.AccAddress
	timeZone               *time.Location
	postCreationDate       time.Time
	postEndPollDate        time.Time
	postEndPollDateExpired time.Time
	answers                types.PollAnswers
	registeredReaction     types.Reaction
	post                   types.Post
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store keys
	postKey := sdk.NewKVStoreKey(common.StoreKey)
	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(postKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc = testCodec()
	suite.paramsKeeper = params.NewKeeper(suite.cdc, paramsKey, paramsTKey)
	suite.keeper = keeper.NewKeeper(suite.cdc, postKey, suite.paramsKeeper.Subspace(types.DefaultParamspace))

	// setup Data
	suite.testData.postID = "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"
	// nolint - errcheck
	suite.testData.postOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	// nolint - errcheck
	suite.testData.timeZone, _ = time.LoadLocation("UTC")
	suite.testData.postCreationDate = time.Date(2020, 1, 1, 15, 15, 00, 000, suite.testData.timeZone)
	suite.testData.postEndPollDate = time.Date(2050, 1, 1, 15, 15, 00, 000, suite.testData.timeZone)
	suite.testData.postEndPollDateExpired = time.Date(2019, 1, 1, 1, 15, 00, 000, suite.testData.timeZone)
	suite.testData.answers = types.PollAnswers{types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}, types.PollAnswer{ID: types.AnswerID(2), Text: "No"}}
	suite.testData.post = types.NewPost(
		suite.testData.postID,
		"",
		"Post message",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		suite.testData.postCreationDate,
		suite.testData.postOwner,
	).WithMedias(types.NewPostMedias(
		types.NewPostMedia("https://uri.com", "text/plain", []sdk.AccAddress{suite.testData.postOwner}),
	)).WithPollData(types.NewPollData(
		"poll?",
		suite.testData.postEndPollDate,
		types.NewPollAnswers(suite.testData.answers[0], suite.testData.answers[1]),
		true,
		true,
		true,
	))

	suite.testData.registeredReaction = types.NewReaction(suite.testData.postOwner, ":smile:", "https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}
