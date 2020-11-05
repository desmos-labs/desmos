package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/app"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc         codec.Marshaler
	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	keeper      keeper.Keeper
	storeKey    sdk.StoreKey
	ms          store.CommitMultiStore
	testData    TestData
}

type TestData struct {
	owner   string
	session types.Session
}

func (suite *KeeperTestSuite) SetupTest() {
	// define store store keys
	magpieKey := sdk.NewKVStoreKey("magpie")
	suite.storeKey = magpieKey

	// create an in-memory db
	memDB := db.NewMemDB()
	suite.ms = store.NewCommitMultiStore(memDB)
	suite.ms.MountStoreWithDB(magpieKey, sdk.StoreTypeIAVL, memDB)
	if err := suite.ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(suite.ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAmino = app.MakeCodecs()
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.storeKey)

	// setup Data
	// nolint - errcheck
	suite.testData.owner = "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
	suite.testData.session = types.Session{
		SessionId:      types.SessionID{Value: 1},
		Owner:          suite.testData.owner,
		CreationTime:   10,
		ExpirationTime: 15,
	}
}

func (suite *KeeperTestSuite) RequireErrorsEqual(expected, actual error) {
	if expected != nil {
		suite.Require().Error(actual)
		suite.Require().Equal(expected.Error(), actual.Error())
	} else {
		suite.Require().NoError(actual)
	}
}
