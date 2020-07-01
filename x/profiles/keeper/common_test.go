package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

func SetupTestInput() (sdk.Context, keeper.Keeper) {

	// define store keys
	profileKey := sdk.NewKVStoreKey("profiles")
	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(profileKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	// create a cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	// define keepers
	paramsKeeper := params.NewKeeper(cdc, paramsKey, paramsTKey)
	subspace := paramsKeeper.Subspace(types.DefaultParamspace)

	return ctx, keeper.NewKeeper(cdc, profileKey, subspace)
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

// newStrPtr allows to easily create a new string pointer starting
// from a string value, for easier test setup
func newStrPtr(value string) *string {
	return &value
}

var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
var testProfile = types.Profile{
	DTag: "dtag",
	Bio:  newStrPtr("biography"),
	Pictures: types.NewPictures(
		newStrPtr("https://shorturl.at/adnX3"),
		newStrPtr("https://shorturl.at/cgpyF"),
	),
	Creator: testPostOwner,
}
