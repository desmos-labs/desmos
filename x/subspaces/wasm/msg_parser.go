package wasm

import (
	"encoding/json"

	"cosmossdk.io/errors"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/cosmwasm"
	"github.com/desmos-labs/desmos/v5/x/commons"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct {
	cdc codec.Codec
}

func NewWasmMsgParser(cdc codec.Codec) MsgsParser {
	return MsgsParser{cdc: cdc}
}

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.SubspacesMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse x/subspaces message from contract %s", contractAddr.String())
	}
	switch {
	case msg.CreateSubspace != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.CreateSubspace, &types.MsgCreateSubspace{})
	case msg.EditSubspace != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.EditSubspace, &types.MsgEditSubspace{})
	case msg.DeleteSubspace != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.DeleteSubspace, &types.MsgDeleteSubspace{})
	case msg.CreateUserGroup != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.CreateUserGroup, &types.MsgCreateUserGroup{})
	case msg.SetUserGroupPermissions != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.SetUserGroupPermissions, &types.MsgSetUserGroupPermissions{})
	case msg.DeleteUserGroup != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.DeleteUserGroup, &types.MsgDeleteUserGroup{})
	case msg.EditUserGroup != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.EditUserGroup, &types.MsgEditUserGroup{})
	case msg.AddUserToUserGroup != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AddUserToUserGroup, &types.MsgAddUserToUserGroup{})
	case msg.RemoveUserFromUserGroup != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RemoveUserFromUserGroup, &types.MsgRemoveUserFromUserGroup{})
	case msg.SetUserPermissions != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.SetUserPermissions, &types.MsgSetUserPermissions{})
	case msg.GrantTreasuryAuthorization != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.GrantTreasuryAuthorization, &types.MsgGrantTreasuryAuthorization{})
	case msg.RevokeTreasuryAuthorization != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RevokeTreasuryAuthorization, &types.MsgRevokeTreasuryAuthorization{})
	case msg.GrantAllowance != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.GrantAllowance, &types.MsgGrantAllowance{})
	case msg.RevokeAllowance != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RevokeAllowance, &types.MsgRevokeAllowance{})
	default:
		return nil, errors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-subspaces-msg-parser: message not supported")
	}
}
