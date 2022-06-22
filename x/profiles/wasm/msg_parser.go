package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	"github.com/desmos-labs/desmos/v3/x/commons"
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

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.ProfilesMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/profiles message from contract %s", contractAddr.String())
	}

	switch {
	case msg.SaveProfile != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.SaveProfile, &types.MsgSaveProfile{})
	case msg.DeleteProfile != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.DeleteProfile, &types.MsgDeleteProfile{})
	case msg.RequestDtagTransfer != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RequestDtagTransfer, &types.MsgRequestDTagTransfer{})
	case msg.AcceptDtagTransferRequest != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AcceptDtagTransferRequest, &types.MsgAcceptDTagTransferRequest{})
	case msg.RefuseDtagTransferRequest != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RefuseDtagTransferRequest, &types.MsgRefuseDTagTransferRequest{})
	case msg.CancelDtagTransferRequest != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.CancelDtagTransferRequest, &types.MsgCancelDTagTransferRequest{})
	case msg.LinkChainAccount != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.LinkChainAccount, &types.MsgLinkChainAccount{})
	case msg.LinkApplication != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.LinkApplication, &types.MsgLinkApplication{})
	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-profiles-msg-parser: message not supported")
	}
}
