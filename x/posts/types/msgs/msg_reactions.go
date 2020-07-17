package msgs

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/commons"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	postsModels "github.com/desmos-labs/desmos/x/posts/types/models"
)

// MsgRegisterReaction represents the message that must be used when wanting
// to register a new reaction shortCode and the associated value
type MsgRegisterReaction struct {
	ShortCode string         `json:"shortcode" yaml:"shortcode"`
	Value     string         `json:"value" yaml:"value"`
	Subspace  string         `json:"subspace" yaml:"subspace"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
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
func (msg MsgRegisterReaction) Route() string { return postsModels.RouterKey }

// Type should return the action
func (msg MsgRegisterReaction) Type() string { return postsModels.ActionRegisterReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterReaction) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if !postsModels.ShortCodeRegEx.MatchString(msg.ShortCode) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a :")
	}

	if !commons.IsURIValid(msg.Value) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction value should be a valid URL")
	}

	if !postsModels.Sha256RegEx.MatchString(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterReaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}
