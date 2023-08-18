package wasm_test

import (
	"encoding/json"
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/reports/keeper"
	"github.com/desmos-labs/desmos/v6/x/reports/types"
)

func buildCreateReportRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReportsMsg{CreateReport: &raw})
	return bz
}

func buildDeleteReportRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReportsMsg{DeleteReport: &raw})
	return bz
}

func buildSupportStandardReasonRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReportsMsg{SupportStandardReason: &raw})
	return bz
}

func buildAddReasonRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReportsMsg{AddReason: &raw})
	return bz
}

func buildRemoveReasonRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ReportsMsg{RemoveReason: &raw})
	return bz
}

func buildReportsQueryRequest(cdc codec.Codec, query *types.QueryReportsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReportsQuery{Reports: &raw})
	return bz
}

func buildReportQueryRequest(cdc codec.Codec, query *types.QueryReportRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReportsQuery{Report: &raw})
	return bz
}

func buildReasonsQueryRequest(cdc codec.Codec, query *types.QueryReasonsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReportsQuery{Reasons: &raw})
	return bz
}

func buildReasonQueryRequest(cdc codec.Codec, query *types.QueryReasonRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ReportsQuery{Reason: &raw})
	return bz
}

type Testsuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey storetypes.StoreKey
	k        keeper.Keeper
}

func (suite *Testsuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Define keeper
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		nil,
		nil,
		nil,
		nil,
		authtypes.NewModuleAddress("gov").String(),
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(Testsuite))
}
