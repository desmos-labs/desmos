package cosmwasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ cosmwasm_plugins.MsgParserInterface = MsgParser{}

type MsgParserInterface interface {
	Parse(contractAddr sdk.AccAddress, msg wasmvmtypes.CosmosMsg) ([]sdk.Msg, error)
	ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error)
}

type ParsersRouter struct {
	Parser map[string]MsgParserInterface
}

func NewParserRouter() ParsersRouter {
	return ParsersRouter{
		Parser: make(map[string]MsgParserInterface),
	}
}

type CustomMsg struct {
	Route   string          `json:"route"`
	MsgData json.RawMessage `json:"msg_data"`
}

func (router ParsersRouter) ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var customMsg CustomMsg
	err := json.Unmarshal(data, &customMsg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if parser, ok := router.Parser[customMsg.Route]; ok {
		return parser.ParseCustomMsgs(contractAddr, customMsg.MsgData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, customMsg.Route)
}
