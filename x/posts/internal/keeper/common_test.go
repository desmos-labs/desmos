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

	// define store store keys
	magpieKey := sdk.NewKVStoreKey("magpie")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(magpieKey, sdk.StoreTypeIAVL, memDB)
	_ = ms.LoadLatestVersion()

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
var answer = types.PollAnswer{ID: uint(1), Text: "Yes"}

var answer2 = types.PollAnswer{ID: uint(2), Text: "No"}

var testPost = types.NewPost(
	types.PostID(3257),
	types.PostID(0),
	"Post message",
	false,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	map[string]string{},
	testPostCreationDate,
	testPostOwner,
	types.PostMedias{types.NewPostMedia(
		"https://uri.com",
		"text/plain"),
	},
	types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
)
