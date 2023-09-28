package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
)

func NewMsgUpdateParams(params Params, authority string) *MsgUpdateParams {
	return &MsgUpdateParams{
		Params:    params,
		Authority: authority,
	}
}

// Route implements legacytx.LegacyMsg
func (msg *MsgUpdateParams) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg *MsgUpdateParams) Type() string {
	return ActionUpdateParams
}

// ValidateBasic implements sdk.Msg
func (msg *MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return err
	}

	return msg.Params.Validate()
}

// GetSignBytes implements sdk.Msg
func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	authority := sdk.MustAccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{authority}
}
