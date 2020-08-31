package msgs

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// MsgSaveProfile defines a SaveProfile message
type MsgSaveProfile struct {
	Dtag       string         `json:"dtag" yaml:"dtag"`
	Moniker    *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio        *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	ProfilePic *string        `json:"profile_picture,omitempty" yaml:"profile_pic,omitempty"`
	CoverPic   *string        `json:"cover_picture,omitempty" yaml:"cover_pic,omitempty"`
	Creator    sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgSaveProfile is a constructor function for MsgSaveProfile
func NewMsgSaveProfile(dtag string, moniker, bio, profilePic, coverPic *string, creator sdk.AccAddress) MsgSaveProfile {
	return MsgSaveProfile{
		Dtag:       dtag,
		Moniker:    moniker,
		Bio:        bio,
		ProfilePic: profilePic,
		CoverPic:   coverPic,
		Creator:    creator,
	}
}

// Route should return the name of the module
func (msg MsgSaveProfile) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgSaveProfile) Type() string { return models.ActionSaveProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgSaveProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", msg.Creator))
	}

	if strings.TrimSpace(msg.Dtag) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaveProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgDeleteProfile
// ----------------------

// MsgDeleteProfile defines a DeleteProfile message
type MsgDeleteProfile struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgDeleteProfile is a constructor function for MsgDeleteProfile
func NewMsgDeleteProfile(creator sdk.AccAddress) MsgDeleteProfile {
	return MsgDeleteProfile{
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProfile) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgDeleteProfile) Type() string { return models.ActionDeleteProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}
