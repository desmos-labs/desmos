package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

// ----------------------
// --- MsgCreateAccount
// ----------------------

// MsgCreateAccount defines a CreateAccount message
type MsgCreateAccount struct {
	Name     string         `json:"name,omitempty"`
	Surname  string         `json:"surname,omitempty"`
	Moniker  string         `json:"moniker"`
	Bio      string         `json:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator"`
}

// NewMsgCreateAccount is a constructor function for MsgCreateAccount
func NewMsgCreateAccount(name string, surname string, moniker string, bio string, pictures *Pictures,
	creator sdk.AccAddress) MsgCreateAccount {
	return MsgCreateAccount{
		Name:     name,
		Surname:  surname,
		Moniker:  moniker,
		Bio:      bio,
		Pictures: pictures,
		Creator:  creator,
	}
}

// Route should return the name of the module
func (msg MsgCreateAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateAccount) Type() string { return ActionCreateAccount }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateAccount) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Moniker)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Moniker cannot be blank or empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgEditPost
// ----------------------

// MsgEditPost defines a EditPost message
type MsgEditAccount struct {
	Name     string         `json:"name,omitempty"`
	Surname  string         `json:"surname,omitempty"`
	Moniker  string         `json:"moniker"`
	Bio      string         `json:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator"`
}

// NewMsgEditAccount is a constructor function for MsgEditAccount
func NewMsgEditAccount(name string, surname string, moniker string, bio string, pictures *Pictures,
	creator sdk.AccAddress) MsgEditAccount {
	return MsgEditAccount{
		Name:     name,
		Surname:  surname,
		Moniker:  moniker,
		Bio:      bio,
		Pictures: pictures,
		Creator:  creator,
	}
}

// Route should return the name of the module
func (msg MsgEditAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditAccount) Type() string { return ActionEditAccount }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditAccount) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Moniker)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Moniker cannot be blank or empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgDeleteAccount
// ----------------------

// MsgDeleteAccount defines a DeletePost message
type MsgDeleteAccount struct {
	Moniker string         `json:"moniker"`
	Creator sdk.AccAddress `json:"creator"`
}

// NewMsgDeleteAccount is a constructor function for MsgDeleteAccount
func NewMsgDeleteAccount(moniker string, creator sdk.AccAddress) MsgDeleteAccount {
	return MsgDeleteAccount{
		Moniker: moniker,
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgDeleteAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteAccount) Type() string { return ActionDeleteAccount }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteAccount) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Moniker)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Moniker cannot be blank or empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}
