package wasm_test

import (
	"encoding/json"
	"testing"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	relationshipskeeper "github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
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

	cdc              codec.Codec
	legacyAminoCdc   *codec.LegacyAmino
	ctx              sdk.Context
	storeKey         storetypes.StoreKey
	k                keeper.Keeper
	ak               authkeeper.AccountKeeper
	rk               relationshipskeeper.Keeper
	sk               subspaceskeeper.Keeper
	paramsKeeper     paramskeeper.Keeper
	bankKeeper       bankkeeper.Keeper
	stakingKeeper    stakingkeeper.Keeper
	upgradeKeeper    upgradekeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(
		types.StoreKey, relationshipstypes.StoreKey, subspacestypes.StoreKey,
		authtypes.StoreKey, paramstypes.StoreKey, ibcexported.StoreKey, capabilitytypes.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, memDB)
	}
	for _, memKey := range memKeys {
		ms.MountStoreWithDB(memKey, storetypes.StoreTypeMemory, nil)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.paramsKeeper = paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)

	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)

	suite.capabilityKeeper = capabilitykeeper.NewKeeper(suite.cdc, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	scopedIBCKeeper := suite.capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedProfilesKeeper := suite.capabilityKeeper.ScopeToModule(types.ModuleName)
	suite.IBCKeeper = ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibcexported.StoreKey],
		suite.paramsKeeper.Subspace(ibcexported.ModuleName),
		suite.stakingKeeper,
		suite.upgradeKeeper,
		scopedIBCKeeper,
	)

	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey], nil, nil)
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		suite.storeKey,
		suite.paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.rk,
		suite.IBCKeeper.ChannelKeeper,
		&suite.IBCKeeper.PortKeeper,
		scopedProfilesKeeper,
	)
}
