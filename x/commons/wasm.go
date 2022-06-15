package commons

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
)

// HandleWasmMsg deserialises the given sdk.Msg and checks whether it is valid or not
func HandleWasmMsg(cdc codec.Codec, data json.RawMessage, msg sdk.Msg) ([]sdk.Msg, error) {
	err := cdc.UnmarshalJSON(data, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return []sdk.Msg{msg}, msg.ValidateBasic()
}
