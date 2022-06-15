package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	"github.com/desmos-labs/desmos/v3/x/commons"
	"github.com/desmos-labs/desmos/v3/x/posts/types"
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
	var msg types.PostsMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/posts message from contract %s", contractAddr.String())
	}

	switch {
	case msg.CreatePost != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.CreatePost, &types.MsgCreatePost{})
	case msg.EditPost != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.EditPost, &types.MsgEditPost{})
	case msg.DeletePost != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.DeletePost, &types.MsgDeletePost{})
	case msg.AddPostAttachment != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AddPostAttachment, &types.MsgAddPostAttachment{})
	case msg.RemovePostAttachment != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RemovePostAttachment, &types.MsgRemovePostAttachment{})
	case msg.AnswerPoll != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AnswerPoll, &types.MsgAnswerPoll{})
	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-posts-msg-parser: message not supported")
	}
}
