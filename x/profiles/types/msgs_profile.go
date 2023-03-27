package types

import (
	"fmt"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// NewMsgSaveProfile returns a new MsgSaveProfile instance
func NewMsgSaveProfile(dTag string, nickname, bio, profilePic, coverPic string, creator string) *MsgSaveProfile {
	return &MsgSaveProfile{
		DTag:           dTag,
		Nickname:       nickname,
		Bio:            bio,
		ProfilePicture: profilePic,
		CoverPicture:   coverPic,
		Creator:        creator,
	}
}

// Route should return the name of the module
func (msg MsgSaveProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSaveProfile) Type() string { return ActionSaveProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgSaveProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaveProfile) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgDeleteProfile is a constructor function for MsgDeleteProfile
func NewMsgDeleteProfile(creator string) *MsgDeleteProfile {
	return &MsgDeleteProfile{
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteProfile) Type() string { return ActionDeleteProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}
