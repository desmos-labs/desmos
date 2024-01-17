package wasm_test

import (
	"encoding/json"
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func buildCreateSubspaceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{CreateSubspace: &raw})
	return bz
}

func buildEditSubspaceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{EditSubspace: &raw})
	return bz
}

func buildDeleteSubspaceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{DeleteSubspace: &raw})
	return bz
}

func buildCreateUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{CreateUserGroup: &raw})
	return bz
}

func buildEditUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{EditUserGroup: &raw})
	return bz
}

func buildSetUserGroupPermissionsRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{SetUserGroupPermissions: &raw})
	return bz
}

func buildDeleteUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{DeleteUserGroup: &raw})
	return bz
}

func buildAddUserToGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{AddUserToUserGroup: &raw})
	return bz
}

func buildRemoveUserFromUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{RemoveUserFromUserGroup: &raw})
	return bz
}

func buildSetUserPermissionsRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{SetUserPermissions: &raw})
	return bz
}

func buildGrantTreasuryAuthorizationRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{GrantTreasuryAuthorization: &raw})
	return bz
}

func buildRevokeTreasuryAuthorizationRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{RevokeTreasuryAuthorization: &raw})
	return bz
}

func buildGrantAllowanceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{GrantAllowance: &raw})
	return bz
}

func buildRevokeAllowanceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(msg))
	bz, _ := json.Marshal(types.SubspacesMsg{RevokeAllowance: &raw})
	return bz
}

func buildSubspacesQueryRequest(cdc codec.Codec, query *types.QuerySubspacesRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{Subspaces: &raw})
	return bz
}

func buildSubspaceQueryRequest(cdc codec.Codec, query *types.QuerySubspaceRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{Subspace: &raw})
	return bz
}

func buildUserGroupsQueryRequest(cdc codec.Codec, query *types.QueryUserGroupsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{UserGroups: &raw})
	return bz
}

func buildUserGroupQueryRequest(cdc codec.Codec, query *types.QueryUserGroupRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{UserGroup: &raw})
	return bz
}

func buildUserGroupMembersQueryRequest(cdc codec.Codec, query *types.QueryUserGroupMembersRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{UserGroupMembers: &raw})
	return bz
}

func buildUserPermissionsQueryRequest(cdc codec.Codec, query *types.QueryUserPermissionsRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{UserPermissions: &raw})
	return bz
}

func buildUserAllowancesQueryRequest(cdc codec.Codec, query *types.QueryUserAllowancesRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{UserAllowances: &raw})
	return bz
}

func buildGroupAllowancesQueryRequest(cdc codec.Codec, query *types.QueryGroupAllowancesRequest) json.RawMessage {
	raw := json.RawMessage(cdc.MustMarshalJSON(query))
	bz, _ := json.Marshal(types.SubspacesQuery{GroupAllowances: &raw})
	return bz
}

type TestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              *keeper.Keeper
	storeKey       storetypes.StoreKey
	ak             authkeeper.AccountKeeper
}

func (suite *TestSuite) SetupTest() {
	// Define store keys
	keys := storetypes.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Define keeper
	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		address.NewBech32Codec("cosmos"),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.ak, nil, "authority")
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
