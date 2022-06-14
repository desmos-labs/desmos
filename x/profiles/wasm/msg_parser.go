package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
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
	var msg types.ProfilesMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/profiles message from contract %s", contractAddr.String())
	}

	switch {
	case msg.SaveProfile != nil:
		return parser.handleMsgSaveProfile(*msg.SaveProfile)
	case msg.DeleteProfile != nil:
		return parser.handleMsgDeleteProfile(*msg.DeleteProfile)
	case msg.RequestDtagTransfer != nil:
		return parser.handleMsgRequestDTagTransfer(*msg.RequestDtagTransfer)
	case msg.AcceptDtagTransferRequest != nil:
		return parser.handleMsgAcceptDTagTransferRequest(*msg.AcceptDtagTransferRequest)
	case msg.RefuseDtagTransferRequest != nil:
		return parser.handleMsgRefuseDTagTransferRequest(*msg.RefuseDtagTransferRequest)
	case msg.CancelDtagTransferRequest != nil:
		return parser.handleMsgCancelDTagTransferRequest(*msg.CancelDtagTransferRequest)
	case msg.LinkChainAccount != nil:
		return parser.handleMsgLinkChainAccount(*msg.LinkChainAccount)
	case msg.LinkApplication != nil:
		return parser.handleMsgLinkApplication(*msg.LinkApplication)
	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-profiles-msg-parser: message not supported")
	}
}

func (parser MsgsParser) handleMsgSaveProfile(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgSaveProfile
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgDeleteProfile(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgDeleteProfile
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgRequestDTagTransfer(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgRequestDTagTransfer
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgAcceptDTagTransferRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgAcceptDTagTransferRequest
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgCancelDTagTransferRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgCancelDTagTransferRequest
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgRefuseDTagTransferRequest(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgRefuseDTagTransferRequest
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgLinkChainAccount(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgLinkChainAccount
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}

func (parser MsgsParser) handleMsgLinkApplication(data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.MsgLinkApplication
	err := parser.cdc.UnmarshalJSON(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{&msg}, msg.ValidateBasic()
}
