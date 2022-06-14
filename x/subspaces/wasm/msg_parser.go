package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct {
	cdc codec.Codec
}

func NewWasmMsgParser(cdc codec.Codec) MsgsParser {
	return MsgsParser{cdc: cdc}
}

func (MsgsParser) Parse(_ sdk.AccAddress, _ wasmvmtypes.CosmosMsg) ([]sdk.Msg, error) {
	return nil, nil
}

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.SubspacesMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/profiles message from contract %s", contractAddr.String())
	}
	switch {
	case msg.CreateSubspace != nil:
		return parser.handleCreateSubspaceRequest(*msg.CreateSubspace)

	case msg.EditSubspace != nil:
		return parser.handleEditSubspaceRequest(*msg.EditSubspace)

	case msg.DeleteSubspace != nil:
		return parser.handleDeleteSubspaceRequest(*msg.DeleteSubspace)

	case msg.CreateUserGroup != nil:
		return parser.handleCreateUserGroupRequest(*msg.CreateUserGroup)

	case msg.SetUserGroupPermissions != nil:
		return parser.handleSetUserGroupPermissionsRequest(*msg.SetUserGroupPermissions)

	case msg.DeleteUserGroup != nil:
		return parser.handleDeleteUserGroupRequest(*msg.DeleteUserGroup)

	case msg.EditUserGroup != nil:
		return parser.handleEditUserGroupRequest(*msg.EditUserGroup)

	case msg.AddUserToUserGroup != nil:
		return parser.handleAddUserToUserGroupRequest(*msg.AddUserToUserGroup)

	case msg.RemoveUserFromUserGroup != nil:
		return parser.handleRemoveUserFromUserGroupRequest(*msg.RemoveUserFromUserGroup)

	case msg.SetUserPermissions != nil:
		return parser.handleSetUserPermissionsRequest(*msg.SetUserPermissions)

	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-subspaces-msg-parser: message not supported")
	}
}

func (parser MsgsParser) handleCreateSubspaceRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgCreateSubspace
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleEditSubspaceRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgEditSubspace
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleDeleteSubspaceRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgDeleteSubspace
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleCreateUserGroupRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgCreateUserGroup
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleSetUserGroupPermissionsRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgSetUserGroupPermissions
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleDeleteUserGroupRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgDeleteUserGroup
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleEditUserGroupRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgEditUserGroup
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleAddUserToUserGroupRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgAddUserToUserGroup
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleRemoveUserFromUserGroupRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgRemoveUserFromUserGroup
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleSetUserPermissionsRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgSetUserPermissions
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}
