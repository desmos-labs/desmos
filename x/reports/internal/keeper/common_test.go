package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/keeper"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models/common"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
	"time"
)

func SetupTestInput() (sdk.Context, keeper.Keeper) {

	// define store keys
	postsKey := sdk.NewKVStoreKey(posts.StoreKey)
	reportsKey := sdk.NewKVStoreKey(common.StoreKey)

	//post keeper
	postsK := posts.NewKeeper(testCodecPosts(), postsKey)

	// create an in-memory db for reports
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(postsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(reportsKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	// create a cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	return ctx, keeper.NewKeeper(postsK, cdc, reportsKey)
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

func testCodecPosts() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	posts.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

var (
	creator, _           = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	postID               = posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	timeZone, _          = time.LoadLocation("UTC")
	testPostCreationDate = time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
)
