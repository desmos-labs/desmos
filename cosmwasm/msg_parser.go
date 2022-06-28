package cosmwasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"
	reactionstypes "github.com/desmos-labs/desmos/v4/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v4/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
	Profiles      *json.RawMessage `json:"profiles"`
	Subspaces     *json.RawMessage `json:"subspaces"`
	Relationships *json.RawMessage `json:"relationships"`
	Posts         *json.RawMessage `json:"posts"`
	Reports       *json.RawMessage `json:"reports"`
	Reactions     *json.RawMessage `json:"reactions"`
}

func (router ParserRouter) ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var customMsg CustomMsg
	err := json.Unmarshal(data, &customMsg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	// get route and msg from data
	var route string
	var msg json.RawMessage
	if customMsg.Profiles != nil {
		route = WasmMsgParserRouteProfiles
		msg = *customMsg.Profiles
	}
	if customMsg.Subspaces != nil {
		route = QueryRouteSubspaces
		msg = *customMsg.Subspaces
	}
	if customMsg.Relationships != nil {
		route = WasmMsgParserRouteRelationships
		msg = *customMsg.Relationships
	}
	if customMsg.Posts != nil {
		route = WasmMsgParserRoutePosts
		msg = *customMsg.Posts
	}
	if customMsg.Reports != nil {
		route = WasmMsgParserRouteReports
		msg = *customMsg.Reports
	}
	if customMsg.Reactions != nil {
		route = WasmMsgParserRouteReactions
		msg = *customMsg.Reactions
	}

	if parser, ok := router.Parsers[route]; ok {
		return parser.ParseCustomMsgs(contractAddr, msg)
	}
	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "unimplemented route")
}
