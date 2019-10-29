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
var testTimeZone, _ = time.LoadLocation("UTC")
var testPost = types.Post{
	PostID:        types.PostID(3257),
	ParentID:      types.PostID(502),
	Message:       "Post message",
	Created:       time.Date(2019, 10, 31, 12, 25, 0, 0, testTimeZone),
	Modified:      time.Date(2019, 11, 1, 12, 25, 0, 0, testTimeZone),
	Likes:         54,
	Owner:         testPostOwner,
	Namespace:     "cosmos",
	ExternalOwner: "cosmos1qe2vysfe8gsqcg0mr0qejd9urknnk7aa9r9fk2",
}
