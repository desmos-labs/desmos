package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/fees/keeper"
	"github.com/desmos-labs/desmos/x/fees/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
	"testing"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc          *codec.Codec
	ctx          sdk.Context
	keeper       keeper.Keeper
	paramsKeeper params.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	//define store keys
	paramsKey := sdk.NewKVStoreKey("params")
	paramsTKey := sdk.NewTransientStoreKey("transient_params")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, memDB)
	ms.MountStoreWithDB(paramsTKey, sdk.StoreTypeTransient, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc = testCodec()
	suite.paramsKeeper = params.NewKeeper(suite.cdc, paramsKey, paramsTKey)
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.paramsKeeper.Subspace(types.DefaultParamspace))
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.Seal()
	return cdc
}
