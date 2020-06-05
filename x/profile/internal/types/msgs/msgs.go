package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// MsgSaveProfile defines a SaveProfile message
type MsgSaveProfile struct {
	Moniker    string         `json:"moniker" yaml:"moniker"`
	Name       *string        `json:"name,omitempty" yaml:"name,omitempty"`
	Surname    *string        `json:"surname,omitempty" yaml:"surname,omitempty"`
	Bio        *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	ProfilePic *string        `json:"profile_pic,omitempty" yaml:"profile_pic,omitempty"`
	ProfileCov *string        `json:"profile_cov,omitempty" yaml:"profile_cov,omitempty"`
	Creator    sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgSaveProfile is a constructor function for MsgSaveProfile
func NewMsgSaveProfile(moniker string, name, surname, bio, profilePic,
	profileCov *string, creator sdk.AccAddress) MsgSaveProfile {
	return MsgSaveProfile{
		Moniker:    moniker,
		Name:       name,
		Surname:    surname,
		Bio:        bio,
		ProfilePic: profilePic,
		ProfileCov: profileCov,
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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	// Todo remove checks
	if len(msg.Moniker) < models.MinMonikerLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile moniker cannot be less than %d characters", models.MinMonikerLength))
	}

	if len(msg.Moniker) > models.MaxMonikerLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile moniker cannot exceed %d characters", models.MaxMonikerLength))
	}

	if msg.Name != nil {
		if len(*msg.Name) < models.MinNameSurnameLength {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile name cannot be less than %d characters", models.MinNameSurnameLength))
		}

		if len(*msg.Name) > models.MaxNameSurnameLength {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile name cannot exceed %d characters", models.MaxNameSurnameLength))
		}
	}

	if msg.Surname != nil {
		if msg.Surname != nil && len(*msg.Surname) < models.MinNameSurnameLength {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile surname cannot be less than %d characters", models.MinNameSurnameLength))
		}

		if len(*msg.Surname) > models.MaxNameSurnameLength {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile surname cannot exceed %d characters", models.MaxNameSurnameLength))
		}
	}

	if msg.Bio != nil && len(*msg.Bio) > models.MaxBioLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile biography cannot exceed %d characters", models.MaxBioLength))
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
