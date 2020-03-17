package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgRegisterReaction represents the message that must be used when wanting
// to register a new reaction shortCode and the associated value
type MsgRegisterReaction struct {
	ShortCode string         `json:"shortcode"`
	Value     string         `json:"value"`
	Subspace  string         `json:"subspace"`
	Creator   sdk.AccAddress `json:"creator"`
}

// NewMsgRegisterReaction is a constructor function for MsgRegisterReaction
func NewMsgRegisterReaction(creator sdk.AccAddress, shortCode, value, subspace string) MsgRegisterReaction {
	return MsgRegisterReaction{
		ShortCode: shortCode,
		Value:     value,
		Subspace:  subspace,
		Creator:   creator,
	}
}

// Route should return the name of the module
func (msg MsgRegisterReaction) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterReaction) Type() string { return ActionRegisterReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterReaction) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.ShortCode)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "reaction short code cannot empty or blank")
	}

	if len(strings.TrimSpace(msg.Value)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "reaction value cannot empty or blank")
	}

	if !ShortCodeRegEx.MatchString(msg.ShortCode) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "reaction short code must be an emoji short code")
	}

	if !URIRegEx.MatchString(msg.Value) || !UnicodeRegEx.MatchString(msg.Value) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "reaction value should be a URL or an emoji unicode")
	}

	if !SubspaceRegEx.MatchString(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "reaction subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterReaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}
