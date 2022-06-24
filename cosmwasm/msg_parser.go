package cosmwasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	reactionstypes "github.com/desmos-labs/desmos/v3/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v3/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

const (
	WasmMsgParserRouteProfiles      = profilestypes.ModuleName
	WasmMsgParserRouteSubspaces     = subspacestypes.ModuleName
	WasmMsgParserRouteRelationships = relationshipstypes.ModuleName
	WasmMsgParserRoutePosts         = poststypes.ModuleName
	WasmMsgParserRouteReports       = reportstypes.ModuleName
	WasmMsgParserRouteReactions     = reactionstypes.ModuleName
)

type MsgParserInterface interface {
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

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if parser, ok := router.Parsers[customMsg.Route]; ok {
		return parser.ParseCustomMsgs(contractAddr, customMsg.MsgData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, customMsg.Route)
}
