package keeper_test

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

func SetupTestInput() (sdk.Context, keeper.Keeper) {

	// define store keys
	magpieKey := sdk.NewKVStoreKey("magpie")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(magpieKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	// create a Cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	return ctx, keeper.NewKeeper(cdc, magpieKey)
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
var timeZone, _ = time.LoadLocation("UTC")
var testPostCreationDate = time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
var testPostEndPollDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)
var testPostEndPollDateExpired = time.Date(2019, 1, 1, 1, 15, 00, 000, timeZone)
var answer = types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}

var answer2 = types.PollAnswer{ID: types.AnswerID(2), Text: "No"}
var postID = types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")

var testPost = types.NewPost(
	postID,
	"",
	"Post message",
	false,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	map[string]string{},
	testPostCreationDate,
	testPostOwner,
).WithMedias(types.NewPostMedias(
	types.NewPostMedia("https://uri.com", "text/plain"),
)).WithPollData(types.NewPollData(
	"poll?",
	testPostEndPollDate,
	types.NewPollAnswers(answer, answer2),
	true,
	true,
	true,
))

var testRegisteredReaction = types.NewReaction(testPostOwner, ":smile:", "https://smile.jpg",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
