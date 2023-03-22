package wasm_test

import (
	"encoding/json"
	"testing"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
	k              keeper.Keeper
	storeKey       storetypes.StoreKey
	ak             authkeeper.AccountKeeper
	paramsKeeper   paramskeeper.Keeper
}

func (suite *TestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, paramstypes.StoreKey, authtypes.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
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

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Define keeper
	suite.paramsKeeper = paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)
	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		suite.paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey, suite.ak, nil)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
