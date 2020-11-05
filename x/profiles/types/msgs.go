package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// NewMsgSaveProfile returns a new MsgSaveProfile instance
func NewMsgSaveProfile(dtag string, moniker, bio, profilePic, coverPic string, creator string) *MsgSaveProfile {
	return &MsgSaveProfile{
		Dtag:           dtag,
		Moniker:        moniker,
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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}

	if strings.TrimSpace(msg.Dtag) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRequestDTagTransfer is a constructor function for MsgRequestDtagTransfer
func NewMsgRequestDTagTransfer(sender, receiver string) *MsgRequestDTagTransfer {
	return &MsgRequestDTagTransfer{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRequestDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRequestDTagTransfer) Type() string { return ActionRequestDtag }

func (msg MsgRequestDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRequestDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRequestDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgCancelDTagTransferRequest is a constructor for MsgCancelDTagTransfer
func NewMsgCancelDTagTransferRequest(sender, receiver string) *MsgCancelDTagTransfer {
	return &MsgCancelDTagTransfer{
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgCancelDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCancelDTagTransfer) Type() string { return CancelDTagTransferRequest }

func (msg MsgCancelDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver == msg.Sender {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCancelDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgAcceptDTagTransfer is a constructor for MsgAcceptDTagTransfer
func NewMsgAcceptDTagTransfer(newDTag string, sender, receiver string) *MsgAcceptDTagTransfer {
	return &MsgAcceptDTagTransfer{
		NewDtag:  newDTag,
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgAcceptDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAcceptDTagTransfer) Type() string { return ActionAcceptDtagTransfer }

func (msg MsgAcceptDTagTransfer) ValidateBasic() error {
	if strings.TrimSpace(msg.NewDtag) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new DTag can't be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", msg.Sender)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address: %s", msg.Receiver)
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRefuseDTagTransferRequest is a constructor for MsgRefuseDTagTransfer
func NewMsgRefuseDTagTransferRequest(sender, receiver string) *MsgRefuseDTagTransfer {
	return &MsgRefuseDTagTransfer{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRefuseDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRefuseDTagTransfer) Type() string { return RefuseDTagTransferRequest }

func (msg MsgRefuseDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRefuseDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRefuseDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}
