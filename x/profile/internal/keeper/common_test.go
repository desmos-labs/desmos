package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
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

	// create a cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	// define keepers
	paramsKeeper := params.NewKeeper(cdc, sdk.NewKVStoreKey("params"), sdk.NewTransientStoreKey("transient_params"))

	return ctx, keeper.NewKeeper(cdc, magpieKey, paramsKeeper.Subspace(types.DefaultParamspace).WithKeyTable(types.ParamKeyTable()))
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
var profilePic = "https://shorturl.at/adnX3"
var profileCov = "https://shorturl.at/cgpyF"
var testPictures = types.NewPictures(&profilePic, &profileCov)
var name = "name"
var surname = "surname"
var bio = "biography"

var testProfile = types.Profile{
	Name:     &name,
	Surname:  &surname,
	Moniker:  "moniker",
	Bio:      &bio,
	Pictures: testPictures,
	Creator:  testPostOwner,
}
