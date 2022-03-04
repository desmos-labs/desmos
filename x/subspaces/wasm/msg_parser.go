package wasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/cosmwasm"
	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct{}

func NewWasmMsgParser() MsgsParser {
	return MsgsParser{}
}

func (MsgsParser) Parse(_ sdk.AccAddress, _ wasmvmtypes.CosmosMsg) ([]sdk.Msg, error) {
	return nil, nil
}

func (MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var route types.SubspacesMsgRoute
	err := json.Unmarshal(data, &route)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse x/profiles message from contract %s", contractAddr.String())
	}
	msg := route.Msg
	switch {
	case msg.CreateSubspace != nil:
		return []sdk.Msg{msg.CreateSubspace}, msg.CreateSubspace.ValidateBasic()
	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "CosmWasm-msg-parser: The msg sent is not one of the supported ones")
	}
}
