package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	"github.com/desmos-labs/desmos/v3/x/commons"
	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct {
	cdc codec.Codec
}

func NewWasmMsgParser(cdc codec.Codec) MsgsParser {
	return MsgsParser{cdc: cdc}
}

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.ReactionsMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/reactions message from contract %s", contractAddr.String())
	}
	switch {
	case msg.AddReaction != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AddReaction, &types.MsgAddReaction{})
	case msg.RemoveReaction != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RemoveReaction, &types.MsgRemoveReaction{})
	case msg.AddRegisteredReaction != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AddRegisteredReaction, &types.MsgAddRegisteredReaction{})
	case msg.EditRegisteredReaction != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.EditRegisteredReaction, &types.MsgEditRegisteredReaction{})
	case msg.RemoveRegisteredReaction != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RemoveRegisteredReaction, &types.MsgRemoveRegisteredReaction{})
	case msg.SetReactionsParams != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.SetReactionsParams, &types.MsgSetReactionsParams{})
	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-reactions-msg-parser: message not supported")
	}
}
