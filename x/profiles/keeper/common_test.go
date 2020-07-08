package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
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
	postOwner sdk.AccAddress
	profile   types.Profile
}

func (suite *KeeperTestSuite) SetupTest() {
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

	suite.ctx = sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc = testCodec()
	suite.paramsKeeper = params.NewKeeper(suite.cdc, paramsKey, paramsTKey)
	suite.keeper = keeper.NewKeeper(suite.cdc, profileKey, suite.paramsKeeper.Subspace(types.DefaultParamspace))

	// setup Data
	// nolint - errcheck
	suite.testData.postOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.testData.profile = types.Profile{
		DTag: "dtag",
		Bio:  newStrPtr("biography"),
		Pictures: types.NewPictures(
			newStrPtr("https://shorturl.at/adnX3"),
			newStrPtr("https://shorturl.at/cgpyF"),
		),
		Creator: suite.testData.postOwner,
	}
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

// newStrPtr allows to easily create a new string pointer starting
// from a string value, for easier test setup
func newStrPtr(value string) *string {
	return &value
}
