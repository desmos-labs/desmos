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

// NewMsgRequestDTagTransfer is a constructor function for MsgRequestDTagTransfer
func NewMsgRequestDTagTransfer(sender, receiver string) *MsgRequestDTagTransfer {
	return &MsgRequestDTagTransfer{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRequestDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRequestDTagTransfer) Type() string { return ActionRequestDTag }

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
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
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
func (msg MsgCancelDTagTransfer) Type() string { return ActionCancelDTagTransferRequest }

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
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
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
		NewDTag:  newDTag,
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgAcceptDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAcceptDTagTransfer) Type() string { return ActionAcceptDTagTransfer }

func (msg MsgAcceptDTagTransfer) ValidateBasic() error {
	if strings.TrimSpace(msg.NewDTag) == "" {
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
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
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
func (msg MsgRefuseDTagTransfer) Type() string { return ActionRefuseDTagTransferRequest }

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
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRefuseDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}
