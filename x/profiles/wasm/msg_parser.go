package wasm

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/cosmwasm"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

var _ cosmwasm.MsgParserInterface = WasmMsgsParser{}

type WasmMsgsParser struct{}

func NewWasmMsgParser() WasmMsgsParser {
	return WasmMsgsParser{}
}

type ProfilesMsg struct {
	SaveProfile   *types.MsgSaveProfile   `json:"save_profile,omitempty"`
	DeleteProfile *types.MsgDeleteProfile `json:"delete_profile,omitempty"`
}

func (WasmMsgsParser) Parse(_ sdk.AccAddress, _ wasmvmtypes.CosmosMsg) ([]sdk.Msg, error) {
	return nil, nil
}

func (WasmMsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
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
	}

	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "")
}
