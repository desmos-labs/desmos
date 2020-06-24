package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
func (msg MsgSaveProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSaveProfile) Type() string { return ActionSaveProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgSaveProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
			fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if !DTagRegEx.MatchString(msg.Dtag) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("Invalid profile dtag provided: '%s'", msg.Dtag))
	}

	if msg.Moniker != nil && (len(*msg.Moniker) < MinMonikerLength || len(*msg.Moniker) > MaxMonikerLength) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("Moniker length should be between %d and %d",
				MinMonikerLength, MaxMonikerLength))
	}

	if msg.Bio != nil && len(*msg.Bio) > MaxBioLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("Profile biography cannot exceed %d characters", MaxBioLength))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
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
func (msg MsgDeleteProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteProfile) Type() string { return ActionDeleteProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}
