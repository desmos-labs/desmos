package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Route should return the name of the module
func (msg MsgCreateLink) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateLink) Type() string { return ActionCreateLink }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateLink) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateLink) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners defines whose signature is required
func (msg *MsgCreateLink) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateLink) MarshalJSON() ([]byte, error) {
	type temp MsgCreateLink
	return json.Marshal(temp(msg))
}
