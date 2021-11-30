package wasm

import (
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

type WasmMsgsParser struct{}

func NewWasmMsgsParser() WasmMsgsParser {
	return WasmMsgsParser{}
}

type DesmosMsgs struct {
	SaveProfile   *types.MsgSaveProfile   `json:"save_profile,omitempty"`
	DeleteProfile *types.MsgDeleteProfile `json:"delete_profile,omitempty"`
}

func (WasmMsgsParser) Parse(_ sdk.AccAddress, _ wasmvmtypes.CosmosMsg) ([]sdk.Msg, error) {
	return nil, nil
}

func (WasmMsgsParser) ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	return nil, sdkerrors.Wrap(wasm.ErrInvalidMsg, "")
}
