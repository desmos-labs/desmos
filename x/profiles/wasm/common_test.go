package wasm_test

import (
	"encoding/json"
	"testing"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func buildSaveProfileRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{SaveProfile: &raw})
	return bz
}

func buildDeleteProfileRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{DeleteProfile: &raw})
	return bz
}

func buildRequestDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{RequestDtagTransfer: &raw})
	return bz
}

func buildAcceptDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{AcceptDtagTransferRequest: &raw})
	return bz
}

func buildRefuseDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{RefuseDtagTransferRequest: &raw})
	return bz
}

func buildCancelDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{CancelDtagTransferRequest: &raw})
	return bz
}

func buildLinkChainAccountRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{LinkChainAccount: &raw})
	return bz
}

func buildUnlinkChainAccountRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{UnlinkChainAccount: &raw})
	return bz
}

func buildSetDefaultExternalAddressRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{SetDefaultExternalAddress: &raw})
	return bz
}

func buildLinkApplicationRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{LinkApplication: &raw})
	return bz
}

func buildUnlinkApplicationRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.ProfilesMsg{UnlinkApplication: &raw})
	return bz
}

func buildProfileQueryRequest(cdc codec.Codec, query *types.QueryProfileRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{Profile: &raw})
	return bz
}

func buildIncomingDtagTransferQueryRequest(cdc codec.Codec, query *types.QueryIncomingDTagTransferRequestsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{IncomingDtagTransferRequests: &raw})
	return bz
}

func buildChainLinksQueryRequest(cdc codec.Codec, query *types.QueryChainLinksRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{ChainLinks: &raw})
	return bz
}

func buildChainLinkOwnersQueryRequest(cdc codec.Codec, query *types.QueryChainLinkOwnersRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{ChainLinkOwners: &raw})
	return bz
}

func buildDefaultExternalAddressesQueryRequest(cdc codec.Codec, query *types.QueryDefaultExternalAddressesRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{DefaultExternalAddresses: &raw})
	return bz
}

func buildAppLinksQueryRequest(cdc codec.Codec, query *types.QueryApplicationLinksRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{ApplicationLinks: &raw})
	return bz
}

func buildApplicationLinkByClientIDQueryRequest(cdc codec.Codec, query *types.QueryApplicationLinkByClientIDRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{ApplicationLinkByClientID: &raw})
	return bz
}

func buildApplicationLinkOwnersQueryRequest(cdc codec.Codec, query *types.QueryApplicationLinkOwnersRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.ProfilesQuery{ApplicationLinkOwners: &raw})
	return bz
}

type TestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       storetypes.StoreKey
	k              *keeper.Keeper
	ak             authkeeper.AccountKeeper
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, types.StoreKey)
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

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		suite.storeKey,
		suite.ak,
		nil,
		nil,
		nil,
		nil,
		authtypes.NewModuleAddress("gov").String(),
	)
}
