package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v2/app"
	"github.com/desmos-labs/desmos/v2/x/fees/keeper"
	"github.com/desmos-labs/desmos/v2/x/fees/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.BinaryCodec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	keeper         keeper.Keeper
	paramsKeeper   paramskeeper.Keeper
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

	suite.ctx = sdk.NewContext(
		ms,
		tmproto.Header{ChainID: "test-chain-id"},
		false,
		log.NewNopLogger(),
	)

	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()
	suite.paramsKeeper = paramskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, paramsKey, paramsTKey)
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.paramsKeeper.Subspace(types.DefaultParamspace))
}
