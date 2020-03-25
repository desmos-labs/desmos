package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

type MsgCreateAccount struct {
	Name             string         `json:"name,omitempty"`
	Surname          string         `json:"surname,omitempty"`
	Moniker          string         `json:"moniker"`
	Bio              string         `json:"bio,omitempty"`
	Pictures         *Pictures      `json:"pictures,omitempty"`
	VerifiedServices []ServiceLink  `json:"verified_services"`
	ChainLinks       []ChainLink    `json:"chain_links"`
	Creator          sdk.AccAddress `json:"creator"`
}

func NewMsgCreateAccount(name string, surname string, moniker string, bio string, pictures *Pictures,
	verifiedServices []ServiceLink, chainLinks []ChainLink, creator sdk.AccAddress) MsgCreateAccount {
	return MsgCreateAccount{
		Name:             name,
		Surname:          surname,
		Moniker:          moniker,
		Bio:              bio,
		Pictures:         pictures,
		VerifiedServices: verifiedServices,
		ChainLinks:       chainLinks,
		Creator:          creator,
	}
}

// Route should return the name of the module
func (msg MsgCreateAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateAccount) Type() string { return ActionCreateAccount }

func (msg MsgCreateAccount) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Moniker)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Moniker cannot be blank or empty")
	}

	for _, serviceLink := range msg.VerifiedServices {
		if err := serviceLink.Validate(); err != nil {
			return err
		}
	}

	for _, chainLink := range msg.ChainLinks {
		if err := chainLink.Validate(); err != nil {
			return err
		}
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

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateAccount) MarshalJSON() ([]byte, error) {
	type temp MsgCreateAccount
	return json.Marshal(temp(msg))
}
