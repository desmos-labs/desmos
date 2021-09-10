package wasm

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmTypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"
	subspacestypes "github.com/desmos-labs/desmos/x/subspaces/types"
)

type WasmMsgParserInterface interface {
	Parse(contractAddr sdk.AccAddress, msg wasmTypes.CosmosMsg) ([]sdk.Msg, error)
	ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error)
}

type MsgParser struct {
	Parsers map[string]WasmMsgParserInterface
}

func NewMsgParser() MsgParser {
	return MsgParser{
		Parsers: make(map[string]WasmMsgParserInterface),
	}
}

type WasmCustomMsg struct {
	Route   string          `json:"route"`
	MsgData json.RawMessage `json:"msg_data"`
}

const (
	WasmMsgParserRouteProfiles  = profilestypes.ModuleName
	WasmMsgParserRoutePosts     = poststypes.ModuleName
	WasmMsgParserRouteSubspaces = subspacestypes.ModuleName
)

func (p MsgParser) ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var customMsg WasmCustomMsg
	err := json.Unmarshal(data, &customMsg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if parser, ok := p.Parsers[customMsg.Route]; ok {
		return parser.ParseCustom(contractAddr, customMsg.MsgData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, customMsg.Route)
}
