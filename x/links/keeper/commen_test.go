package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctesting "github.com/cosmos/cosmos-sdk/x/ibc/testing"
	"github.com/desmos-labs/desmos/x/links/keeper"
	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/suite"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	queryClient types.QueryClient

	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       sdk.StoreKey
	k              keeper.Keeper
	testData       TestData
}

type TestData struct {
	user      string
	otherUser string
	link      types.Link
}

func (suite *KeeperTestSuite) SetupTest() {
	// ibc setting
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), suite.chainA.App.InterfaceRegistry())
	suite.queryClient = types.NewQueryClient(queryHelper)

	// define store keys
	linkKey := sdk.NewKVStoreKey(types.StoreKey)
	suite.storeKey = linkKey

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(linkKey, sdk.StoreTypeIAVL, memDB)

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
