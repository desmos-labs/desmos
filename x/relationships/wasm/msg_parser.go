package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct {
	cdc codec.Codec
}

func NewWasmMsgParser(cdc codec.Codec) MsgsParser {
	return MsgsParser{
		cdc: cdc,
	}
}

func (MsgsParser) Parse(_ sdk.AccAddress, _ wasmvmtypes.CosmosMsg) ([]sdk.Msg, error) {
	return nil, nil
}

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.RelationshipsMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/relationships message from contract %s", contractAddr.String())
	}

	switch {
	case msg.CreateRelationship != nil:
		return parser.handleMsgCreateRelationship(*msg.CreateRelationship)
	case msg.DeleteRelationship != nil:
		return parser.handleMsgDeleteRelationship(*msg.DeleteRelationship)
	case msg.BlockUser != nil:
		return parser.handleMsgBlockUser(*msg.BlockUser)
	case msg.UnblockUser != nil:
		return parser.handleMsgUnblockUser(*msg.UnblockUser)

	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-relationships-msg-parser: message not supported")
	}
}

func (parser MsgsParser) handleMsgCreateRelationship(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgCreateRelationship
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgDeleteRelationship(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgDeleteRelationship
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgBlockUser(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgBlockUser
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgUnblockUser(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgUnblockUser
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}
