package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
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
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	// create a cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	magpieKeeper := keeper.NewKeeper(cdc, magpieKey)

	return ctx, magpieKeeper
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	auth.RegisterCodec(cdc)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var testSession = types.Session{
	SessionID: types.SessionID(1),
	Owner:     testOwner,
	Created:   10,
	Expiry:    15,
}
