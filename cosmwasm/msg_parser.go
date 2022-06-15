package cosmwasm

import (
	"encoding/json"
	"log"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	relationshipstypes "github.com/desmos-labs/desmos/v3/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

const (
	WasmMsgParserRouteProfiles      = profilestypes.ModuleName
	WasmMsgParserRouteSubspaces     = subspacestypes.ModuleName
	WasmMsgParserRouteRelationships = relationshipstypes.ModuleName
	WasmMsgParserRouteReports       = reportstypes.ModuleName
)

type MsgParserInterface interface {
	Parse(contractAddr sdk.AccAddress, msg wasmvmtypes.CosmosMsg) ([]sdk.Msg, error)
	ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error)
}

type ParserRouter struct {
	Parsers map[string]MsgParserInterface
}

func NewParserRouter() ParserRouter {
	return ParserRouter{
		Parsers: make(map[string]MsgParserInterface),
	}
}

type CustomMsg struct {
	Route   string          `json:"route"`
	MsgData json.RawMessage `json:"msg_data"`
}

func (router ParserRouter) ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var customMsg CustomMsg
	err := json.Unmarshal(data, &customMsg)

	log.Println("[!] CosmWasm contract msg routed to module: ", customMsg.Route)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if parser, ok := router.Parsers[customMsg.Route]; ok {
		return parser.ParseCustomMsgs(contractAddr, customMsg.MsgData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, customMsg.Route)
}
