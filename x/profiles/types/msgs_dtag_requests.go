package types

import (
	"fmt"
	"strings"

	"cosmossdk.io/errors"
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
func (msg *MsgRequestDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgRequestDTagTransfer) Type() string { return ActionRequestDTag }

func (msg *MsgRequestDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Sender == msg.Receiver {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRequestDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgRequestDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgCancelDTagTransferRequest is a constructor for MsgCancelDTagTransferRequest
func NewMsgCancelDTagTransferRequest(sender, receiver string) *MsgCancelDTagTransferRequest {
	return &MsgCancelDTagTransferRequest{
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg *MsgCancelDTagTransferRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgCancelDTagTransferRequest) Type() string { return ActionCancelDTagTransferRequest }

func (msg *MsgCancelDTagTransferRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver == msg.Sender {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCancelDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCancelDTagTransferRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgAcceptDTagTransferRequest is a constructor for MsgAcceptDTagTransferRequest
func NewMsgAcceptDTagTransferRequest(newDTag string, sender, receiver string) *MsgAcceptDTagTransferRequest {
	return &MsgAcceptDTagTransferRequest{
		NewDTag:  newDTag,
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg *MsgAcceptDTagTransferRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgAcceptDTagTransferRequest) Type() string { return ActionAcceptDTagTransfer }

func (msg *MsgAcceptDTagTransferRequest) ValidateBasic() error {
	if strings.TrimSpace(msg.NewDTag) == "" {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "new DTag can't be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", msg.Sender)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address: %s", msg.Receiver)
	}

	if msg.Sender == msg.Receiver {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgAcceptDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgAcceptDTagTransferRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRefuseDTagTransferRequest is a constructor for MsgRefuseDTagTransferRequest
func NewMsgRefuseDTagTransferRequest(sender, receiver string) *MsgRefuseDTagTransferRequest {
	return &MsgRefuseDTagTransferRequest{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg *MsgRefuseDTagTransferRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgRefuseDTagTransferRequest) Type() string { return ActionRefuseDTagTransferRequest }

func (msg *MsgRefuseDTagTransferRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender == msg.Receiver {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRefuseDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgRefuseDTagTransferRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}
