package wasm_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/desmos-labs/desmos/v3/app"
	"github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"
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

func buildSubspacesQueryRequest(cdc codec.Codec, query *types.QuerySubspacesRequest) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesQuery{Subspaces: cdc.MustMarshalJSON(query)})
	return bz
}

func buildSubspaceQueryRequest(cdc codec.Codec, query *types.QuerySubspaceRequest) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesQuery{Subspace: cdc.MustMarshalJSON(query)})
	return bz
}

func buildUserGroupsQueryRequest(cdc codec.Codec, query *types.QueryUserGroupsRequest) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesQuery{UserGroups: cdc.MustMarshalJSON(query)})
	return bz
}

func buildUserGroupQueryRequest(cdc codec.Codec, query *types.QueryUserGroupRequest) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesQuery{UserGroup: cdc.MustMarshalJSON(query)})
	return bz
}

func buildUserGroupMembersQueryRequest(cdc codec.Codec, query *types.QueryUserGroupMembersRequest) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesQuery{UserGroupMembers: cdc.MustMarshalJSON(query)})
	return bz
}

func buildUserPermissionsQueryRequest(cdc codec.Codec, query *types.QueryUserPermissionsRequest) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesQuery{UserPermissions: cdc.MustMarshalJSON(query)})
	return bz
}

type Testsuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              keeper.Keeper
	paramsKeeper   paramskeeper.Keeper
	storeKey       sdk.StoreKey
}

func (suite *Testsuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, paramstypes.StoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Define keeper
	suite.k = keeper.NewKeeper(suite.cdc, suite.storeKey)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(Testsuite))
}
