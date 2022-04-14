package wasm_test

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func buildCreateSubspaceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{CreateSubspace: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildEditSubspaceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{EditSubspace: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildDeleteSubspaceRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{DeleteSubspace: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildCreateUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{CreateUserGroup: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildEditUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{EditUserGroup: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildSetUserGroupPermissionsRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{SetUserGroupPermissions: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildDeleteUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{DeleteUserGroup: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildAddUserToGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{AddUserToUserGroup: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildRemoveUserFromUserGroupRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{RemoveUserFromUserGroup: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildSetUserPermissionsRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.SubspacesMsg{SetUserPermissions: cdc.MustMarshalJSON(msg)})
	return bz
}
