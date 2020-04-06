package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgCreateProfile
// ----------------------

// MsgCreateProfile defines a CreateProfile message
type MsgCreateProfile struct {
	Moniker  string         `json:"moniker"`
	Name     string         `json:"name,omitempty"`
	Surname  string         `json:"surname,omitempty"`
	Bio      string         `json:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator"`
}

// NewMsgCreateProfile is a constructor function for MsgCreateProfile
func NewMsgCreateProfile(name string, surname string, moniker string, bio string, pictures *Pictures,
	creator sdk.AccAddress) MsgCreateProfile {
	return MsgCreateProfile{
		Moniker:  moniker,
		Name:     name,
		Surname:  surname,
		Bio:      bio,
		Pictures: pictures,
		Creator:  creator,
	}
}

// Route should return the name of the module
func (msg MsgCreateProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateProfile) Type() string { return ActionCreateProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(msg.Name) != 0 && len(msg.Name) < MinNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile name cannot be less than %d characters", MinNameSurnameLength))
	}

	if len(msg.Name) > MaxNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile name cannot exceed %d characters", MaxNameSurnameLength))
	}

	if len(msg.Surname) != 0 && len(msg.Surname) < MinNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile surname cannot be less than %d characters", MinNameSurnameLength))
	}

	if len(msg.Surname) > MaxNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile surname cannot exceed %d characters", MaxNameSurnameLength))
	}

	if len(strings.TrimSpace(msg.Moniker)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile moniker cannot be blank or empty")
	}

	if len(msg.Moniker) > MaxMonikerLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile moniker cannot exceed %d characters", MaxMonikerLength))
	}

	if len(msg.Bio) > MaxBioLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile biography cannot exceed %d characters", MaxBioLength))
	}

	if msg.Pictures != nil {
		if err := msg.Pictures.Validate(); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgEditProfile
// ----------------------

// MsgEditProfile defines a EditProfile message
type MsgEditProfile struct {
	NewMoniker string         `json:"new_moniker"`
	Name       string         `json:"name,omitempty"`
	Surname    string         `json:"surname,omitempty"`
	Bio        string         `json:"bio,omitempty"`
	ProfilePic string         `json:"profile_pic,omitempty"`
	ProfileCov string         `json:"profile_cov,omitempty"`
	Creator    sdk.AccAddress `json:"creator"`
}

// NewMsgEditProfile is a constructor function for MsgEditProfile
func NewMsgEditProfile(newMoniker string, name string, surname string, bio string, profilePic string,
	profileCov string, creator sdk.AccAddress) MsgEditProfile {
	return MsgEditProfile{
		NewMoniker: newMoniker,
		Name:       name,
		Surname:    surname,
		Bio:        bio,
		ProfilePic: profilePic,
		ProfileCov: profileCov,
		Creator:    creator,
	}
}

// Route should return the name of the module
func (msg MsgEditProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditProfile) Type() string { return ActionEditProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(msg.NewMoniker) > MaxMonikerLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile new moniker cannot exceed %d characters", MaxMonikerLength))
	}

	if len(msg.Name) != 0 && len(msg.Name) < MinNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile name cannot be less than %d characters", MinNameSurnameLength))
	}

	if len(msg.Name) > MaxNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile name cannot exceed %d characters", MaxNameSurnameLength))
	}

	if len(msg.Surname) != 0 && len(msg.Surname) < MinNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile surname cannot be less than %d characters", MinNameSurnameLength))
	}

	if len(msg.Surname) > MaxNameSurnameLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile surname cannot exceed %d characters", MaxNameSurnameLength))
	}

	if len(msg.Bio) > MaxBioLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile biography cannot exceed %d characters", MaxBioLength))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgDeleteProfile
// ----------------------

// MsgDeleteProfile defines a DeleteProfile message
type MsgDeleteProfile struct {
	Creator sdk.AccAddress `json:"creator"`
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
