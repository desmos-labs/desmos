package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc      *codec.Codec
	ctx      sdk.Context
	keeper   keeper.Keeper
	ms       store.CommitMultiStore
	testData TestData
}

type TestData struct {
	owner   sdk.AccAddress
	session types.Session
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store store keys
	magpieKey := sdk.NewKVStoreKey("magpie")

	// create an in-memory db
	memDB := db.NewMemDB()
	suite.ms = store.NewCommitMultiStore(memDB)
	suite.ms.MountStoreWithDB(magpieKey, sdk.StoreTypeIAVL, memDB)
	if err := suite.ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(suite.ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc = testCodec()
	suite.keeper = keeper.NewKeeper(suite.cdc, magpieKey)

	// setup Data
	// nolint - errcheck
	suite.testData.owner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.testData.session = types.Session{
		SessionID: types.SessionID(1),
		Owner:     suite.testData.owner,
		Created:   10,
		Expiry:    15,
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	auth.RegisterCodec(cdc)
	types.RegisterLegacyAminoCodec(cdc)

	cdc.Seal()
	return cdc
}
