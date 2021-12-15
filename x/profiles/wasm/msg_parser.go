package wasm

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/cosmwasm"
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
	var msg ProfilesMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse profiles message from contract %s", contractAddr.String())
	}

	switch {
	case msg.SaveProfile != nil:
		return []sdk.Msg{msg.SaveProfile}, msg.SaveProfile.ValidateBasic()
	case msg.DeleteProfile != nil:
		return []sdk.Msg{msg.DeleteProfile}, msg.DeleteProfile.ValidateBasic()
	case msg.RequestDtagTransfer != nil:
		return []sdk.Msg{msg.RequestDtagTransfer}, msg.RequestDtagTransfer.ValidateBasic()
	case msg.AcceptDtagTransferRequest != nil:
		return []sdk.Msg{msg.AcceptDtagTransferRequest}, msg.AcceptDtagTransferRequest.ValidateBasic()
	case msg.RefuseDtagTransferRequest != nil:
		return []sdk.Msg{msg.RefuseDtagTransferRequest}, msg.RefuseDtagTransferRequest.ValidateBasic()
	case msg.CancelDtagTransferRequest != nil:
		return []sdk.Msg{msg.CancelDtagTransferRequest}, msg.CancelDtagTransferRequest.ValidateBasic()
	default:
		return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "Cosmwasm-msg-parser: The msg sent is not one of the supported ones")
	}
}
